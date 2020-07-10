package command

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cmoscofian/meliponto/src/context"
	"github.com/cmoscofian/meliponto/src/model"
	"github.com/cmoscofian/meliponto/src/service"
	"github.com/cmoscofian/meliponto/src/util"
	"github.com/cmoscofian/meliponto/src/util/constants"
)

type DayCommand Command

func NewDayCommand() *DayCommand {
	return &DayCommand{
		fs: dayFlagSet,
	}
}

func (d *DayCommand) Name() string {
	return d.fs.Name()
}

func (d *DayCommand) Init(args []string) error {
	return d.fs.Parse(args)
}

func (d *DayCommand) Run(ctx *context.Configuration) error {
	if d.fs.Parsed() {
		chbs := make(chan []byte)
		cher := make(chan error)
		var response []byte
		var err error
		var bodys [][]byte
		var login model.LoginResponse

		if help {
			d.fs.Usage()
			return nil
		}

		if d.fs.NArg() == 0 {
			return errors.New(constants.MissingDateError)
		}

		day, err := util.ParseFlagDate(d.fs.Arg(0))
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

		err = dailyCheck(ctx, day, &bodys, gard)
		if err != nil {
			return err
		}

		for _, b := range bodys {
			go service.Punch(token, b, chbs, cher)
		}

		for range bodys {
			select {
			case response = <-chbs:
				pr := new(model.PunchResponse)
				json.Unmarshal(response, pr)
				fmt.Printf("Punch successfull! [id: %s][date: %s][message: %s][state: %s]\n", pr.ID, pr.Date, pr.Message, pr.State)
			case err = <-cher:
				return err
			}
		}

		return nil
	}

	return errors.New(constants.FlagsUnparsedError)
}
