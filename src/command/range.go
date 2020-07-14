package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/cmoscofian/meliponto/src/context"
	"github.com/cmoscofian/meliponto/src/model"
	"github.com/cmoscofian/meliponto/src/service"
	"github.com/cmoscofian/meliponto/src/util"
	"github.com/cmoscofian/meliponto/src/util/constants"
)

type RangeCommand Command

func NewRangeCommand() *RangeCommand {
	return &RangeCommand{
		fs: rangeFlagSet,
	}
}

func (d *RangeCommand) Name() string {
	return d.fs.Name()
}

func (d *RangeCommand) Init(args []string) error {
	return d.fs.Parse(args)
}

func (d *RangeCommand) Run(ctx *context.Configuration) error {
	if d.fs.Parsed() {
		chbs := make(chan []byte)
		cher := make(chan error)
		var response []byte
		var err error
		var bodys [][]byte
		var login model.LoginResponse
		var gstart time.Time
		var gend time.Time
		var grange int
		withGard := false

		if help {
			d.fs.Usage()
			return nil
		}

		if d.fs.NArg() < 2 {
			return errors.New(constants.MissingDatesError)
		}

		if onGard != "" && offGard == "" || onGard == "" && offGard != "" {
			return errors.New(constants.MissingGardFlagError)
		}

		if onGard != "" && offGard != "" {
			withGard = true
		}

		start, _, drange, err := util.RangeBetweenDatesInDays(d.fs.Arg(0), d.fs.Arg(1))
		if err != nil {
			return err
		}

		if withGard {
			gstart, gend, grange, err = util.RangeBetweenDatesInDays(onGard, offGard)
			if err != nil {
				return err
			}
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

		for i := 0; i < drange; i++ {
			err = dailyCheck(ctx, start, &bodys, false)
			if err != nil {
				return err
			}

			start = start.Add(24 * time.Hour)
		}

		if withGard {
			gdate := gstart
			for i := 0; i < grange; i++ {
				err := dailyCheckOnGard(ctx, gdate, gstart, gend, &bodys)
				if err != nil {
					return err
				}
				gdate = gdate.Add(24 * time.Hour)
			}
		}

		for _, b := range bodys {
			go service.Punch(token, b, chbs, cher)
		}

		for range bodys {
			select {
			case response = <-chbs:
				pr := new(model.PunchResponse)
				_ = json.Unmarshal(response, pr)
				fmt.Printf("Punch successfull! [id: %s][date: %s][message: %s][state: %s]\n", pr.ID, pr.Date, pr.Message, pr.State)
			case err = <-cher:
				return err
			}
		}

		return nil
	}

	return errors.New(constants.FlagsUnparsedError)
}
