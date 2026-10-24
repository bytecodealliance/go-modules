package gen

import (
	"os"
	"path/filepath"
	"testing"

	"go.bytecodealliance.org/internal/relpath"
)

func TestPackagePath(t *testing.T) {
	wd, err := relpath.Getwd()
	if err != nil {
		t.Error(err)
	}

	got, err := PackagePath(wd)
	if err != nil {
		t.Error(err)
	}
	want := "go.bytecodealliance.org/internal/go/gen"
	if got != want {
		t.Errorf("PackagePath(%q): got %s, expected %s", wd, got, want)
	}

	tmp := os.TempDir()
	_, err = PackagePath(tmp)
	if err == nil {
		t.Errorf("PackagePath(%q): expected error, got nil", tmp)
	}
}

func TestPackagePathNestedModule(t *testing.T) {
	root, err := os.MkdirTemp("", t.Name()+"-")
	if err != nil {
		t.Fatal(err)
	}
	module := filepath.Join(root, "go.mod")
	internal := filepath.Join(root, "internal")
	nested := filepath.Join(internal, "bindings")
	// TinyGo lacks recursive removal on Windows and WASI.
	t.Cleanup(func() {
		for _, path := range []string{nested, internal, module, root} {
			if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
				t.Errorf("remove test fixture %s: %v", path, err)
			}
		}
	})
	if err := os.WriteFile(module, []byte("module example.com/plugin\n\ngo 1.26.0\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(nested, 0o700); err != nil {
		t.Fatal(err)
	}
	got, err := PackagePath(nested)
	if err != nil {
		t.Fatal(err)
	}
	if want := "example.com/plugin/internal/bindings"; got != want {
		t.Errorf("PackagePath(%q) = %q, want %q", nested, got, want)
	}
}
