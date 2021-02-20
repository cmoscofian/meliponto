package command

import (
	"errors"
	"flag"

	"github.com/cmoscofian/meliponto/src/cli/util/constant"
	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
	"github.com/cmoscofian/meliponto/src/shared/domain/repository"
)

// rangeCommand is the implementation of the `range` command.
// A punch command for handling full range of punches based
// on a valid context config file.
type rangeCommand struct {
	fs       *flag.FlagSet
	injected bool
	ls       repository.LoginService
}

// NewRange returns a new RangeCommand pointer setting up
// it's valid flagset.
func NewRange() Command {
	return &rangeCommand{
		fs:       rangeFlagSet,
		injected: false,
	}
}

// Match returns a bool evaluating if the given
// option matches this particular command.
func (r rangeCommand) Match(option string) bool {
	return r.fs.Name() == option
}

// Parse evaluates and parses all given flags and
// arguments. It returns an error when unable to
// to parse all given arguments
func (r rangeCommand) Parse(args []string) error {
	return r.fs.Parse(args)
}

// Inject handles injecting all required dependencies
// for this particular command.
func (r *rangeCommand) Inject() {
	r.injected = true
}

// Run is responsible for the logic implementation of the
// command given a valid configuration context.
func (r rangeCommand) Run(ctx *entity.Context) error {
	if r.fs.Parsed() {
		return nil
	}

	return errors.New(constant.FlagsUnparsedError)
}
