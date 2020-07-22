package command

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cmoscofian/meliponto/src/context"
	"github.com/cmoscofian/meliponto/src/handlers"
	"github.com/cmoscofian/meliponto/src/model"
	"github.com/cmoscofian/meliponto/src/service"
	"github.com/cmoscofian/meliponto/src/usecase"
	"github.com/cmoscofian/meliponto/src/util"
	"github.com/cmoscofian/meliponto/src/util/constants"
)

// SingleCommand is the implementation of the `single` command.
// A punch command for handling a single punch.
type SingleCommand Command

// NewSingleCommand returns a new SingleCommand pointer setting up
// it's valid flagset.
func NewSingleCommand() *SingleCommand {
	return &SingleCommand{
		fs: singleFlagSet,
	}
}

// Name return the string name set for flagset command.
func (d *SingleCommand) Name() string {
	return d.fs.Name()
}

// Init parses all the valid flags of the command.
func (d *SingleCommand) Init(args []string) error {
	return d.fs.Parse(args)
}

// Run is responsible for the logic implementation of the command given a valid
// configuration context.
func (d *SingleCommand) Run(ctx *context.Configuration) error {
	if d.fs.Parsed() {
		chbs := make(chan []byte)
		cher := make(chan error)

		if help {
			d.fs.Usage()
			return nil
		}

		if date == "" {
			return errors.New(constants.MissingDateFlagError)
		}

		if ptime == "" {
			return errors.New(constants.MissingTimeFlagError)
		}

		day, err := util.ParseFlagDatetime(date, ptime)
		if err != nil {
			return err
		}

		if token == "" {
			token, err = handlers.HandleLogin(ctx, chbs, cher)
			if err != nil {
				return err
			}
		}

		uc := usecase.NewUsecase(util.GetDefaultMessage(ctx, message, day), day, gard)
		body, err := uc.Create()
		if err != nil {
			return err
		}

		go service.PostPunch(token, body, chbs, cher)

		select {
		case response := <-chbs:
			pr := new(model.PunchResponse)
			_ = json.Unmarshal(response, pr)
			fmt.Printf(constants.PunchSuccessful, pr.ID, pr.Date, pr.Message, pr.State)
		case err := <-cher:
			return err
		}

		return nil
	}

	return errors.New(constants.FlagsUnparsedError)
}
