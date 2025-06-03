# `go.bytecodealliance.org/x/wasihttp`

Package `go.bytecodealliance.org/x/wasihttp` implements [`wasi:http/proxy`](https://github.com/WebAssembly/wasi-http/blob/v0.2.0/proxy.md) for Go using standard [`net/http`](https://pkg.go.dev/net/http) interfaces.

This package allows Go applications to use the familiar `net/http` API when building WebAssembly modules that use the WASI HTTP standard.

## Prerequisites

To use this package, you'll need:

- [TinyGo](https://tinygo.org/) 0.34.0 or later.
- [Wasmtime](https://wasmtime.dev/) 26.0.0 or later.

## Usage

### Server

Create a standard Go HTTP server using the familiar `net/http` API:

```go
package main

import (
    "net/http"
    _ "go.bytecodealliance.org/x/wasihttp" // enable wasi-http
)

func init() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("X-Go", "Gopher")
        w.Write([]byte("Hello world!\n"))
    })
}

func main() {}
```

Compile with:

```bash
tinygo build -target=wasip2-http.json -o server.wasm ./main.go
```

## Examples

For more examples, see the `examples` directory.
