package command

import (
	"errors"
	"fmt"

	"github.com/cmoscofian/meliponto/src/context"
	"github.com/cmoscofian/meliponto/src/util/constants"
)

// ConfigCommand is the implementation of the `config` command.
// A general purpose command for handling updates to the config file
// from the command line.
type ConfigCommand Command

// NewConfigCommand returns a new ConfigCommand pointer setting up
// it's valid flagset.
func NewConfigCommand() *ConfigCommand {
	return &ConfigCommand{
		fs: configFlagSet,
	}
}

// Name return the string name set for flagset command.
func (d *ConfigCommand) Name() string {
	return d.fs.Name()
}

// Init parses all the valid flags of the command.
func (d *ConfigCommand) Init(args []string) error {
	return d.fs.Parse(args)
}

// Run is responsible for the logic implementation of the command given a valid
// configuration context.
func (d *ConfigCommand) Run(ctx *context.Configuration) error {
	if help {
		d.fs.Usage()
		return nil
	}

	if !generate && userID == "" && companyID == "" {
		return errors.New(constants.MissingFlagsError)
	}

	if generate {
		context.Generate()
		return nil
	}

	if userID != "" {
		if err := ctx.SetUserID(userID); err != nil {
			return err
		}
		fmt.Print(constants.ConfigUpdatedSuccessful)
	}

	if companyID != "" {
		if err := ctx.SetCompanyID(companyID); err != nil {
			return err
		}
		fmt.Print(constants.ConfigUpdatedSuccessful)
	}

	return nil
}
