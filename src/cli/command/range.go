package command

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/cmoscofian/meliponto/src/cli/util/constant"
	"github.com/cmoscofian/meliponto/src/handler"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/domain/repositories"
	"github.com/cmoscofian/meliponto/src/shared/util"
	shared "github.com/cmoscofian/meliponto/src/shared/util/constant"
)

// rangeCommand is the implementation of the `range` command.
// A punch command for handling full range of punches based
// on a valid context config file.
type rangeCommand struct {
	fs *flag.FlagSet
	ls repositories.LoginService
}

// NewRange returns a new RangeCommand pointer setting up
// it's valid flagset.
func NewRange(ls repositories.LoginService) Command {
	return &rangeCommand{
		fs: rangeFlagSet,
		ls: ls,
	}
}

// Name return the string name set for flagset command.
func (d *rangeCommand) Name() string {
	return d.fs.Name()
}

// Init parses all the valid flags of the command.
func (d *rangeCommand) Init(args []string) error {
	return d.fs.Parse(args)
}

// Run is responsible for the logic implementation of the command given a valid
// configuration context.
func (d *rangeCommand) Run(ctx *entities.Context) error {
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
			return errors.New(shared.MissingDatesError)
		}

		if onGard != "" && offGard == "" || onGard == "" && offGard != "" {
			return errors.New(constant.MissingGardFlagError)
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
			token, err = d.ls.HandleLogin(ctx, "")
			if err != nil {
				return err
			}
		}

		if err := handler.HandleFetchToPunch(token, start, end, chbs, cher); err != nil {
			return err
		}

		for i := 0; i < drange; i++ {
			if err := handler.HandlePunch(ctx, start, &bodys, false); err != nil {
				return err
			}

			start = start.Add(24 * time.Hour)
		}

		if withGard {
			gdate := gstart
			for i := 0; i < grange; i++ {
				if err := handler.HandleOnGardPunch(ctx, true, gdate, gstart, gend, &bodys); err != nil {
					return err
				}

				gdate = gdate.Add(24 * time.Hour)
			}
		}

		// for _, b := range bodys {
		// 	go service.PostPunch(token, b, chbs, cher)
		// }

		for range bodys {
			select {
			case response := <-chbs:
				pr := new(entities.PunchResponse)
				_ = json.Unmarshal(response, pr)
				fmt.Printf(constant.PunchSuccessful, pr.ID, pr.Date, pr.Message, pr.State)
			case err := <-cher:
				return err
			}
		}

		return nil
	}

	return errors.New(constant.FlagsUnparsedError)
}
