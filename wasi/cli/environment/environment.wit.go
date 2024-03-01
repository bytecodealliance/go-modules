// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package environment represents the interface "wasi:cli/environment@0.2.0".
package environment

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// GetArguments represents function "wasi:cli/environment@0.2.0#get-arguments".
//
// Get the POSIX-style arguments to the program.
//
//	get-arguments: func() -> list<string>
//
//go:nosplit
func GetArguments() cm.List[string] {
	var result cm.List[string]
	getArguments(&result)
	return result
}

//go:wasmimport wasi:cli/environment@0.2.0 get-arguments
//go:noescape
func getArguments(result *cm.List[string])

// GetEnvironment represents function "wasi:cli/environment@0.2.0#get-environment".
//
// Get the POSIX-style environment variables.
//
// Each environment variable is provided as a pair of string variable names
// and string value.
//
// Morally, these are a value import, but until value imports are available
// in the component model, this import function should return the same
// values each time it is called.
//
//	get-environment: func() -> list<tuple<string, string>>
//
//go:nosplit
func GetEnvironment() cm.List[[2]string] {
	var result cm.List[[2]string]
	getEnvironment(&result)
	return result
}

//go:wasmimport wasi:cli/environment@0.2.0 get-environment
//go:noescape
func getEnvironment(result *cm.List[[2]string])

// InitialCWD represents function "wasi:cli/environment@0.2.0#initial-cwd".
//
// Return a path that programs should use as their initial current working
// directory, interpreting `.` as shorthand for this.
//
//	initial-cwd: func() -> option<string>
//
//go:nosplit
func InitialCWD() cm.Option[string] {
	var result cm.Option[string]
	initialCWD(&result)
	return result
}

//go:wasmimport wasi:cli/environment@0.2.0 initial-cwd
//go:noescape
func initialCWD(result *cm.Option[string])
