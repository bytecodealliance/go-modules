//go:build tinygo

package wit

import (
	"strings"
	"testing"
)

func TestLoadWITUnsupportedBackend(t *testing.T) {
	for _, input := range []struct {
		name string
		load func() (*Resolve, error)
	}{
		{"file", func() (*Resolve, error) { return LoadWIT("example.wit") }},
		{"reader", func() (*Resolve, error) { return DecodeWIT(strings.NewReader("package example:paths;")) }},
	} {
		t.Run(input.name, func(t *testing.T) {
			resolved, err := input.load()
			if resolved != nil || err == nil || !strings.Contains(err.Error(), "wasm-tools disabled under TinyGo") {
				t.Fatalf("expected disabled-backend error, got %v, %v", resolved, err)
			}
		})
	}
}
