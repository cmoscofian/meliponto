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

// DayCommand is the implementation of the `day` command.
// A punch command for handling full day punches based
// on a valid context config file.
type DayCommand Command

// NewDayCommand returns a new DayCommand pointer setting up
// it's valid flagset.
func NewDayCommand() *DayCommand {
	return &DayCommand{
		fs: dayFlagSet,
	}
}

// Name return the string name set for flagset command.
func (d *DayCommand) Name() string {
	return d.fs.Name()
}

// Init parses all the valid flags of the command.
func (d *DayCommand) Init(args []string) error {
	return d.fs.Parse(args)
}

// Run is responsible for the logic implementation of the command given a valid
// configuration context.
func (d *DayCommand) Run(ctx *context.Configuration) error {
	if d.fs.Parsed() {
		chbs := make(chan []byte)
		cher := make(chan error)
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

		if err := handlers.HandleFetchToPunch(token, day, day, chbs, cher); err != nil {
			return err
		}

		if err := handlers.HandlePunch(ctx, day, &bodys, gard); err != nil {
			return err
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
