package command

import (
	"errors"
	"flag"
	"fmt"

	"github.com/cmoscofian/meliponto/src/cli/util/constant"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/domain/repositories"
	"github.com/cmoscofian/meliponto/src/shared/util"
)

// single is the implementation of the `single` command.
// A punch command for handling a single punch.
type single struct {
	fs *flag.FlagSet
	ls repositories.LoginService
}

// NewSingle returns a new SingleCommand pointer setting up
// it's valid flagset.
func NewSingle(ls repositories.LoginService) Command {
	return &single{
		fs: singleFlagSet,
		ls: ls,
	}
}

// Name return the string name set for flagset command.
func (d *single) Name() string {
	return d.fs.Name()
}

// Init parses all the valid flags of the command.
func (d *single) Init(args []string) error {
	return d.fs.Parse(args)
}

// Run is responsible for the logic implementation of the command given a valid
// configuration context.
func (d *single) Run(ctx *entities.Context) error {
	if d.fs.Parsed() {
		if help {
			d.fs.Usage()
			return nil
		}

		if date == "" {
			return errors.New(constant.MissingDateFlagError)
		}

		if ptime == "" {
			return errors.New(constant.MissingTimeFlagError)
		}

		day, err := util.ParseInputDateTime(date, ptime)
		if err != nil {
			return err
		}

		fmt.Println(day)

		if token == "" {
			token, err = d.ls.HandleLogin(ctx, "")
			if err != nil {
				return err
			}
		}

		return nil
	}

	return errors.New(constant.FlagsUnparsedError)
}
