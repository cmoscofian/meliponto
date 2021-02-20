package command

import (
	"flag"

	"github.com/cmoscofian/meliponto/src/cli/util"
	"github.com/cmoscofian/meliponto/src/cli/util/constant"
	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
)

// Command is the interface that implements every possible command.
type Command interface {
	Match(string) bool
	Parse([]string) error
	Inject()
	Run(*entity.Context) error
}

var configFlagSet *flag.FlagSet
var getTokenFlagSet *flag.FlagSet
var gardFlagSet *flag.FlagSet
var singleFlagSet *flag.FlagSet
var dayFlagSet *flag.FlagSet
var rangeFlagSet *flag.FlagSet
var reportFlagSet *flag.FlagSet
var versionFlagSet *flag.FlagSet

var token string
var gard bool
var notFull bool

var message string
var ptime string
var date string
var help bool

var onGard string
var offGard string

var destination string

var userID string
var companyID string
var generate bool

func init() {
	// General Commands
	configFlagSet = flag.NewFlagSet(constant.ConfigKey, flag.ExitOnError)
	configFlagSet.StringVar(&userID, constant.REFlag, "", constant.ReUsageMessage)
	configFlagSet.StringVar(&companyID, constant.CompanyFlag, "", constant.CompanyUsageMessage)
	configFlagSet.BoolVar(&generate, constant.GenerateFlag, false, constant.GenerateUsageMessage)
	configFlagSet.BoolVar(&help, constant.HelpFlag, false, constant.HelpUsageMessage)
	configFlagSet.Usage = util.PrintUsage

	getTokenFlagSet = flag.NewFlagSet(constant.GetTokenKey, flag.ExitOnError)
	getTokenFlagSet.BoolVar(&help, constant.HelpFlag, false, constant.HelpUsageMessage)
	getTokenFlagSet.Usage = util.PrintUsage

	reportFlagSet = flag.NewFlagSet(constant.ReportKey, flag.ExitOnError)
	reportFlagSet.StringVar(&token, constant.TokenFlag, "", constant.TokenUsageMessage)
	reportFlagSet.StringVar(&destination, constant.DestinationFlag, "", constant.DestinationUsageMessage)
	reportFlagSet.BoolVar(&gard, constant.GardFlag, false, constant.GardUsageMessage)
	reportFlagSet.BoolVar(&help, constant.HelpFlag, false, constant.HelpUsageMessage)
	reportFlagSet.Usage = util.PrintUsage

	versionFlagSet = flag.NewFlagSet(constant.VersionKey, flag.ExitOnError)

	// Regular Punch Commands
	singleFlagSet = flag.NewFlagSet(constant.SingleKey, flag.ExitOnError)
	singleFlagSet.StringVar(&token, constant.TokenFlag, "", constant.TokenUsageMessage)
	singleFlagSet.StringVar(&message, constant.MessageFlag, "", constant.MessageUsageMessage)
	singleFlagSet.StringVar(&ptime, constant.TimeFlag, "", constant.TimeUsageMessage)
	singleFlagSet.StringVar(&date, constant.DateFlag, "", constant.DateUsageMessage)
	singleFlagSet.BoolVar(&help, constant.HelpFlag, false, constant.HelpUsageMessage)
	singleFlagSet.Usage = util.PrintUsage

	dayFlagSet = flag.NewFlagSet(constant.DayKey, flag.ExitOnError)
	dayFlagSet.StringVar(&token, constant.TokenFlag, "", constant.TokenUsageMessage)
	dayFlagSet.BoolVar(&gard, constant.OnGardFlag, false, constant.IsOnGardUsageMessage)
	dayFlagSet.BoolVar(&help, constant.HelpFlag, false, constant.HelpUsageMessage)
	dayFlagSet.Usage = util.PrintUsage

	rangeFlagSet = flag.NewFlagSet(constant.RangeKey, flag.ExitOnError)
	rangeFlagSet.StringVar(&token, constant.TokenFlag, "", constant.TokenUsageMessage)
	rangeFlagSet.StringVar(&onGard, constant.OnGardFlag, "", constant.OnGardUsageMessage)
	rangeFlagSet.StringVar(&offGard, constant.OffGardFlag, "", constant.OffGardUsageMessage)
	rangeFlagSet.BoolVar(&help, constant.HelpFlag, false, constant.HelpUsageMessage)
	rangeFlagSet.Usage = util.PrintUsage

	// Gard Punch Commands
	gardFlagSet = flag.NewFlagSet(constant.GardKey, flag.ExitOnError)
	gardFlagSet.StringVar(&token, constant.TokenFlag, "", constant.TokenUsageMessage)
	gardFlagSet.BoolVar(&notFull, constant.NotFullGardFlag, false, constant.FullGardUsageMessage)
	gardFlagSet.BoolVar(&help, constant.HelpFlag, false, constant.HelpUsageMessage)
	gardFlagSet.Usage = util.PrintUsage
}
