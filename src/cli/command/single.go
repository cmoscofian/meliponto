package command

import (
	"errors"
	"flag"

	"github.com/cmoscofian/meliponto/src/cli/util/constant"
	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
	"github.com/cmoscofian/meliponto/src/shared/domain/repository"
)

// single is the implementation of the `single` command.
// A punch command for handling a single punch.
type single struct {
	fs       *flag.FlagSet
	injected bool
	ls       repository.LoginService
}

// NewSingle returns a new SingleCommand pointer setting up
// it's valid flagset.
func NewSingle() Command {
	return &single{
		fs:       singleFlagSet,
		injected: false,
	}
}

// Match returns a bool evaluating if the given
// option matches this particular command.
func (s *single) Match(option string) bool {
	return s.fs.Name() == option
}

// Parse evaluates and parses all given flags and
// arguments. It returns an error when unable to
// to parse all given arguments
func (s *single) Parse(args []string) error {
	return s.fs.Parse(args)
}

// Inject handles injecting all required dependencies
// for this particular command.
func (s *single) Inject() {
	s.injected = true
}

// Run is responsible for the logic implementation of the
// command given a valid configuration context.
func (s *single) Run(ctx *entity.Context) error {
	if s.fs.Parsed() {
		return nil
	}

	return errors.New(constant.FlagsUnparsedError)
}
