package command

import (
	"errors"
	"flag"

	"github.com/cmoscofian/meliponto/src/cli/util/constant"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/domain/repositories"
)

// day is the implementation of the `day` command.
// A punch command for handling full day punches based
// on a valid context config file.
type day struct {
	fs *flag.FlagSet
	ls repositories.LoginService
}

// NewDay returns a new DayCommand pointer setting up
// it's valid flagset.
func NewDay(ls repositories.LoginService) Command {
	return &day{
		fs: dayFlagSet,
		ls: ls,
	}
}

// Name return the string name set for flagset command.
func (d *day) Name() string {
	return d.fs.Name()
}

// Init parses all the valid flags of the command.
func (d *day) Init(args []string) error {
	return d.fs.Parse(args)
}

// Run is responsible for the logic implementation of the command given a valid
// configuration context.
func (d *day) Run(ctx *entities.Context) error {
	if d.fs.Parsed() {
		return nil
	}

	return errors.New(constant.FlagsUnparsedError)
}
