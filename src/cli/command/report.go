package command

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/cmoscofian/meliponto/src/cli/util/constant"
	"github.com/cmoscofian/meliponto/src/handler"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/domain/repositories"
	"github.com/cmoscofian/meliponto/src/shared/util"
	shared "github.com/cmoscofian/meliponto/src/shared/util/constant"
)

// report is the implementation of the `report`` command.
// A general purpose command for generating a report with information
// regarding the range passed as paramethers.
type report struct {
	fs *flag.FlagSet
	ls repositories.LoginService
}

// NewReport returns a new ReportCommand pointer setting up
// it's valid flagset.
func NewReport(ls repositories.LoginService) Command {
	return &report{
		fs: reportFlagSet,
		ls: ls,
	}
}

// Name return the string name set for flagset command.
func (d *report) Name() string {
	return d.fs.Name()
}

// Init parses all the valid flags of the command.
func (d *report) Init(args []string) error {
	return d.fs.Parse(args)
}

// Run is responsible for the logic implementation of the command given a valid
// configuration context.
func (d *report) Run(ctx *entities.Context) error {
	if d.fs.Parsed() {
		chbs := make(chan []byte)
		cher := make(chan error)

		if help {
			d.fs.Usage()
			return nil
		}

		if d.fs.NArg() < 2 {
			return errors.New(shared.MissingDatesError)
		}

		start, end, _, err := util.RangeBetweenDatesInDays(d.fs.Arg(0), d.fs.Arg(1))
		if err != nil {
			return err
		}

		if token == "" {
			token, err = d.ls.HandleLogin(ctx, "")
			if err != nil {
				return err
			}
		}

		query, err := handler.HandleFetch(token, start, end, chbs, cher)
		if err != nil {
			return err
		}

		if !query.HasData() {
			fmt.Println(shared.NoPreviousPunchesError)
			return nil
		}

		var punches []*entities.PunchResponse
		if gard {
			punches = query.GetAllowance()
		} else {
			punches = query.GetRegular()
		}
		bs, err := util.FormatCSVMessage(ctx, punches)
		if err != nil {
			return err
		}

		if destination == "" {
			destination = os.Getenv("HOME")
		}

		filename := path.Join(destination, fmt.Sprintf("%s_%s_%s.csv", ctx.UserID, start.Format(shared.DateLayout), end.Format(shared.DateLayout)))

		if err := ioutil.WriteFile(filename, bs, 0666); err != nil {
			return err
		}

		fmt.Printf(constant.ReportSuccessful, filename)
		return nil
	}

	return errors.New(constant.FlagsUnparsedError)
}
