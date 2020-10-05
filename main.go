package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/cmoscofian/meliponto/src/command"
	"github.com/cmoscofian/meliponto/src/context"
	"github.com/cmoscofian/meliponto/src/util"
	"github.com/cmoscofian/meliponto/src/util/constants"
)

func root(args []string) error {
	ctx := context.Create()

	cmds := []command.Commander{
		command.NewConfigCommand(),
		command.NewGetTokenCommand(),
		command.NewGardCommand(),
		command.NewSingleCommand(),
		command.NewDayCommand(),
		command.NewRangeCommand(),
		command.NewReportCommand(),
		command.NewVersionCommand(),
	}

	if len(args) < 1 {
		return errors.New(constants.NoCommandError)
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

	return fmt.Errorf(constants.InvalidCommandError, option)
}

func main() {
	if err := root(os.Args[1:]); err != nil {
		fmt.Printf("Error: %s\n\n", err)
		util.PrintUsage()
		os.Exit(1)
	}
}
