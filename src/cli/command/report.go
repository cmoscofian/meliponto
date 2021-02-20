package command

import (
	"errors"
	"flag"

	"github.com/cmoscofian/meliponto/src/cli/util/constant"
	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
	"github.com/cmoscofian/meliponto/src/shared/domain/repository"
)

// report is the implementation of the `report`` command.
// A general purpose command for generating a report with information
// regarding the range passed as paramethers.
type report struct {
	fs       *flag.FlagSet
	injected bool
	ls       repository.LoginService
}

// NewReport returns a new ReportCommand pointer setting up
// it's valid flagset.
func NewReport() Command {
	return &report{
		fs:       reportFlagSet,
		injected: false,
	}
}

// Match returns a bool evaluating if the given
// option matches this particular command.
func (r report) Match(option string) bool {
	return r.fs.Name() == option
}

// Parse evaluates and parses all given flags and
// arguments. It returns an error when unable to
// to parse all given arguments
func (r report) Parse(args []string) error {
	return r.fs.Parse(args)
}

// Inject handles injecting all required dependencies
// for this particular command.
func (r *report) Inject() {
	r.injected = true
}

// Run is responsible for the logic implementation of the
// command given a valid configuration context.
func (r report) Run(ctx *entity.Context) error {
	if r.fs.Parsed() {
		return nil
	}

	return errors.New(constant.FlagsUnparsedError)
}
