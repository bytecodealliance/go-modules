package cm

import "unsafe"

// Resource represents an opaque Component Model [resource handle].
// It is represented in the [Canonical ABI] as an 32-bit integer.
//
// [resource handle]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/Explainer.md#handle-types
// [Canonical ABI]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md
type Resource uint32

// ResourceNone is a sentinel value indicating a null or uninitialized resource.
// This is a reserved value specified in the [Canonical ABI runtime state].
//
// [Canonical ABI runtime state]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#runtime-state
const ResourceNone = 0

// Rep represents an opaque resource representation.
// It can be any value that fits in 32 bits, but is typically a pointer on wasm32.
type Rep uint32

// Into performs an unsafe cast from [Rep] rep into T.
func Into[T Resourcer](rep Rep) T {
	return *(*T)(unsafe.Pointer(&rep))
}

// IntoRep performs an unsafe cast from T into Rep.
func IntoRep[T any, R RepTypes[T]](v R) Rep {
	return *(*Rep)(unsafe.Pointer(&v))
}

// RepTypes is a type constraint for a concrete resource representation,
// currently represented in the [Canonical ABI] as a 32-bit integer value.
//
// [Canonical ABI]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md
type RepTypes[T any] interface {
	~int32 | ~uint32 | ~uintptr | *T
}

type Own[T Resourcer] struct {
	handle uint32
}

var _ [unsafe.Sizeof(Own[Resourcer]{})]byte = [unsafe.Sizeof(uint32(0))]byte{}

type Resourcer interface {
	// ResourceRep() Rep
	ResourceDestructor()
}
