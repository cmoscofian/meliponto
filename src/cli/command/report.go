package command

import (
	"errors"
	"flag"

	"github.com/cmoscofian/meliponto/src/cli/util/constant"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/domain/repositories"
)

// report is the implementation of the `report`` command.
// A general purpose command for generating a report with information
// regarding the range passed as paramethers.
type report struct {
	fs *flag.FlagSet
	ls repositories.LoginService
}

// NewReport returns a new ReportCommand pointer setting up
// it's valid flagset.
func NewReport(ls repositories.LoginService) Command {
	return &report{
		fs: reportFlagSet,
		ls: ls,
	}
}

// Name return the string name set for flagset command.
func (d *report) Name() string {
	return d.fs.Name()
}

// Init parses all the valid flags of the command.
func (d *report) Init(args []string) error {
	return d.fs.Parse(args)
}

// Run is responsible for the logic implementation of the command given a valid
// configuration context.
func (d *report) Run(ctx *entities.Context) error {
	if d.fs.Parsed() {
		return nil
	}

	return errors.New(constant.FlagsUnparsedError)
}
