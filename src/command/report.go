package command

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/cmoscofian/meliponto/src/context"
	"github.com/cmoscofian/meliponto/src/handlers"
	"github.com/cmoscofian/meliponto/src/model"
	"github.com/cmoscofian/meliponto/src/util"
	"github.com/cmoscofian/meliponto/src/util/constants"
)

// ReportCommand is the implementation of the `report`` command.
// A general purpose command for generating a report with information
// regarding the range passed as paramethers.
type ReportCommand Command

// NewReportCommand returns a new ReportCommand pointer setting up
// it's valid flagset.
func NewReportCommand() *ReportCommand {
	return &ReportCommand{
		fs: reportFlagSet,
	}
}

// Name return the string name set for flagset command.
func (d *ReportCommand) Name() string {
	return d.fs.Name()
}

// Init parses all the valid flags of the command.
func (d *ReportCommand) Init(args []string) error {
	return d.fs.Parse(args)
}

// Run is responsible for the logic implementation of the command given a valid
// configuration context.
func (d *ReportCommand) Run(ctx *context.Configuration) error {
	if d.fs.Parsed() {
		chbs := make(chan []byte)
		cher := make(chan error)

		if help {
			d.fs.Usage()
			return nil
		}

		if d.fs.NArg() < 2 {
			return errors.New(constants.MissingDatesError)
		}

		start, end, _, err := util.RangeBetweenDatesInDays(d.fs.Arg(0), d.fs.Arg(1))
		if err != nil {
			return err
		}

		if token == "" {
			token, err = handlers.HandleLogin(ctx, chbs, cher)
			if err != nil {
				return err
			}
		}

		query, err := handlers.HandleFetch(token, start, end, chbs, cher)
		if err != nil {
			return err
		}

		if !query.HasData() {
			fmt.Println(constants.NoPreviousPunchesError)
			return nil
		}

		var punches []*model.PunchResponse
		if gard {
			punches = query.GetAllowance()
		} else {
			punches = query.GetRegular()
		}

		if len(punches) <= 0 {
			fmt.Println(constants.NoPreviousPunchesError)
			return nil
		}

		bs, err := util.FormatCSVMessage(ctx, punches)
		if err != nil {
			return err
		}

		if destination == "" {
			destination = os.Getenv("HOME")
		}

		filename := path.Join(destination, fmt.Sprintf("%s_%s_%s.csv", ctx.UserID, start.Format(constants.DateLayout), end.Format(constants.DateLayout)))

		if err := ioutil.WriteFile(filename, bs, 0666); err != nil {
			return err
		}

		fmt.Printf(constants.ReportSuccessful, filename)
		return nil
	}

	return errors.New(constants.FlagsUnparsedError)
}
