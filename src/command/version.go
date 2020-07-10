package command

import (
	"fmt"

	"github.com/cmoscofian/meliponto/src/context"
)

type VersionCommand Command

var (
	version string = "0.0.1"
)

func NewVersionCommand() *VersionCommand {
	return &VersionCommand{
		fs: versionFlagSet,
	}
}

func (d *VersionCommand) Name() string {
	return d.fs.Name()
}

func (d *VersionCommand) Init(args []string) error {
	return d.fs.Parse(args)
}

func (d *VersionCommand) Run(ctx *context.Configuration) error {
	fmt.Printf("meliponto version v%s\n", version)
	return nil
}
