// Package json enables JSON support for [Component Model] types defined in package cm.
// There are no exported symbols in this package.
//
// To enable JSON support, import this package as _:
//
//	import _ "go.bytecodealliance.org/cm/json"
//
// [Component Model]: https://component-model.bytecodealliance.org/introduction.html
package json

import (
	"encoding/json"
	_ "unsafe"
)

//go:linkname marshal go.bytecodealliance.org/cm.jsonMarshal
func marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

//go:linkname unmarshal go.bytecodealliance.org/cm.jsonUnmarshal
func unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
