package wasihttp

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"go.bytecodealliance.org/cm"
	"go.bytecodealliance.org/x/wasihttp/internal/wasi/http/types"
	"go.bytecodealliance.org/x/wasihttp/internal/wasi/io/streams"
)

func incomingRequest(req types.IncomingRequest) (*http.Request, error) {
	r := &http.Request{
		Method: fromMethod(req.Method()),
		URL:    incomingURL(req),
		// TODO: Proto, ProtoMajor, ProtoMinor
		Header: fromFields(req.Headers()),
		Host:   req.Authority().Value(),
	}

	body, _, isErr := req.Consume().Result()
	if isErr {
		return nil, errors.New("error consuming wasi-http request")
	}

	r.Body = newBodyReader(body, func(h http.Header) {
		r.Trailer = h
	})

	return r, nil
}

func fromMethod(m types.Method) string {
	if o := m.Other(); o != nil {
		return strings.ToUpper(*o)
	}
	return strings.ToUpper(m.String())
}

func incomingURL(req types.IncomingRequest) *url.URL {
	u := &url.URL{
		Scheme: fromScheme(req.Scheme().Value()),
		Host:   req.Authority().Value(),
	}
	u, _ = u.Parse(req.PathWithQuery().Value())
	return u
}

func fromScheme(s types.Scheme) string {
	if o := s.Other(); o != nil {
		return strings.ToLower(*o)
	}
	return strings.ToLower(s.String())
}

func incomingResponse(res types.IncomingResponse) (*http.Response, error) {
	r := &http.Response{
		Status:     http.StatusText(int(res.Status())),
		StatusCode: int(res.Status()),
		Header:     fromFields(res.Headers()),
	}

	body, _, isErr := res.Consume().Result()
	if isErr {
		return nil, errors.New("error consuming wasi-http response")
	}

	r.Body = newBodyReader(body, func(h http.Header) { r.Trailer = h })

	return r, nil
}

var _ io.ReadCloser = &bodyReader{}

type bodyReader struct {
	body     types.IncomingBody
	trailer  func(http.Header)
	stream   streams.InputStream
	finished bool
}

func newBodyReader(body types.IncomingBody, trailer func(http.Header)) *bodyReader {
	return &bodyReader{
		body:    body,
		trailer: trailer,
	}
}

// TODO: implement buffered reads
func (r *bodyReader) Read(p []byte) (int, error) {
	if r.finished {
		return 0, http.ErrBodyReadAfterClose
	}

	if r.stream == cm.ResourceNone {
		// the first call should always return OK
		r.stream, _, _ = r.body.Stream().Result()
	}

	// TODO: coordinate with runtime to block on multiple pollables.
	poll := r.stream.Subscribe()
	poll.Block()
	poll.ResourceDrop()

	list, err, isErr := r.stream.Read(uint64(len(p))).Result()
	if isErr {
		if err.Closed() {
			err2 := r.finish() // read trailers
			if err2 != nil {
				return 0, err2
			}
			return 0, io.EOF
		}
		return 0, fmt.Errorf("failed to read from InputStream %s", err.LastOperationFailed().ToDebugString())
	}

	copy(p, list.Slice())
	return int(list.Len()), nil
}

func (r *bodyReader) Close() error {
	return r.finish()
}

func (r *bodyReader) finish() error {
	if r.finished {
		return nil
	}
	r.finished = true
	if r.stream != cm.ResourceNone {
		r.stream.ResourceDrop()
	}

	future := types.IncomingBodyFinish(r.body)
	defer future.ResourceDrop()
	p := future.Subscribe()
	p.Block()
	p.ResourceDrop()
	trailersReady := future.Get()
	// TODO: figure out a better way to handle option<result<result<option<trailers>, error-code>>>
	someTrailers, err, isErr := trailersReady.Some().OK().Result()
	if isErr {
		return errors.New(err.String())
	}
	trailers := someTrailers.Some()
	if trailers != nil {
		r.trailer(fromFields(*trailers))
	}

	return nil
}

var (
	_ io.Writer    = &bodyWriter{}
	_ http.Flusher = &bodyWriter{}
)

type bodyWriter struct {
	body     types.OutgoingBody
	trailer  func() http.Header
	stream   streams.OutputStream
	finished bool
}

// newBodyWriter takes ownership of body, allowing it to be written to.
// Call finish to send the HTTP trailers provided by the trailer callback.
func newBodyWriter(body types.OutgoingBody, trailer func() http.Header) *bodyWriter {
	return &bodyWriter{
		body:    body,
		trailer: trailer,
	}
}

// TODO: buffer writes
func (w *bodyWriter) Write(p []byte) (n int, err error) {
	if w.stream == cm.ResourceNone {
		w.stream, _, _ = w.body.Write().Result()
	}
	res := w.stream.BlockingWriteAndFlush(cm.ToList(p))
	if res.IsErr() {
		return 0, fmt.Errorf("wasihttp: %v", res.Err())
	}
	return len(p), nil
}

// TODO: buffer writes
func (w *bodyWriter) Flush() {
	if w.finished {
		return
	}
	if w.stream != cm.ResourceNone {
		w.stream.Flush()
	}
}

func (w *bodyWriter) finish() error {
	if w.finished {
		return nil
	}
	w.finished = true
	if w.stream != cm.ResourceNone {
		w.stream.Flush()
		w.stream.ResourceDrop()
	}

	var trailers cm.Option[types.Trailers]
	if w.trailer != nil {
		trailers = toTrailers(w.trailer())
	}
	finished := types.OutgoingBodyFinish(w.body, trailers)
	if finished.IsErr() {
		return errors.New(finished.Err().String())
	}

	return nil
}

func toScheme(s string) types.Scheme {
	switch s {
	case "http":
		return types.SchemeHTTP()
	case "https":
		return types.SchemeHTTPS()
	default:
		// TODO: when should we set the scheme to `cm.None` if `req.URL.Scheme` is empty?
		return types.SchemeOther(s)
	}
}

func toMethod(s string) types.Method {
	switch s {
	case http.MethodGet:
		return types.MethodGet()
	case http.MethodHead:
		return types.MethodHead()
	case http.MethodPost:
		return types.MethodPost()
	case http.MethodPut:
		return types.MethodPut()
	case http.MethodPatch:
		return types.MethodPatch()
	case http.MethodDelete:
		return types.MethodDelete()
	case http.MethodConnect:
		return types.MethodConnect()
	case http.MethodOptions:
		return types.MethodOptions()
	case http.MethodTrace:
		return types.MethodTrace()
	default:
		// TODO: is other method allowed? or should we return GET?
		// https://cs.opensource.google/go/go/+/refs/tags/go1.23.2:src/net/http/method.go
		// https://github.com/WebAssembly/wasi-http/blob/main/wit/types.wit#L340C41-L341C69
		return types.MethodOther(s)
	}
}

func fromFields(f types.Fields) http.Header {
	h := http.Header{}
	for _, e := range f.Entries().Slice() {
		h.Add(string(e.F0), string(e.F1.Slice()))
	}
	return h
}

func toFields(h http.Header) types.Fields {
	fields := types.NewFields()
	for k, v := range h {
		vals := make([]types.FieldValue, 0, len(v))
		for _, vv := range v {
			vals = append(vals, types.FieldValue(cm.ToList([]uint8(vv))))
		}
		fields.Set(types.FieldKey(k), cm.ToList(vals))
	}
	return fields
}

func toTrailers(h http.Header) cm.Option[types.Trailers] {
	if h == nil || len(h) == 0 {
		return cm.None[types.Trailers]()
	}
	return cm.Some(toFields(h))
}
