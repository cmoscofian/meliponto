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

// gardCommand is the implementation of the `gard` command.
// A punch command for handling full gard punches based
// on a valid context config file.
type gardCommand struct {
	fs *flag.FlagSet
	ls repositories.LoginService
}

// NewGard returns a new GardCommand pointer setting up
// it's valid flagset.
func NewGard(ls repositories.LoginService) Command {
	return &gardCommand{
		fs: gardFlagSet,
		ls: ls,
	}
}

// Name return the string name set for flagset command.
func (d *gardCommand) Name() string {
	return d.fs.Name()
}

// Init parses all the valid flags of the command.
func (d *gardCommand) Init(args []string) error {
	return d.fs.Parse(args)
}

// Run is responsible for the logic implementation of the command given a valid
// configuration context.
func (d *gardCommand) Run(ctx *entities.Context) error {
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
			return errors.New(shared.MissingDateError)
		}

		if token == "" {
			token, err = d.ls.HandleLogin(ctx, "")
			if err != nil {
				return err
			}
		}

		if d.fs.NArg() > 1 {
			start, end, drange, err = util.RangeBetweenDatesInDays(d.fs.Arg(0), d.fs.Arg(1))
			if err != nil {
				return err
			}

			if err := handler.HandleFetchToPunch(token, start, end, chbs, cher); err != nil {
				return err
			}

			day := start
			for i := 0; i < drange; i++ {
				if err := handler.HandleOnGardPunch(ctx, !notFull, day, start, end, &bodys); err != nil {
					return err
				}
				day = day.Add(24 * time.Hour)
			}
		} else {
			day, err := util.ParseInputDate(d.fs.Arg(0))
			if err != nil {
				return err
			}

			if err := handler.HandleFetchToPunch(token, day, day, chbs, cher); err != nil {
				return err
			}

			if err := handler.HandleOnGardPunch(ctx, false, day, day, day, &bodys); err != nil {
				return err
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
