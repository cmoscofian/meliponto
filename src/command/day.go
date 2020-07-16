package command

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cmoscofian/meliponto/src/context"
	"github.com/cmoscofian/meliponto/src/handlers"
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
			if token, err = handlers.HandleLogin(ctx, chbs, cher); err != nil {
				return err
			}
		}

		if err := handlers.HandleFetch(token, day, day, chbs, cher); err != nil {
			return err
		}

		if err := dailyCheck(ctx, day, &bodys, gard); err != nil {
			return err
		}

		for _, b := range bodys {
			go service.PostPunch(token, b, chbs, cher)
		}

		for range bodys {
			select {
			case response = <-chbs:
				pr := new(model.PunchResponse)
				_ = json.Unmarshal(response, pr)
				fmt.Printf(constants.PunchSuccessful, pr.ID, pr.Date, pr.Message, pr.State)
			case err = <-cher:
				return err
			}
		}

		return nil
	}

	return errors.New(constants.FlagsUnparsedError)
}
