package cm

import (
	"unsafe"
)

func Pointers(v any) func(yield func(unsafe.Pointer) bool) {
	switch v := v.(type) {
	case string:
		return pointerSeq(unsafe.Pointer(unsafe.StringData(v)))
	case *string:
		return pointerSeq(unsafe.Pointer(unsafe.StringData(*v)))
	case pointerer:
		return pointerSeq(v.pointer())
	}
	return func(yield func(unsafe.Pointer) bool) {}
}

func pointerSeq(ptr unsafe.Pointer) func(yield func(unsafe.Pointer) bool) {
	return func(yield func(unsafe.Pointer) bool) {
		yield(ptr)
	}
}

func pointersSeq(ptrs ...unsafe.Pointer) func(yield func(unsafe.Pointer) bool) {
	return func(yield func(unsafe.Pointer) bool) {
		for _, ptr := range ptrs {
			if !yield(ptr) {
				return
			}
		}
	}
}

type pointerer interface {
	pointer() unsafe.Pointer
}

func Collect[T any](func(yield func(T) bool)) []T {
	return nil
}
