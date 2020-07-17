package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/cmoscofian/meliponto/src/context"
	"github.com/cmoscofian/meliponto/src/handlers"
	"github.com/cmoscofian/meliponto/src/model"
	"github.com/cmoscofian/meliponto/src/service"
	"github.com/cmoscofian/meliponto/src/util"
	"github.com/cmoscofian/meliponto/src/util/constants"
)

type GardCommand Command

func NewGardCommand() *GardCommand {
	return &GardCommand{
		fs: gardFlagSet,
	}
}

func (d *GardCommand) Name() string {
	return d.fs.Name()
}

func (d *GardCommand) Init(args []string) error {
	return d.fs.Parse(args)
}

func (d *GardCommand) Run(ctx *context.Configuration) error {
	if d.fs.Parsed() {
		chbs := make(chan []byte)
		cher := make(chan error)
		var bodys [][]byte
		var start, end time.Time
		var drange int
		var err error

		if help {
			d.fs.Usage()
			return nil
		}

		if d.fs.NArg() < 1 {
			return errors.New(constants.MissingDateError)
		}

		if token == "" {
			token, err = handlers.HandleLogin(ctx, chbs, cher)
			if err != nil {
				return err
			}
		}

		if d.fs.NArg() > 1 {
			start, end, drange, err = util.RangeBetweenDatesInDays(d.fs.Arg(0), d.fs.Arg(1))
			if err != nil {
				return err
			}

			if err := handlers.HandleFetch(token, start, end, chbs, cher); err != nil {
				return err
			}

			day := start
			for i := 0; i < drange; i++ {
				if err := handlers.HandleOnGardPunch(ctx, !notFull, day, start, end, &bodys); err != nil {
					return err
				}
				day = day.Add(24 * time.Hour)
			}
		} else {
			day, err := util.ParseFlagDate(d.fs.Arg(0))
			if err != nil {
				return err
			}

			if err := handlers.HandleFetch(token, day, day, chbs, cher); err != nil {
				return err
			}

			if err := handlers.HandleOnGardPunch(ctx, false, day, day, day, &bodys); err != nil {
				return err
			}
		}

		for _, b := range bodys {
			go service.PostPunch(token, b, chbs, cher)
		}

		for range bodys {
			select {
			case response := <-chbs:
				pr := new(model.PunchResponse)
				_ = json.Unmarshal(response, pr)
				fmt.Printf(constants.PunchSuccessful, pr.ID, pr.Date, pr.Message, pr.State)
			case err := <-cher:
				return err
			}
		}

		return nil
	}

	return errors.New(constants.FlagsUnparsedError)
}
