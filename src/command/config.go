package command

import (
	"fmt"

	"github.com/cmoscofian/meliponto/src/context"
)

type ConfigCommand Command

func NewConfigCommand() *ConfigCommand {
	return &ConfigCommand{
		fs: configFlagSet,
	}
}

func (d *ConfigCommand) Name() string {
	return d.fs.Name()
}

func (d *ConfigCommand) Init(args []string) error {
	return d.fs.Parse(args)
}

func (d *ConfigCommand) Run(ctx *context.Configuration) error {
	if help {
		d.fs.Usage()
		return nil
	}

	if generate {
		context.Generate()
		return nil
	}

	if userID != "" {
		if err := ctx.SetUserID(userID); err != nil {
			return err
		}
		fmt.Print("Config file updated successfully!")
	}

	if companyID != "" {
		if err := ctx.SetCompanyID(companyID); err != nil {
			return err
		}
		fmt.Print("Config file updated successfully!")
	}

	return nil
}
