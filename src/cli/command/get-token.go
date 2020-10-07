package command

import (
	"errors"
	"flag"
	"fmt"

	"github.com/cmoscofian/meliponto/src/cli/util/constant"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/domain/repositories"
)

// getToken is the implementation of the `get-token` command.
// A general purpose command for fetching a valid JWT token from the
// for authenticating from the command line
type getToken struct {
	fs *flag.FlagSet
	ls repositories.LoginService
}

// NewGetToken returns a new GetTokenCommand pointer setting up
// it's valid flagset.
func NewGetToken(ls repositories.LoginService) Command {
	return &getToken{
		fs: getTokenFlagSet,
		ls: ls,
	}
}

// Name return the string name set for flagset command.
func (d *getToken) Name() string {
	return d.fs.Name()
}

// Init parses all the valid flags of the command.
func (d *getToken) Init(args []string) error {
	return d.fs.Parse(args)
}

// Run is responsible for the logic implementation of the command given a valid
// configuration context.
func (d *getToken) Run(ctx *entities.Context) error {
	if d.fs.Parsed() {
		if help {
			d.fs.Usage()
			return nil
		}

		token, err := d.ls.HandleLogin(ctx, "")
		if err != nil {
			return err
		}

		fmt.Println(token)

		return nil
	}

	return errors.New(constant.FlagsUnparsedError)
}
