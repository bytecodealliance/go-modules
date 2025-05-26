package wasihttp

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"go.bytecodealliance.org/cm"
	outgoinghandler "go.bytecodealliance.org/x/wasihttp/internal/wasi/http/outgoing-handler"
	"go.bytecodealliance.org/x/wasihttp/internal/wasi/http/types"
)

func init() {
	// Override the default net/http transport and client.
	http.DefaultTransport = &Transport{}
	if http.DefaultClient != nil {
		http.DefaultClient.Transport = &Transport{} // TinyGo has custom implementation of net/http
	}
}

// Transport implements [http.RoundTripper] using [wasi-http] APIs.
//
// [wasi-http]: https://github.com/webassembly/wasi-http
type Transport struct{}

// RoundTrip executes a single HTTP transaction.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Only close the body if it's not nil
	if req.Body != nil {
		defer req.Body.Close()
	}

	// TODO: wrap this into a helper func outgoingRequest?
	r := types.NewOutgoingRequest(toFields(req.Header))
	r.SetAuthority(cm.Some(requestAuthority(req))) // TODO: when should this be cm.None?
	r.SetMethod(toMethod(req.Method))
	r.SetPathWithQuery(requestPath(req))
	r.SetScheme(cm.Some(toScheme(req.URL.Scheme))) // TODO: when should this be cm.None?

	body, _, _ := r.Body().Result() // the first call should always return OK

	// TODO: when are [options] used?
	// [options]: https://github.com/WebAssembly/wasi-http/blob/main/wit/handler.wit#L38-L39
	incoming, err, isErr := outgoinghandler.Handle(r, cm.None[types.RequestOptions]()).Result()
	if isErr {
		// outgoing request is invalid or not allowed to be made
		return nil, errors.New(err.String())
	}
	defer incoming.ResourceDrop()

	// Write request body
	w := newBodyWriter(body, func() http.Header {
		// TODO: extract request trailers
		return nil
	})

	// Only copy from req.Body if it's not nil
	if req.Body != nil {
		if _, err := io.Copy(w, req.Body); err != nil {
			return nil, fmt.Errorf("wasihttp: %v", err)
		}
	}
	w.finish()

	// Wait for response
	poll := incoming.Subscribe()
	if !poll.Ready() {
		poll.Block()
	}
	poll.ResourceDrop()

	future := incoming.Get()
	if future.None() {
		return nil, fmt.Errorf("wasihttp: future response is None after blocking")
	}
	// TODO: figure out a better way to handle option<result<result<incoming-response, error-code>>>
	response, err, isErr := future.Some().OK().Result() // the first call should always return OK
	if isErr {
		// TODO: what do we do with the HTTP proxy error-code?
		return nil, errors.New(err.String())
	}
	// TODO: when should an incoming-response be dropped?
	// defer response.ResourceDrop()

	return incomingResponse(response)
}

func requestAuthority(req *http.Request) string {
	if req.Host == "" {
		return req.URL.Host
	} else {
		return req.Host
	}
}

func requestPath(req *http.Request) cm.Option[string] {
	path := req.URL.RequestURI()
	if path == "" {
		return cm.None[string]()
	}
	return cm.Some(path)
}
