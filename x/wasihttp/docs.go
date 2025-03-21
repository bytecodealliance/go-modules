// Package wasihttp implements [`wasi:http/proxy`](https://github.com/WebAssembly/wasi-http/blob/v0.2.0/proxy.md) for Go using standard [`net/http`](https://pkg.go.dev/net/http) interfaces
//
// It provides an implementation of the wasi:http/proxy interface that allows Go applications
// to use the standard net/http library for both client and server operations within a WebAssembly
// module. This allows Go developers to use familiar HTTP APIs when building WebAssembly modules
// that need to make or handle HTTP requests.
//
// To use this package in your Go application, simply import it:
//
//	import _ "go.bytecodealliance.org/x/wasihttp"
package wasihttp
