package cm

import _ "unsafe"

// Package cm is imported by the TinyGo package syscall,
// so it cannot directly import encoding/json.
// We define stub functions and use go:linkname to get around this.

//go:linkname jsonMarshal
func jsonMarshal(v any) ([]byte, error)

//go:linkname jsonUnmarshal
func jsonUnmarshal(data []byte, v any) error
