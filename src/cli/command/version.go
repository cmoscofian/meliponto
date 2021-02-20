package command

import (
	"flag"
	"fmt"

	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
)

// version is the implementation of the `version` command.
// A general purpose command for fetching the version of current
// app installed.
type version struct {
	fs       *flag.FlagSet
	injected bool
	ver      string
}

var (
	currentVersion string = "0.0.1"
)

// NewVersion returns a new VersionCommand pointer setting up
// it's valid flagset.
func NewVersion() Command {
	return &version{
		fs:  versionFlagSet,
		ver: currentVersion,
	}
}

// Match returns a bool evaluating if the given
// option matches this particular command.
func (v version) Match(option string) bool {
	return v.fs.Name() == option
}

// Parse evaluates and parses all given flags and
// arguments. It returns an error when unable to
// to parse all given arguments
func (v version) Parse(args []string) error {
	return v.fs.Parse(args)
}

// Inject handles injecting all required dependencies
// for this particular command.
func (v *version) Inject() {
	v.injected = true
}

// Run is responsible for the logic implementation of the
// command given a valid configuration context.
func (v version) Run(ctx *entity.Context) error {
	fmt.Println("meliponto version v" + v.ver)
	return nil
}
