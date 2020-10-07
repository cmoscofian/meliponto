package command

import (
	"bytes"
	"errors"
	"flag"
	"fmt"

	"github.com/cmoscofian/meliponto/src/cli/context"
	"github.com/cmoscofian/meliponto/src/cli/util/constant"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
)

// config is the implementation of the `config` command.
// A general purpose command for handling updates to the config file
// from the command line.
type config struct {
	fs *flag.FlagSet
}

// NewConfig returns a new ConfigCommand pointer setting up
// it's valid flagset.
func NewConfig() Command {
	return &config{
		fs: configFlagSet,
	}
}

// Name return the string name set for flagset command.
func (d *config) Name() string {
	return d.fs.Name()
}

// Init parses all the valid flags of the command.
func (d *config) Init(args []string) error {
	return d.fs.Parse(args)
}

// Run is responsible for the logic implementation of the command given a valid
// configuration context.
func (d *config) Run(ctx *entities.Context) error {
	if help {
		d.fs.Usage()
		return nil
	}

	if !generate && userID == "" && companyID == "" {
		return errors.New(constant.MissingFlagsError)
	}

	if generate {
		context.GenerateConfig()
		return nil
	}

	if userID != "" {
		if err := ctx.SetUserID(userID, bytes.NewBuffer([]byte("1"))); err != nil {
			return err
		}
		fmt.Print(constant.ConfigUpdatedSuccessful)
	}

	if companyID != "" {
		if err := ctx.SetCompanyID(companyID, bytes.NewBuffer([]byte("1"))); err != nil {
			return err
		}
		fmt.Print(constant.ConfigUpdatedSuccessful)
	}

	return nil
}
