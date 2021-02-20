package command

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/cmoscofian/meliponto/src/cli/context"
	"github.com/cmoscofian/meliponto/src/cli/util/constant"
	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
)

// config is the implementation of the `config` command.
// A general purpose command for handling updates to the config file
// from the command line.
type config struct {
	fs       *flag.FlagSet
	injected bool
}

// NewConfig returns a new ConfigCommand pointer setting up
// it's valid flagset.
func NewConfig() Command {
	return &config{
		fs:       configFlagSet,
		injected: false,
	}
}

// Match returns a bool evaluating if the given
// option matches this particular command.
func (c config) Match(option string) bool {
	return c.fs.Name() == option
}

// Parse evaluates and parses all given flags and
// arguments. It returns an error when unable to
// to parse all given arguments
func (c config) Parse(args []string) error {
	return c.fs.Parse(args)
}

// Inject handles injecting all required dependencies
// for this particular command.
func (c *config) Inject() {
	c.injected = true
}

// Run is responsible for the logic implementation of the
// command given a valid configuration context.
func (c config) Run(ctx *entity.Context) error {
	if help {
		c.fs.Usage()
		return nil
	}

	if !generate && userID == "" && companyID == "" {
		return errors.New(constant.MissingFlagsError)
	}

	if generate {
		context.GenerateConfig()
		return nil
	}

	file, err := os.OpenFile(context.ConfigPath, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	if userID != "" {
		if err := ctx.SetUserID(userID, file); err != nil {
			return err
		}
		fmt.Print(constant.ConfigUpdatedSuccessful)
	}

	if companyID != "" {
		if err := ctx.SetCompanyID(companyID, file); err != nil {
			return err
		}
		fmt.Print(constant.ConfigUpdatedSuccessful)
	}

	return nil
}
