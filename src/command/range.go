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

// RangeCommand is the implementation of the `range` command.
// A punch command for handling full range of punches based
// on a valid context config file.
type RangeCommand Command

// NewRangeCommand returns a new RangeCommand pointer setting up
// it's valid flagset.
func NewRangeCommand() *RangeCommand {
	return &RangeCommand{
		fs: rangeFlagSet,
	}
}

// Name return the string name set for flagset command.
func (d *RangeCommand) Name() string {
	return d.fs.Name()
}

// Init parses all the valid flags of the command.
func (d *RangeCommand) Init(args []string) error {
	return d.fs.Parse(args)
}

// Run is responsible for the logic implementation of the command given a valid
// configuration context.
func (d *RangeCommand) Run(ctx *context.Configuration) error {
	if d.fs.Parsed() {
		chbs := make(chan []byte)
		cher := make(chan error)
		var bodys [][]byte
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

		start, end, drange, err := util.RangeBetweenDatesInDays(d.fs.Arg(0), d.fs.Arg(1))
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
			token, err = handlers.HandleLogin(ctx, chbs, cher)
			if err != nil {
				return err
			}
		}

		if err := handlers.HandleFetch(token, start, end, chbs, cher); err != nil {
			return err
		}

		for i := 0; i < drange; i++ {
			if err := handlers.HandlePunch(ctx, start, &bodys, false); err != nil {
				return err
			}

			start = start.Add(24 * time.Hour)
		}

		if withGard {
			gdate := gstart
			for i := 0; i < grange; i++ {
				if err := handlers.HandleOnGardPunch(ctx, true, gdate, gstart, gend, &bodys); err != nil {
					return err
				}

				gdate = gdate.Add(24 * time.Hour)
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
