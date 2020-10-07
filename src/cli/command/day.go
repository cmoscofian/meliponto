package command

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"

	"github.com/cmoscofian/meliponto/src/cli/util/constant"
	"github.com/cmoscofian/meliponto/src/handler"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/domain/repositories"
	"github.com/cmoscofian/meliponto/src/shared/util"
	shared "github.com/cmoscofian/meliponto/src/shared/util/constant"
)

// day is the implementation of the `day` command.
// A punch command for handling full day punches based
// on a valid context config file.
type day struct {
	fs *flag.FlagSet
	ls repositories.LoginService
}

// NewDay returns a new DayCommand pointer setting up
// it's valid flagset.
func NewDay(ls repositories.LoginService) Command {
	return &day{
		fs: dayFlagSet,
		ls: ls,
	}
}

// Name return the string name set for flagset command.
func (d *day) Name() string {
	return d.fs.Name()
}

// Init parses all the valid flags of the command.
func (d *day) Init(args []string) error {
	return d.fs.Parse(args)
}

// Run is responsible for the logic implementation of the command given a valid
// configuration context.
func (d *day) Run(ctx *entities.Context) error {
	if d.fs.Parsed() {
		chbs := make(chan []byte)
		cher := make(chan error)
		var bodys [][]byte

		if help {
			d.fs.Usage()
			return nil
		}

		if d.fs.NArg() == 0 {
			return errors.New(shared.MissingDateError)
		}

		day, err := util.ParseInputDate(d.fs.Arg(0))
		if err != nil {
			return err
		}

		if token == "" {
			if token, err = d.ls.HandleLogin(ctx, ""); err != nil {
				return err
			}
		}

		if err := handler.HandleFetchToPunch(token, day, day, chbs, cher); err != nil {
			return err
		}

		if err := handler.HandlePunch(ctx, day, &bodys, gard); err != nil {
			return err
		}

		// for _, b := range bodys {
		// go service.PostPunch(token, b, chbs, cher)
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
