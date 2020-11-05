package command

import (
	"fmt"

	"github.com/cmoscofian/meliponto/src/context"
)

// VersionCommand is the implementation of the `version` command.
// A general purpose command for fetching the version of current
// app installed.
type VersionCommand Command

var (
	version string = "0.0.1"
)

// NewVersionCommand returns a new VersionCommand pointer setting up
// it's valid flagset.
func NewVersionCommand() *VersionCommand {
	return &VersionCommand{
		fs: versionFlagSet,
	}
}

// Name return the string name set for flagset command.
func (d *VersionCommand) Name() string {
	return d.fs.Name()
}

// Init parses all the valid flags of the command.
func (d *VersionCommand) Init(args []string) error {
	return d.fs.Parse(args)
}

// Run is responsible for the logic implementation of the command given a valid
// configuration context.
func (d *VersionCommand) Run(ctx *context.Configuration) error {
	fmt.Printf("meliponto version v%s\n", version)
	return nil
}
