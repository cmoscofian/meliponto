package command

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/cmoscofian/meliponto/src/cli/service"
	cliutil "github.com/cmoscofian/meliponto/src/cli/util"
	cliconstant "github.com/cmoscofian/meliponto/src/cli/util/constant"
	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
	"github.com/cmoscofian/meliponto/src/shared/domain/repository"
	"github.com/cmoscofian/meliponto/src/shared/infrastructure/restclient"
	"github.com/cmoscofian/meliponto/src/shared/repositories/rest"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

// getToken is the implementation of the `get-token` command.
// A general purpose command for fetching a valid JWT token from the
// for authenticating from the command line
type getToken struct {
	fs       *flag.FlagSet
	injected bool
	ls       repository.LoginService
}

// NewGetToken returns a new GetTokenCommand pointer setting up
// it's valid flagset.
func NewGetToken() Command {
	return &getToken{
		fs:       getTokenFlagSet,
		injected: false,
	}
}

// Match returns a bool evaluating if the given
// option matches this particular command.
func (g getToken) Match(option string) bool {
	return g.fs.Name() == option
}

// Parse evaluates and parses all given flags and
// arguments. It returns an error when unable to
// to parse all given arguments
func (g getToken) Parse(args []string) error {
	return g.fs.Parse(args)
}

// Inject handles injecting all required dependencies
// for this particular command.
func (g *getToken) Inject() {
	defaultClient := restclient.NewRestClientPool(constant.BaseURI, nil, time.Minute)
	loginClient := rest.NewLogin(defaultClient)
	loginService := service.NewLogin(loginClient)

	g.injected = true
	g.ls = loginService
}

// Run is responsible for the logic implementation of the
// command given a valid configuration context.
func (g getToken) Run(ctx *entity.Context) error {
	if g.fs.Parsed() && g.injected {
		if help {
			g.fs.Usage()
			return nil
		}

		password, err := cliutil.GetPassword()
		if err != nil {
			return err
		}

		token, err = g.ls.HandleLogin(ctx, password)
		if err != nil {
			return err
		}

		fmt.Println(token)

		return nil
	}

	return errors.New(cliconstant.FlagsUnparsedError)
}
