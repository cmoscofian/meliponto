package command

import (
	"flag"

	"github.com/cmoscofian/meliponto/src/context"
	"github.com/cmoscofian/meliponto/src/util"
	"github.com/cmoscofian/meliponto/src/util/constants"
)

// Commander is the interface that implements every possible command.
type Commander interface {
	Init(args []string) error
	Run(ctx *context.Configuration) error
	Name() string
}

// Command is a struct type that implements every Commander.
// It has a single field fs (FlagSet) responsible for all valid flags to this command
type Command struct {
	fs *flag.FlagSet // nolint
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
	configFlagSet = flag.NewFlagSet(constants.ConfigKey, flag.ExitOnError)
	configFlagSet.StringVar(&userID, constants.REFlag, "", constants.ReUsageMessage)
	configFlagSet.StringVar(&companyID, constants.CompanyFlag, "", constants.CompanyUsageMessage)
	configFlagSet.BoolVar(&generate, constants.GenerateFlag, false, constants.GenerateUsageMessage)
	configFlagSet.BoolVar(&help, constants.HelpFlag, false, constants.HelpUsageMessage)
	configFlagSet.Usage = util.PrintUsage

	getTokenFlagSet = flag.NewFlagSet(constants.GetTokenKey, flag.ExitOnError)
	getTokenFlagSet.BoolVar(&help, constants.HelpFlag, false, constants.HelpUsageMessage)
	getTokenFlagSet.Usage = util.PrintUsage

	reportFlagSet = flag.NewFlagSet(constants.ReportKey, flag.ExitOnError)
	reportFlagSet.StringVar(&token, constants.TokenFlag, "", constants.TokenUsageMessage)
	reportFlagSet.StringVar(&destination, constants.DestinationFlag, "", constants.DestinationUsageMessage)
	reportFlagSet.BoolVar(&gard, constants.GardFlag, false, constants.GardUsageMessage)
	reportFlagSet.BoolVar(&help, constants.HelpFlag, false, constants.HelpUsageMessage)
	reportFlagSet.Usage = util.PrintUsage

	versionFlagSet = flag.NewFlagSet(constants.VersionKey, flag.ExitOnError)

	// Regular Punch Commands
	singleFlagSet = flag.NewFlagSet(constants.SingleKey, flag.ExitOnError)
	singleFlagSet.StringVar(&token, constants.TokenFlag, "", constants.TokenUsageMessage)
	singleFlagSet.StringVar(&message, constants.MessageFlag, "", constants.MessageUsageMessage)
	singleFlagSet.StringVar(&ptime, constants.TimeFlag, "", constants.TimeUsageMessage)
	singleFlagSet.StringVar(&date, constants.DateFlag, "", constants.DateUsageMessage)
	singleFlagSet.BoolVar(&help, constants.HelpFlag, false, constants.HelpUsageMessage)
	singleFlagSet.Usage = util.PrintUsage

	dayFlagSet = flag.NewFlagSet(constants.DayKey, flag.ExitOnError)
	dayFlagSet.StringVar(&token, constants.TokenFlag, "", constants.TokenUsageMessage)
	dayFlagSet.BoolVar(&gard, constants.OnGardFlag, false, constants.IsOnGardUsageMessage)
	dayFlagSet.BoolVar(&help, constants.HelpFlag, false, constants.HelpUsageMessage)
	dayFlagSet.Usage = util.PrintUsage

	rangeFlagSet = flag.NewFlagSet(constants.RangeKey, flag.ExitOnError)
	rangeFlagSet.StringVar(&token, constants.TokenFlag, "", constants.TokenUsageMessage)
	rangeFlagSet.StringVar(&onGard, constants.OnGardFlag, "", constants.OnGardUsageMessage)
	rangeFlagSet.StringVar(&offGard, constants.OffGardFlag, "", constants.OffGardUsageMessage)
	rangeFlagSet.BoolVar(&help, constants.HelpFlag, false, constants.HelpUsageMessage)
	rangeFlagSet.Usage = util.PrintUsage

	// Gard Punch Commands
	gardFlagSet = flag.NewFlagSet(constants.GardKey, flag.ExitOnError)
	gardFlagSet.StringVar(&token, constants.TokenFlag, "", constants.TokenUsageMessage)
	gardFlagSet.BoolVar(&notFull, constants.NotFullGardFlag, false, constants.FullGardUsageMessage)
	gardFlagSet.BoolVar(&help, constants.HelpFlag, false, constants.HelpUsageMessage)
	gardFlagSet.Usage = util.PrintUsage
}
