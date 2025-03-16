//go:build !module.std

package cm

// This file contains JSON-related functionality for Component Model list types.
// To avoid a cyclical dependency on package encoding/json when using this package
// in a Go or TinyGo standard library, do not include files named *_json.go.

import (
	"bytes"
	"encoding/json"
	"unsafe"
)

// MarshalJSON implements json.Marshaler.
func (l list[T]) MarshalJSON() ([]byte, error) {
	if l.len == 0 {
		return []byte("[]"), nil
	}

	s := l.Slice()
	var zero T
	if unsafe.Sizeof(zero) == 1 {
		// The default Go json.Encoder will marshal []byte as base64.
		// We override that behavior so all int types have the same serialization format.
		// []uint8{1,2,3} -> [1,2,3]
		// []uint32{1,2,3} -> [1,2,3]
		return json.Marshal(sliceOf(s))
	}
	return json.Marshal(s)
}

type slice[T any] []entry[T]

func sliceOf[S ~[]E, E any](s S) slice[E] {
	return *(*slice[E])(unsafe.Pointer(&s))
}

type entry[T any] [1]T

func (v entry[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(v[0])
}

// UnmarshalJSON implements json.Unmarshaler.
func (l *list[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullLiteral) {
		return nil
	}

	var s []T
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	l.data = unsafe.SliceData([]T(s))
	l.len = uintptr(len(s))

	return nil
}

// nullLiteral is the JSON representation of a null literal.
// By convention, to approximate the behavior of Unmarshal itself,
// Unmarshalers implement UnmarshalJSON([]byte("null")) as a no-op.
// See https://pkg.go.dev/encoding/json#Unmarshaler for more information.
var nullLiteral = []byte("null")
