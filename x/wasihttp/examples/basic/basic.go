// This example implements a basic web server.
//
// To run: `tinygo run -target=wasip2-http.json ./examples/basic`
// Test /: `curl -v 'http://0.0.0.0:8080/'`
// Test /error: `curl -v 'http://0.0.0.0:8080/error'`

package main

import (
	"net/http"

	_ "go.bytecodealliance.org/x/wasihttp"
)

func init() {
	// TODO: use "GET /" when TinyGo supports net/http from Go 1.22+
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Go", "Gopher")
		w.Write([]byte("Hello world!\n"))
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		// do nothing, force default response handling
	})
}

func main() {}
