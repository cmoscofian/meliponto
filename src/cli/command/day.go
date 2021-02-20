package command

import (
	"errors"
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/cmoscofian/meliponto/src/cli/service"
	cliutil "github.com/cmoscofian/meliponto/src/cli/util"
	cliconstant "github.com/cmoscofian/meliponto/src/cli/util/constant"
	"github.com/cmoscofian/meliponto/src/shared/chain"
	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
	"github.com/cmoscofian/meliponto/src/shared/domain/repository"
	"github.com/cmoscofian/meliponto/src/shared/infrastructure/restclient"
	"github.com/cmoscofian/meliponto/src/shared/repositories/rest"
	"github.com/cmoscofian/meliponto/src/shared/util"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

// day is the implementation of the `day` command.
// A punch command for handling full day punches based
// on a valid context config file.
type day struct {
	fs       *flag.FlagSet
	injected bool
	ls       repository.LoginService
	bs       repository.GenerateBodys
	fps      repository.FetchPunchService
	cpc      repository.CreatePunchClient
}

// NewDay returns a new DayCommand pointer setting up
// it's valid flagset.
func NewDay() Command {
	return &day{
		fs:       dayFlagSet,
		injected: false,
	}
}

// Match returns a bool evaluating if the given
// option matches this particular command.
func (d day) Match(option string) bool {
	return d.fs.Name() == option
}

// Parse evaluates and parses all given flags and
// arguments. It returns an error when unable to
// to parse all given arguments
func (d day) Parse(args []string) error {
	return d.fs.Parse(args)
}

// Inject handles injecting all required dependencies
// for this particular command.
func (d *day) Inject() {
	defaultClient := restclient.NewRestClientPool(constant.BaseURI, nil, time.Minute)
	loginClient := rest.NewLogin(defaultClient)
	fetchPunchClient := rest.NewFetchPunch(defaultClient)
	createPunchClient := rest.NewCreatePunch(defaultClient)

	chainBuilder := chain.New()

	loginService := service.NewLogin(loginClient)
	fetchService := service.NewFetch(fetchPunchClient)
	usecaseService := service.NewUsecase(chainBuilder)

	d.ls = loginService
	d.bs = usecaseService
	d.cpc = createPunchClient
	d.fps = fetchService
	d.injected = true
}

// Run is responsible for the logic implementation of the
// command given a valid configuration context.
func (d day) Run(ctx *entity.Context) error {
	if d.fs.Parsed() && d.injected {
		if help {
			d.fs.Usage()
			return nil
		}

		if d.fs.NArg() == 0 {
			return errors.New(cliconstant.MissingDateFlagError)
		}

		if token == "" {
			password, err := cliutil.GetPassword()
			if err != nil {
				return err
			}

			if token, err = d.ls.HandleLogin(ctx, password); err != nil {
				return err
			}
		}

		day, err := util.ParseInputDate(d.fs.Arg(0))
		if err != nil {
			return err
		}

		query, err := d.fps.HandleFetchRange(ctx, day, day, token)
		if err != nil {
			return err
		}

		cliutil.AllowPunch(query)

		bodys := d.bs.HandleRegularBody(ctx, day)
		if gard {
			bodys = append(bodys, d.bs.HandleGardBody(ctx, day, nil, nil)...)
		}

		var wg sync.WaitGroup

		ch := make(chan interface{}, len(bodys))
		wg.Add(len(bodys))

		for _, b := range bodys {
			go func(p *entity.PunchRequest) {
				defer wg.Done()
				resp, err := d.cpc.CreatePunch(token, p)
				if err != nil {
					ch <- err
					return
				}
				ch <- resp
			}(b)
		}

		wg.Wait()
		close(ch)

		for c := range ch {
			switch r := c.(type) {
			case *entity.PunchResponse:
				fmt.Printf(cliconstant.PunchSuccessful, r.ID, r.Date, r.Message, r.State)
			case error:
				err = r
			}
		}

		if err != nil {
			return err
		}

		return nil
	}

	return errors.New(cliconstant.FlagsUnparsedError)
}
