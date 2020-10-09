package command

import (
	"flag"
	"fmt"

	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
)

// version is the implementation of the `version` command.
// A general purpose command for fetching the version of current
// app installed.
type version struct {
	fs *flag.FlagSet
}

var (
	currentVersion string = "0.0.1"
)

// NewVersion returns a new VersionCommand pointer setting up
// it's valid flagset.
func NewVersion() Command {
	return &version{
		fs: versionFlagSet,
	}
}

// Name return the string name set for flagset command.
func (d *version) Name() string {
	return d.fs.Name()
}

// Init parses all the valid flags of the command.
func (d *version) Init(args []string) error {
	return d.fs.Parse(args)
}

// Run is responsible for the logic implementation of the command given a valid
// configuration context.
func (d *version) Run(ctx *entities.Context) error {
	fmt.Printf("meliponto version v%s\n", currentVersion)
	return nil
}
