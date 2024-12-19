package main

import (
	"context"
	"fmt"
	"os"

	"go.bytecodealliance.org/internal/wasmtools"
)

func main() {
	ctx := context.Background()
	wasmTools, err := wasmtools.New(ctx)
	if err != nil {
		exit(err)
	}

	var args []string
	if len(os.Args) != 0 {
		args = os.Args[1:]
	}

	err = wasmTools.Run(ctx, os.Stdin, os.Stdout, os.Stderr, nil, args...)
	if err != nil {
		exit(err)
	}
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "wasm-tools: %v\n", err)
	os.Exit(1)
}
