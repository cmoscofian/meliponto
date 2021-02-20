package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/cmoscofian/meliponto/src/cli/command"
	"github.com/cmoscofian/meliponto/src/cli/context"
	"github.com/cmoscofian/meliponto/src/cli/util"
	cliconstant "github.com/cmoscofian/meliponto/src/cli/util/constant"
)

func root(args []string) error {
	ctx := context.New()
	cmds := []command.Command{
		command.NewConfig(),
		command.NewDay(),
		command.NewGard(),
		command.NewGetToken(),
		command.NewRange(),
		command.NewReport(),
		command.NewSingle(),
		command.NewVersion(),
	}

	if len(args) < 1 {
		return errors.New(cliconstant.NoCommandError)
	}

	for _, cmd := range cmds {
		if cmd.Match(args[0]) {
			if err := cmd.Parse(os.Args[2:]); err != nil {
				return err
			}
			cmd.Inject()

			return cmd.Run(ctx)
		}
	}

	return fmt.Errorf(cliconstant.InvalidCommandError, args[0])
}

func main() {
	if err := root(os.Args[1:]); err != nil {
		fmt.Printf("Error: %s\n\n", err)
		util.PrintUsage()
		os.Exit(1)
	}
}
