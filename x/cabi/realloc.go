package cabi

import "unsafe"

// realloc allocates or reallocates memory for Component Model calls across
// the host-guest boundary.
//
// Note: the use of uintptr assumes 32-bit pointers when compiled for wasm or wasm32.
func realloc(ptr unsafe.Pointer, size, align, newsize uintptr) unsafe.Pointer {
	if newsize == 0 {
		return nil
	}
	if newsize <= size && offset(uintptr(ptr), align) == 0 {
		return ptr
	}
	newptr := alloc(newsize, align)
	if size > 0 {
		copy(unsafe.Slice((*byte)(newptr), newsize), unsafe.Slice((*byte)(ptr), min(size, newsize)))
	}
	return newptr
}

// offset returns the delta between the aligned value of ptr and ptr
// so it can be passed to unsafe.Add. The return value is guaranteed to be >= 0.
func offset(ptr, align uintptr) uintptr {
	return -ptr & (allocationAlignment(align) - 1)
}

// alloc allocates a block of memory with size bytes.
// Padding allows alignment without reducing the usable allocation.
// Values of align other than 1, 2, 4, or 8 select 16-byte alignment.
func alloc(size, align uintptr) unsafe.Pointer {
	padding := allocationAlignment(align) - 1
	if size > uintptr(^uint(0)>>1)-padding {
		panic("cabi allocation size overflows int")
	}
	s := make([]byte, max(size, 1)+padding)
	ptr := unsafe.Pointer(unsafe.SliceData(s))
	return unsafe.Add(ptr, offset(uintptr(ptr), align))
}

func allocationAlignment(align uintptr) uintptr {
	switch align {
	case 1, 2, 4, 8:
		return align
	default:
		return 16
	}
}
