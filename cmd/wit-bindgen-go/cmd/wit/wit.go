package wit

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
	"github.com/bytecodealliance/wasm-tools-go/internal/witcli"
)

// Command is the CLI command for wit.
var Command = &cli.Command{
	Name:   "wit",
	Usage:  "reverses a WIT JSON file into WIT syntax",
	Action: action,
}

func action(ctx context.Context, cmd *cli.Command) error {
	path, err := witcli.ParsePaths(cmd.Args().Slice()...)
	if err != nil {
		return err
	}
	res, err := witcli.LoadOne(cmd.Bool("force-wit"), path)
	if err != nil {
		return err
	}
	fmt.Println(res.WIT(nil, ""))
	return nil
}
