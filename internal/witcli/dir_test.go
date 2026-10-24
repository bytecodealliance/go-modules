package witcli

import (
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestFindOrCreateDirFind(t *testing.T) {
	temp, err := os.MkdirTemp("", t.Name()+"-")
	if err != nil {
		t.Fatal(err)
	}

	dir := filepath.Join(temp, "existing")
	cleanupEmptyDirs(t, dir, temp)
	if err := os.Mkdir(dir, os.ModeDir|os.ModePerm); err != nil {
		t.Fatal(err)
	}

	fi, err := os.Stat(dir)
	if err != nil {
		t.Fatal(err)
	}

	info, err := FindOrCreateDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	gotModTime := info.ModTime()
	wantModTime := fi.ModTime()
	if gotModTime != wantModTime {
		t.Errorf("FindOrCreateDirExisting returned directory with modtime %v; want %v", gotModTime, wantModTime)
	}

	// Skip remaining tests on incompatible platforms
	if runtime.Compiler == "tinygo" || strings.Contains(runtime.GOARCH, "wasm") {
		return
	}

	gotPerm := info.Mode().Perm()
	wantPerm := fi.Mode().Perm()
	if gotPerm != wantPerm {
		t.Errorf("FindOrCreateDirExisting returned directory with file permissions %v; want %v", gotPerm, wantPerm)
	}
}

func TestFindOrCreateDirCreate(t *testing.T) {
	temp, err := os.MkdirTemp("", t.Name()+"-")
	if err != nil {
		t.Fatal(err)
	}
	control := filepath.Join(temp, "control")
	dir := filepath.Join(temp, "new")
	cleanupEmptyDirs(t, dir, control, temp)
	if err := os.Mkdir(control, 0o755); err != nil {
		t.Fatal(err)
	}
	controlInfo, err := os.Stat(control)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		t.Errorf("Expected %s to not exist, but it did.", dir)
	}

	info, err := FindOrCreateDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if info == nil || !info.IsDir() {
		t.Fatalf("FindOrCreateDir returned non-directory information: %v", info)
	}
	createdInfo, err := os.Stat(dir)
	if err != nil {
		t.Fatal(err)
	}
	if !createdInfo.IsDir() {
		t.Fatal("FindOrCreateDir did not create a directory")
	}

	// Skip remaining tests on incompatible platforms
	if runtime.Compiler == "tinygo" || strings.Contains(runtime.GOARCH, "wasm") {
		return
	}

	gotPerm := info.Mode().Perm()
	// Compare filesystem-applied permissions, including umask and Windows mode mapping.
	wantPerm := controlInfo.Mode().Perm()
	if gotPerm != wantPerm {
		t.Errorf("FindOrCreateDir returned directory with file permissions %v; want %v", gotPerm, wantPerm)
	}
}

// TinyGo lacks recursive removal on Windows and WASI; remove owned leaves first.
func cleanupEmptyDirs(t *testing.T, paths ...string) {
	t.Helper()
	t.Cleanup(func() {
		for _, path := range paths {
			if err := os.Remove(path); err != nil && !errors.Is(err, fs.ErrNotExist) {
				t.Errorf("remove test directory %s: %v", path, err)
			}
		}
	})
}

func TestFindOrCreateDirPermissionDenied(t *testing.T) {
	// Skip test on incompatible platforms
	if runtime.Compiler == "tinygo" || strings.Contains(runtime.GOARCH, "wasm") {
		return
	}

	temp := t.TempDir()

	protected := filepath.Join(temp, "protected")
	if err := os.Mkdir(protected, 0o755); err != nil {
		t.Fatal(err)
	}
	if runtime.GOOS == "windows" {
		// Windows directory creation uses ACLs; chmod does not remove this permission.
		icacls := filepath.Join(os.Getenv("SystemRoot"), "System32", "icacls.exe")
		if output, err := exec.Command(icacls, protected, "/deny", "*S-1-1-0:(AD)").CombinedOutput(); err != nil {
			t.Fatalf("deny subdirectory creation: %v\n%s", err, output)
		}
		t.Cleanup(func() {
			if output, err := exec.Command(icacls, protected, "/remove:d", "*S-1-1-0").CombinedOutput(); err != nil {
				t.Errorf("restore directory permissions: %v\n%s", err, output)
			}
		})
	} else {
		if err := os.Chmod(protected, 0o400); err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() {
			if err := os.Chmod(protected, 0o755); err != nil {
				t.Errorf("restore directory permissions: %v", err)
			}
		})
	}

	dir := filepath.Join(protected, "permission-denied")

	info, err := FindOrCreateDir(dir)
	if !errors.Is(err, fs.ErrPermission) {
		t.Errorf("Expected creation within protected directory to fail, but it did not: %v", err)
	}
	if info != nil {
		t.Errorf("failed creation returned file information: %v", info)
	}
}
