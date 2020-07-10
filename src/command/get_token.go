package command

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cmoscofian/meliponto/src/context"
	"github.com/cmoscofian/meliponto/src/model"
	"github.com/cmoscofian/meliponto/src/service"
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
		var response []byte
		var err error
		var login model.LoginResponse

		if help {
			d.fs.Usage()
			return nil
		}

		go service.Login(ctx, chbs, cher)

		select {
		case response = <-chbs:
			err := json.Unmarshal(response, &login)
			if err != nil {
				return err
			}

			if login.Status == model.SuccessStatus {
				fmt.Println(login.Token)
			} else {
				if login.Message != "" {
					return errors.New(login.Message)
				}
				return errors.New(constants.InvalidLoginError)
			}

		case err = <-cher:
			return err
		}

		return nil
	}

	return errors.New(constants.FlagsUnparsedError)
}
