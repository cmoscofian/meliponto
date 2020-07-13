package command

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cmoscofian/meliponto/src/context"
	"github.com/cmoscofian/meliponto/src/model"
	"github.com/cmoscofian/meliponto/src/service"
	"github.com/cmoscofian/meliponto/src/usecase"
	"github.com/cmoscofian/meliponto/src/util"
	"github.com/cmoscofian/meliponto/src/util/constants"
)

type SingleCommand Command

func NewSingleCommand() *SingleCommand {
	return &SingleCommand{
		fs: singleFlagSet,
	}
}

func (d *SingleCommand) Name() string {
	return d.fs.Name()
}

func (d *SingleCommand) Init(args []string) error {
	return d.fs.Parse(args)
}

func (d *SingleCommand) Run(ctx *context.Configuration) error {
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

		if date == "" {
			return errors.New(constants.MissingDateFlagError)
		}

		if ptime == "" {
			return errors.New(constants.MissingTimeFlagError)
		}

		day, err := util.ParseFlagDatetime(date, ptime)
		if err != nil {
			return err
		}

		if token == "" {
			go service.Login(ctx, chbs, cher)

			select {
			case response = <-chbs:
				err := json.Unmarshal(response, &login)
				if err != nil {
					return err
				}

				if login.Status == model.SuccessStatus {
					token = login.Token
				} else {
					if login.Message != "" {
						return errors.New(login.Message)
					}
					return errors.New(constants.InvalidLoginError)
				}

			case err = <-cher:
				return err
			}
		}

		uc := usecase.NewUsecase(util.GetDefaultMessage(ctx, message, day), day, gard)
		body, err := uc.Create()
		if err != nil {
			return err
		}

		go service.Punch(token, body, chbs, cher)

		select {
		case response = <-chbs:
			pr := new(model.PunchResponse)
			_ = json.Unmarshal(response, pr)
			fmt.Printf("Punch successfull! [id: %s][date: %s][message: %s][state: %s]\n", pr.ID, pr.Date, pr.Message, pr.State)
		case err = <-cher:
			return err
		}

		return nil
	}

	return errors.New(constants.FlagsUnparsedError)
}
