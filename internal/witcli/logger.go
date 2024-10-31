package witcli

import (
	"os"

	"github.com/bytecodealliance/wasm-tools-go/wit/logging"
)

// Logger returns a [logging.Logger] that outputs to stdout.
func Logger(verbose, debug bool) logging.Logger {
	level := logging.LevelWarn
	if debug {
		level = logging.LevelDebug
	} else if verbose {
		level = logging.LevelInfo
	}
	return logging.NewLogger(os.Stderr, level)
}
