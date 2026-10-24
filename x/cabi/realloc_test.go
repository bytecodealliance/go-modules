package cabi

import (
	"fmt"
	"slices"
	"testing"
	"unsafe"
)

func TestOffset(t *testing.T) {
	for _, align := range []uintptr{1, 2, 4, 8, 16} {
		for _, ptr := range []uintptr{0, 1, 3, 9, 16, ^uintptr(0) - 1, ^uintptr(0)} {
			want := (align - ptr%align) % align
			if got := offset(ptr, align); got != want {
				t.Errorf("offset(%d, %d) = %d, want %d", ptr, align, got, want)
			}
		}
	}
}

func TestAlloc(t *testing.T) {
	for _, align := range []uintptr{0, 1, 2, 3, 4, 8, 16, 32} {
		wantAlign := align
		if align != 1 && align != 2 && align != 4 && align != 8 {
			wantAlign = 16
		}
		for _, size := range []uintptr{0, 1, 3, 7, 9, 17, 100, 257} {
			t.Run(fmt.Sprintf("align%d/size%d", align, size), func(t *testing.T) {
				ptr := alloc(size, align)
				if uintptr(ptr)%wantAlign != 0 {
					t.Fatalf("allocation %p is not aligned to %d", ptr, wantAlign)
				}
				span := unsafe.Slice((*byte)(ptr), size)
				for i := range span {
					span[i] = byte(i + 1)
				}
				for i, got := range span {
					if want := byte(i + 1); got != want {
						t.Fatalf("byte %d = %d, want %d", i, got, want)
					}
				}
			})
		}
	}
}

func TestAllocOverflow(t *testing.T) {
	for _, align := range []uintptr{1, 2, 4, 8, 16} {
		t.Run(fmt.Sprintf("align%d", align), func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Error("oversized allocation did not panic")
				}
			}()
			alloc(^uintptr(0), align)
		})
	}
}

func TestRealloc(t *testing.T) {
	for _, align := range []uintptr{1, 2, 4, 8, 16} {
		for _, sizes := range [][2]uintptr{{0, 0}, {0, 100}, {17, 17}, {17, 7}, {7, 100}, {17, 0}} {
			size, newsize := sizes[0], sizes[1]
			t.Run(fmt.Sprintf("align%d/%dto%d", align, size, newsize), func(t *testing.T) {
				var ptr unsafe.Pointer
				if size > 0 {
					ptr = alloc(size, align)
				}
				original := unsafe.Slice((*byte)(ptr), size)
				for i := range original {
					original[i] = byte(i + 1)
				}
				want := slices.Clone(original[:min(size, newsize)])
				got := realloc(ptr, size, align, newsize)
				if newsize == 0 {
					if got != nil {
						t.Fatalf("zero-sized reallocation returned %p", got)
					}
					return
				}
				if uintptr(got)%align != 0 {
					t.Fatalf("reallocation %p is not aligned to %d", got, align)
				}
				span := unsafe.Slice((*byte)(got), newsize)
				if !slices.Equal(span[:len(want)], want) {
					t.Fatalf("reallocation prefix = %v, want %v", span[:len(want)], want)
				}
				for i := range span {
					span[i] = byte(i + 1)
				}
				if newsize <= size && got != ptr {
					t.Error("aligned shrinking allocation moved")
				}
			})
		}
	}
}

func TestReallocRealignedShrink(t *testing.T) {
	for _, align := range []uintptr{2, 4, 8, 16} {
		t.Run(fmt.Sprintf("align%d", align), func(t *testing.T) {
			ptr := unsafe.Add(alloc(33, align), 1)
			original := unsafe.Slice((*byte)(ptr), 32)
			for i := range original {
				original[i] = byte(i + 1)
			}
			want := slices.Clone(original[:17])
			got := realloc(ptr, 32, align, 17)
			if uintptr(got)%align != 0 {
				t.Fatalf("reallocation %p is not aligned to %d", got, align)
			}
			span := unsafe.Slice((*byte)(got), 17)
			if !slices.Equal(span, want) {
				t.Fatalf("realigned allocation = %v, want %v", span, want)
			}
		})
	}
}
