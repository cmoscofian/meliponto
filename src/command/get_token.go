package command

import (
	"errors"
	"fmt"

	"github.com/cmoscofian/meliponto/src/context"
	"github.com/cmoscofian/meliponto/src/handlers"
	"github.com/cmoscofian/meliponto/src/util/constants"
)

// GetTokenCommand is the implementation of the `get-token` command.
// A general purpose command for fetching a valid JWT token from the
// for authenticating from the command line
type GetTokenCommand Command

// NewGetTokenCommand returns a new GetTokenCommand pointer setting up
// it's valid flagset.
func NewGetTokenCommand() *GetTokenCommand {
	return &GetTokenCommand{
		fs: getTokenFlagSet,
	}
}

// Name return the string name set for flagset command.
func (d *GetTokenCommand) Name() string {
	return d.fs.Name()
}

// Init parses all the valid flags of the command.
func (d *GetTokenCommand) Init(args []string) error {
	return d.fs.Parse(args)
}

// Run is responsible for the logic implementation of the command given a valid
// configuration context.
func (d *GetTokenCommand) Run(ctx *context.Configuration) error {
	if d.fs.Parsed() {
		chbs := make(chan []byte)
		cher := make(chan error)

		if help {
			d.fs.Usage()
			return nil
		}

		token, err := handlers.HandleLogin(ctx, chbs, cher)
		if err != nil {
			return err
		}

		fmt.Println(token)

		return nil
	}

	return errors.New(constants.FlagsUnparsedError)
}
