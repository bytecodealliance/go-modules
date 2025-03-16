//go:build !wasm

package cm

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestDependencies(t *testing.T) {
	cmd := exec.Command("go", "list", "-f", "{{.Imports}}", "-tags", "module.std", ".")
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout
	err := cmd.Run()
	if err != nil {
		t.Fatal(err)
	}

	got := strings.TrimSpace(stdout.String())
	const want = "[errors structs unsafe]" // Should not include "encoding/json"
	if got != want {
		t.Errorf("Expected dependencies %s, got %s", want, got)
	}
}
