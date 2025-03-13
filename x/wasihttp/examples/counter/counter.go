// This example implements a web server with a counter running in a goroutine.
// This demonstrates instance reuse by the host.
//
// To run: `tinygo run -target=wasip2-http.json ./examples/counter`
// Test /: `curl -v 'http://0.0.0.0:8080/'`

package main

import (
	"fmt"
	"net/http"

	_ "go.bytecodealliance.org/x/wasihttp"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n := <-c
		fmt.Fprintf(w, "%d", n)
	})

	go func() {
		var n int
		for {
			c <- n
			n++
		}
	}()
}

var c = make(chan int, 1)

func main() {}
