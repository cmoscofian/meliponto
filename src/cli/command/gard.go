package command

import (
	"errors"
	"flag"

	"github.com/cmoscofian/meliponto/src/cli/util/constant"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/domain/repositories"
)

// gardCommand is the implementation of the `gard` command.
// A punch command for handling full gard punches based
// on a valid context config file.
type gardCommand struct {
	fs *flag.FlagSet
	ls repositories.LoginService
}

// NewGard returns a new GardCommand pointer setting up
// it's valid flagset.
func NewGard(ls repositories.LoginService) Command {
	return &gardCommand{
		fs: gardFlagSet,
		ls: ls,
	}
}

// Name return the string name set for flagset command.
func (d *gardCommand) Name() string {
	return d.fs.Name()
}

// Init parses all the valid flags of the command.
func (d *gardCommand) Init(args []string) error {
	return d.fs.Parse(args)
}

// Run is responsible for the logic implementation of the command given a valid
// configuration context.
func (d *gardCommand) Run(ctx *entities.Context) error {
	if d.fs.Parsed() {
		return nil
	}

	return errors.New(constant.FlagsUnparsedError)
}
