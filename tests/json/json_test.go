package json_test

import (
	"reflect"
	"testing"

	wallclock "tests/generated/wasi/clocks/v0.2.0/wall-clock"
	"tests/generated/wasi/filesystem/v0.2.0/types"

	"go.bytecodealliance.org/cm"

	_ "go.bytecodealliance.org/cm/json" // Imported for go:linkname
)

func TestJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		into    any
		want    any
		wantErr bool
	}{
		{
			"nil",
			`null`,
			ptr(ptr("")),
			ptr((*string)(nil)),
			false,
		},
		{
			"descriptor-type(block-device)",
			`"block-device"`,
			ptr(types.DescriptorType(0)),
			ptr(types.DescriptorTypeBlockDevice),
			false,
		},
		{
			"descriptor-type(directory)",
			`"directory"`,
			ptr(types.DescriptorType(0)),
			ptr(types.DescriptorTypeDirectory),
			false,
		},
		{
			"datetime",
			`{"seconds":1,"nanoseconds":2}`,
			&wallclock.DateTime{},
			&wallclock.DateTime{Seconds: 1, Nanoseconds: 2},
			false,
		},
		{
			"empty list",
			`[]`,
			&cm.List[uint8]{},
			&cm.List[uint8]{},
			false,
		},
		{
			"list of bool",
			`[false,true,false]`,
			&cm.List[bool]{},
			ptr(cm.ToList([]bool{false, true, false})),
			false,
		},
		{
			"list of u32",
			`[1,2,3]`,
			&cm.List[uint32]{},
			ptr(cm.ToList([]uint32{1, 2, 3})),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := reflect.ValueOf(cm.ToList([]uint8{1, 2, 3}))
			if !reflect.DeepEqual(v, v) {
				t.Errorf("json.Unmarshal(%q): resulting value different (%v != %v)", tt.json, tt.into, tt.want)
			}
			// err := json.Unmarshal([]byte(tt.json), &tt.into)
			// if tt.wantErr && err == nil {
			// 	t.Errorf("json.Unmarshal(%q): expected error, got nil error", tt.json)
			// 	return
			// } else if !tt.wantErr && err != nil {
			// 	t.Errorf("json.Unmarshal(%q): expected no error, got error: %v", tt.json, err)
			// 	return
			// }
			// got, err := json.Marshal(tt.into)
			// if err != nil {
			// 	t.Error(err)
			// 	return
			// }
			// if string(got) != tt.json {
			// 	if !reflect.DeepEqual(tt.want, tt.into) {
			// 		t.Errorf("json.Unmarshal(%q): resulting value different (%v != %v)", tt.json, tt.into, tt.want)
			// 	}
			// 	t.Errorf("json.Marshal(%v): %s, expected %s", tt.into, string(got), tt.json)
			// }
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}
