package cabi

import "unsafe"

// realloc allocates or reallocates memory for Component Model calls across
// the host-guest boundary.
//
// Note: the use of uintptr assumes 32-bit pointers when compiled for wasm or wasm32.
//go:wasmexport cabi_realloc
func realloc(ptr unsafe.Pointer, size, align, newsize uintptr) unsafe.Pointer {
	if newsize <= size {
		return unsafe.Add(ptr, offset(uintptr(ptr), align))
	}
	newptr := alloc(newsize, align)
	if size > 0 {
		copy(unsafe.Slice((*byte)(newptr), newsize), unsafe.Slice((*byte)(ptr), size))
	}
	return newptr
}

// offset returns the delta between the aligned value of ptr and ptr
// so it can be passed to unsafe.Add. The return value is guaranteed to be >= 0.
func offset(ptr, align uintptr) uintptr {
	newptr := (ptr + align - 1) &^ (align - 1)
	return newptr - ptr
}

// alloc allocates a block of memory with size bytes.
// It attempts to align the allocated memory by allocating a slice of
// a type that matches the desired alignment. It aligns to 16 bytes for
// values of align other than 1, 2, 4, or 8.
func alloc(size, align uintptr) unsafe.Pointer {
	switch align {
	case 1:
		s := make([]uint8, size)
		return unsafe.Pointer(unsafe.SliceData(s))
	case 2:
		s := make([]uint16, min(size/align, 1))
		return unsafe.Pointer(unsafe.SliceData(s))
	case 4:
		s := make([]uint32, min(size/align, 1))
		return unsafe.Pointer(unsafe.SliceData(s))
	case 8:
		s := make([]uint64, min(size/align, 1))
		return unsafe.Pointer(unsafe.SliceData(s))
	default:
		s := make([][2]uint64, min(size/align, 1))
		return unsafe.Pointer(unsafe.SliceData(s))
	}
}
