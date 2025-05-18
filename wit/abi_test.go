package wit

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestAlign(t *testing.T) {
	tests := []struct {
		ptr   uintptr
		align uintptr
		want  uintptr
	}{
		{0, 1, 0},
		{0, 2, 0},
		{0, 4, 0},
		{0, 8, 0},
		{1, 1, 1},
		{1, 2, 2},
		{1, 4, 4},
		{1, 8, 8},
		{2, 1, 2},
		{2, 2, 2},
		{2, 4, 4},
		{2, 8, 8},
		{3, 1, 3},
		{3, 2, 4},
		{3, 4, 4},
		{3, 8, 8},
		{4, 1, 4},
		{4, 2, 4},
		{4, 4, 4},
		{4, 8, 8},
		{5, 1, 5},
		{5, 2, 6},
		{5, 4, 8},
		{5, 8, 8},
		{6, 1, 6},
		{6, 2, 6},
		{6, 4, 8},
		{6, 8, 8},
		{7, 1, 7},
		{7, 2, 8},
		{7, 4, 8},
		{7, 8, 8},
		{8, 1, 8},
		{8, 2, 8},
		{8, 4, 8},
		{8, 8, 8},
		{9, 1, 9},
		{9, 2, 10},
		{9, 4, 12},
		{9, 8, 16},
		{10, 1, 10},
		{10, 2, 10},
		{10, 4, 12},
		{10, 8, 16},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%d,%d=%d", tt.ptr, tt.align, tt.want)
		t.Run(name, func(t *testing.T) {
			got := Align(tt.ptr, tt.align)
			if got != tt.want {
				t.Errorf("Align(%d, %d): expected %d, got %d", tt.ptr, tt.align, tt.want, got)
			}
		})
	}
}

func TestDiscriminant(t *testing.T) {
	tests := []struct {
		n    int
		want Type
	}{
		{0, U8{}},
		{1, U8{}},
		{5, U8{}},
		{8, U8{}},
		{255, U8{}},
		{256, U8{}},
		{257, U16{}},
		{10000, U16{}},
		{32768, U16{}},
		{65536, U16{}},
		{65537, U32{}},
		{1 << 24, U32{}},
		{math.MaxInt32, U32{}},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%d", tt.n)
		t.Run(name, func(t *testing.T) {
			got := Discriminant(tt.n)
			if got != tt.want {
				t.Errorf("Discriminant(%d): expected %T, got %T", tt.n, tt.want, got)
			}
		})
	}
}

func TestTypeSizeAndAlign(t *testing.T) {
	tests := []struct {
		name  string
		v     Type
		size  uintptr
		align uintptr
	}{
		{"bool", Bool{}, 1, 1},
		{"s8", S8{}, 1, 1},
		{"u8", U8{}, 1, 1},
		{"s16", S16{}, 2, 2},
		{"u16", U16{}, 2, 2},
		{"s32", S32{}, 4, 4},
		{"u32", U32{}, 4, 4},
		{"s64", S64{}, 8, 8},
		{"u64", U64{}, 8, 8},
		{"f32", F32{}, 4, 4},
		{"f64", F64{}, 8, 8},
		{"char", Char{}, 4, 4},
		{"string", String{}, 8, 4},
		{"option<string>", &TypeDef{Kind: &Option{Type: String{}}}, 12, 4},
		{"option<f32>", &TypeDef{Kind: &Option{Type: F32{}}}, 8, 4},
		{"variant", &TypeDef{Kind: &Variant{Cases: []Case{{Type: String{}}, {Type: F64{}}}}}, 16, 8},
		{"record", &TypeDef{Kind: &Record{Fields: []Field{{Type: U16{}}, {Type: &TypeDef{Kind: &Result{OK: U64{}}}}, {Type: U8{}}}}}, 32, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size := tt.v.Size()
			if size != tt.size {
				t.Errorf("(Type).Size(): %d, expected %d", size, tt.size)
			}
			align := tt.v.Align()
			if align != tt.align {
				t.Errorf("(Type).Align(): %d, expected %d", align, tt.align)
			}
		})
	}
}

func TestTypeFlat(t *testing.T) {
	tests := []struct {
		name string
		v    Type
		want []Type
	}{
		{"bool", Bool{}, []Type{U32{}}},
		{"s8", S8{}, []Type{U32{}}},
		{"u8", U8{}, []Type{U32{}}},
		{"s16", S16{}, []Type{U32{}}},
		{"u16", U16{}, []Type{U32{}}},
		{"s32", S32{}, []Type{U32{}}},
		{"u32", U32{}, []Type{U32{}}},
		{"s64", S64{}, []Type{U64{}}},
		{"u64", U64{}, []Type{U64{}}},
		{"f32", F32{}, []Type{F32{}}},
		{"f64", F64{}, []Type{F64{}}},
		{"char", Char{}, []Type{U32{}}},
		{"string", String{}, []Type{PointerTo(U8{}), U32{}}},
		{"option<string>", &TypeDef{Kind: &Option{Type: String{}}}, []Type{U32{}, PointerTo(U8{}), U32{}}},
		{"option<f32>", &TypeDef{Kind: &Option{Type: F32{}}}, []Type{U32{}, F32{}}},
		{"variant", &TypeDef{Kind: &Variant{Cases: []Case{{Type: String{}}, {Type: F64{}}}}}, []Type{U32{}, U64{}, U32{}}},
		{"record", &TypeDef{Kind: &Record{Fields: []Field{{Type: U16{}}, {Type: &TypeDef{Kind: &Result{OK: U64{}}}}, {Type: U8{}}}}}, []Type{U32{}, U32{}, U64{}, U32{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v.Flat()
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("(Type).Flat(): %v, expected %v", witFor(got...), witFor(tt.want...))
			}
		})
	}
}

func witFor[T Node](nodes ...T) []string {
	out := make([]string, len(nodes))
	for i, node := range nodes {
		out[i] = node.WIT(nil, "")
	}
	return out
}

// TestHasBorrow verifies that HasBorrow returns true for WIT types that contain a Borrow type.
func TestHasBorrow(t *testing.T) {
	testCases := []struct {
		name     string
		typeDef  *TypeDef
		expected bool
	}{
		{"Simple Borrow", makeBorrow(), true},
		{"Record with Borrow", makeTypeDef(&Record{Fields: []Field{{Type: makeBorrow()}}}), true},
		{"List of Record with Borrow", makeTypeDef(&List{Type: makeTypeDef(&Record{Fields: []Field{{Type: makeBorrow()}}})}), true},
		{"Record without Borrow", makeTypeDef(&Record{Fields: []Field{{Type: makeTypeDef(&String{})}}}), false},
		{"Nested Option with Borrow", makeTypeDef(&Option{Type: makeBorrow()}), true},
		{"Nested Option without Borrow", makeTypeDef(&Option{Type: makeTypeDef(&String{})}), false},
		{"Variant with Borrow", makeTypeDef(&Variant{Cases: []Case{{Type: makeBorrow()}}}), true},
		{"Variant without Borrow", makeTypeDef(&Variant{Cases: []Case{{Type: makeTypeDef(&String{})}}}), false},
		{"Result with Borrow in Ok", makeTypeDef(&Result{OK: makeBorrow(), Err: makeTypeDef(&String{})}), true},
		{"Result with Borrow in Err", makeTypeDef(&Result{OK: makeTypeDef(&String{}), Err: makeBorrow()}), true},
		{"Result without Borrow", makeTypeDef(&Result{OK: makeTypeDef(&String{}), Err: makeTypeDef(&String{})}), false},
		{"Future with Borrow", makeTypeDef(&Future{Type: makeBorrow()}), true},
		{"Future without Borrow", makeTypeDef(&Future{Type: makeTypeDef(&String{})}), false},
		{"Tuple with Borrow", makeTypeDef(&Tuple{Types: []Type{makeTypeDef(&String{}), makeBorrow()}}), true},
		{"Tuple without Borrow", makeTypeDef(&Tuple{Types: []Type{makeTypeDef(&String{}), makeTypeDef(&String{})}}), false},
		{"Enum without Borrow", makeTypeDef(&Enum{Cases: []EnumCase{{Name: "A"}, {Name: "B"}}}), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := HasBorrow(tc.typeDef); result != tc.expected {
				t.Errorf("HasBorrow(%q) = %t; want %t", tc.name, result, tc.expected)
			}
		})
	}
}

func makeTypeDef(kind TypeDefKind) *TypeDef {
	return &TypeDef{Kind: kind}
}

func makeBorrow() *TypeDef {
	return &TypeDef{Kind: &Borrow{}}
}
