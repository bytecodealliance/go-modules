package environment_test

import (
	"flag"
	"os"
	"testing"
	"unsafe"

	"tests/generated/wasi/cli/v0.2.0/environment"
)

func TestEnvironment(t *testing.T) {
	env := environment.GetEnvironment()
	t.Logf("env len: %d", env.Len())
	for i, kv := range env.Slice() {
		t.Logf("%02d: %s = %s", i, truncate(kv[0], 32), truncate(kv[1], 32))
	}
}

var bigArg = flag.String("big", "", "big argument")

func TestArguments(t *testing.T) {
	args := environment.GetArguments().Slice()

	// t.Errorf("os.Args: %v\nargs: %v", os.Args, args)
	t.Errorf("os.Args: %v", unsafe.SliceData(os.Args))
	t.Errorf("os.Args[0]: %v", unsafe.StringData(os.Args[0]))
	t.Errorf("len(os.Args): %d", len(os.Args))
	t.Errorf("os.Args[1]: %s", os.Args[1])

	for i, arg := range args {
		t.Logf("%02d: %s", i, truncate(arg, 64))
		if !testing.Verbose() && arg == "-test.v" {
			t.Errorf("testing.Verbose() == false with %s arg", arg)
		}
	}

	for i := 0; i < 1000; i++ {
		args2 := environment.GetArguments().Slice()
		if want, got := len(args), len(args2); got != want {
			t.Errorf("len(args2): %d, expected %d", got, want)
		}
	}
}

func truncate(s string, n int) string {
	if len(s) > n {
		s = s[:n] + "…"
	}
	return s
}
