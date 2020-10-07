package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/cmoscofian/meliponto/src/cli/context"
	"github.com/cmoscofian/meliponto/src/cli/util"
	"github.com/cmoscofian/meliponto/src/cli/util/constant"
)

func root(args []string) error {
	ctx := context.New()
	cmds := buildCommands()

	if len(args) < 1 {
		return errors.New(constant.NoCommandError)
	}

	option := args[0]

	for _, cmd := range cmds {
		if cmd.Name() == option {
			if err := cmd.Init(os.Args[2:]); err != nil {
				return err
			}
			return cmd.Run(ctx)
		}
	}

	return fmt.Errorf(constant.InvalidCommandError, option)
}

func main() {
	if err := root(os.Args[1:]); err != nil {
		fmt.Printf("Error: %s\n\n", err)
		util.PrintUsage()
		os.Exit(1)
	}
}
