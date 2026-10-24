//go:build !tinygo

package wit

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadWITPaths(t *testing.T) {
	// Relative paths require the fixture and working directory to share a volume.
	root, err := os.MkdirTemp(".", ".wit-load-paths-*")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(root); err != nil {
			t.Error(err)
		}
	})
	root, err = filepath.Abs(root)
	if err != nil {
		t.Fatal(err)
	}
	file := filepath.Join(root, "nested", "example.wit")
	if err := os.MkdirAll(filepath.Dir(file), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(file, []byte("package example:paths; world example {}"), 0o644); err != nil {
		t.Fatal(err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	rel, err := filepath.Rel(cwd, file)
	if err != nil {
		t.Fatal(err)
	}
	for _, input := range []struct {
		name string
		path string
	}{
		{"absolute", file},
		{"absolute slash", filepath.ToSlash(file)},
		{"relative", rel},
		{"relative slash", filepath.ToSlash(rel)},
	} {
		t.Run(input.name, func(t *testing.T) {
			resolved, err := LoadWIT(input.path)
			if err != nil {
				t.Fatal(err)
			}
			if len(resolved.Worlds) != 1 || resolved.Worlds[0].Name != "example" {
				t.Fatalf("resolved worlds: %v", resolved.Worlds)
			}
		})
	}
}

func TestLoadWITDirectoryDependencies(t *testing.T) {
	root := t.TempDir()
	dependency := filepath.Join(root, "deps", "dep")
	if err := os.MkdirAll(dependency, 0o755); err != nil {
		t.Fatal(err)
	}
	for name, text := range map[string]string{
		filepath.Join(root, "main.wit"):      "package example:paths; world example { import example:dep/api; }",
		filepath.Join(dependency, "api.wit"): "package example:dep; interface api { ping: func(); }",
	} {
		if err := os.WriteFile(name, []byte(text), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	resolved, err := LoadWIT(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(resolved.Packages) != 2 || len(resolved.Worlds) != 1 || resolved.Worlds[0].Name != "example" {
		t.Fatalf("resolved packages/worlds: %v / %v", resolved.Packages, resolved.Worlds)
	}
}
