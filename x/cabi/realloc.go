package cabi

import "unsafe"

// The Go runtime-internal sbrk function.
// Note that this isn't *really* sbrk, and doesn't necessarily result in heap
// growth. Instead, it reserves a block of memory of size n bytes, which then
// isn't used by Go's GC allocator.
// The heap will only actually grown if the requested allocation doesn't fit.
//go:linkname sbrk runtime.sbrk
func sbrk(n uintptr) unsafe.Pointer

// Returns the memory region from `v` to `v + n` to the Go runtime, enabling
// its use by the garbage collector.
//go:linkname sysFreeOS runtime.sysFreeOS
func sysFreeOS(v unsafe.Pointer, n uintptr)

var useGCAllocations = false

func init() {
	// Once `init` is called, the runtime has been initialized, and we can
	// start using managed memory.
	useGCAllocations = true
}

// realloc allocates or reallocates memory for Component Model calls across
// the host-guest boundary.
//
// Note: while in Go, as opposed to TinyGo, `uintptr` is 64-bit, `go:wasmexport`
// properly coerces the parameters/return values to and from 32-bit values when
// compiled for wasm or wasm32.
// TODO: file bug for sysReserveAlignedSbrk not always releasing the memlock.
//go:wasmexport cabi_realloc
func realloc(ptr unsafe.Pointer, size, align, newsize uintptr) unsafe.Pointer {

	// If the Go runtime has not been initialized, we need to
	// allocate memory using sbrk directly.
	// This happens because the runtime itself does two calls to
	// imported functions during initialization, which result in
	// calls to `cabi_realloc` before the runtime is fully initialized.
	if !useGCAllocations {
		if newsize == 0 {
			if ptr != nil {
				// Free the old pointer if it is not nil.
				// Note that in practice, this case isn't ever hit:
				// `cabi_realloc` isn't called with a non-nil pointer before the
				// Go runtime is fully initialized.
				sysFreeOS(ptr, size)
			}

			return nil
		}

		alignedSize := newsize + offset(newsize, align)
		unaligned := sbrk(alignedSize)
		off := offset(uintptr(unaligned), align)
		newptr := unsafe.Add(unaligned, off)
		if ptr != nil && newptr != nil && size > 0 {
			// Copy the old data to the new pointer.
			// Note that in practice, this case isn't ever hit:
			// `cabi_realloc` isn't called with a non-nil pointer.
			copy(unsafe.Slice((*byte)(newptr), size), unsafe.Slice((*byte)(ptr), size))
		}
		if ptr != nil {
			// Free the old pointer if it is not nil.
			sysFreeOS(ptr, size)
		}
		return newptr
	}

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
//
// The allocation itself is done using the Go runtime's GC allocator.
// We then return a raw pointer to the underlying memory, without keeping
// any references to the allocated object itself. That would be very bad
// indeed if there was a risk of the GC running before we're done using
// the memory, or before it's rooted by another GC object.
//
// Since WebAssembly is a single-threaded environment, and the use of the
// returned memory is tightly controlled by generated bindings code, neither
// of these issues is a problem in practice.
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
