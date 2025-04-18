package stack

import (
	"fmt"
	"runtime"
)

func Trace() []string {
	var out []string
	for i := 1; i < 1024; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		out = append(out, fmt.Sprintf("%s:%d", file, line))
	}
	return out
}
