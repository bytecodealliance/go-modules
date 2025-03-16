package cm

import (
	"bytes"
	"testing"
)

func TestListMethods(t *testing.T) {
	want := []byte("hello world")
	type myList List[uint8]
	l := myList(ToList(want))
	got := l.Slice()
	if !bytes.Equal(want, got) {
		t.Errorf("got (%s) != want (%s)", string(got), string(want))
	}
}
