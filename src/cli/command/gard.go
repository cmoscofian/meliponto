package command

import (
	"errors"
	"flag"

	"github.com/cmoscofian/meliponto/src/cli/util/constant"
	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
	"github.com/cmoscofian/meliponto/src/shared/domain/repository"
)

// gardCommand is the implementation of the `gard` command.
// A punch command for handling full gard punches based
// on a valid context config file.
type gardCommand struct {
	fs       *flag.FlagSet
	injected bool
	ls       repository.LoginService
}

// NewGard returns a new GardCommand pointer setting up
// it's valid flagset.
func NewGard() Command {
	return &gardCommand{
		fs:       gardFlagSet,
		injected: false,
	}
}

// Match returns a bool evaluating if the given
// option matches this particular command.
func (g gardCommand) Match(option string) bool {
	return g.fs.Name() == option
}

// Parse evaluates and parses all given flags and
// arguments. It returns an error when unable to
// to parse all given arguments
func (g gardCommand) Parse(args []string) error {
	return g.fs.Parse(args)
}

// Inject handles injecting all required dependencies
// for this particular command.
func (g *gardCommand) Inject() {
	g.injected = true
}

// Run is responsible for the logic implementation of the
// command given a valid configuration context.
func (g gardCommand) Run(ctx *entity.Context) error {
	if g.fs.Parsed() {
		return nil
	}

	return errors.New(constant.FlagsUnparsedError)
}
