// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package exit represents the interface "wasi:cli/exit@0.2.0".
package exit

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// Exit represents function "wasi:cli/exit@0.2.0#exit".
//
// Exit the current instance and any linked instances.
//
//	exit: func(status: result)
//
//go:nosplit
func Exit(status cm.Result) {
	exit(status)
}

//go:wasmimport wasi:cli/exit@0.2.0 exit
//go:noescape
func exit(status cm.Result)
