package command

import (
	"errors"
	"fmt"

	"github.com/cmoscofian/meliponto/src/context"
	"github.com/cmoscofian/meliponto/src/handlers"
	"github.com/cmoscofian/meliponto/src/util/constants"
)

type GetTokenCommand Command

func NewGetTokenCommand() *GetTokenCommand {
	return &GetTokenCommand{
		fs: getTokenFlagSet,
	}
}

func (d *GetTokenCommand) Name() string {
	return d.fs.Name()
}

func (d *GetTokenCommand) Init(args []string) error {
	return d.fs.Parse(args)
}

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
