package command

import (
	"errors"
	"flag"

	"github.com/cmoscofian/meliponto/src/cli/util/constant"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/domain/repositories"
)

// rangeCommand is the implementation of the `range` command.
// A punch command for handling full range of punches based
// on a valid context config file.
type rangeCommand struct {
	fs *flag.FlagSet
	ls repositories.LoginService
}

// NewRange returns a new RangeCommand pointer setting up
// it's valid flagset.
func NewRange(ls repositories.LoginService) Command {
	return &rangeCommand{
		fs: rangeFlagSet,
		ls: ls,
	}
}

// Name return the string name set for flagset command.
func (d *rangeCommand) Name() string {
	return d.fs.Name()
}

// Init parses all the valid flags of the command.
func (d *rangeCommand) Init(args []string) error {
	return d.fs.Parse(args)
}

// Run is responsible for the logic implementation of the command given a valid
// configuration context.
func (d *rangeCommand) Run(ctx *entities.Context) error {
	if d.fs.Parsed() {
		return nil
	}

	return errors.New(constant.FlagsUnparsedError)
}
