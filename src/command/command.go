package command

import (
	"flag"

	"github.com/cmoscofian/meliponto/src/context"
	"github.com/cmoscofian/meliponto/src/util"
	"github.com/cmoscofian/meliponto/src/util/constants"
)

type Commander interface {
	Init(args []string) error
	Run(ctx *context.Configuration) error
	Name() string
}

type Command struct {
	fs *flag.FlagSet
}

var configFlagSet *flag.FlagSet
var getTokenFlagSet *flag.FlagSet
var singleFlagSet *flag.FlagSet
var dayFlagSet *flag.FlagSet
var rangeFlagSet *flag.FlagSet
var versionFlagSet *flag.FlagSet

var token string
var gard bool

var message string
var begin string
var end string
var ptime string
var date string
var week int
var month int
var help bool

var onGard string
var offGard string

var userID string
var companyID string
var generate bool

func init() {
	configFlagSet = flag.NewFlagSet(constants.ConfigKey, flag.ContinueOnError)
	configFlagSet.StringVar(&userID, "re", "", constants.ReUsageMessage)
	configFlagSet.StringVar(&companyID, "company", "", constants.CompanyUsageMessage)
	configFlagSet.BoolVar(&generate, "generate", false, constants.GenerateUsageMessage)
	configFlagSet.BoolVar(&help, "help", false, constants.HelpUsageMessage)
	configFlagSet.Usage = util.PrintUsage

	getTokenFlagSet = flag.NewFlagSet(constants.GetTokenKey, flag.ContinueOnError)
	getTokenFlagSet.BoolVar(&help, "help", false, constants.HelpUsageMessage)
	getTokenFlagSet.Usage = util.PrintUsage

	singleFlagSet = flag.NewFlagSet(constants.SingleKey, flag.ContinueOnError)
	singleFlagSet.StringVar(&token, "token", "", constants.TokenUsageMessage)
	singleFlagSet.StringVar(&message, "message", "", constants.MessageUsageMessage)
	singleFlagSet.StringVar(&ptime, "time", "", constants.TimeUsageMessage)
	singleFlagSet.StringVar(&date, "date", "", constants.DateUsageMessage)
	singleFlagSet.BoolVar(&help, "help", false, constants.HelpUsageMessage)
	singleFlagSet.Usage = util.PrintUsage

	dayFlagSet = flag.NewFlagSet(constants.DayKey, flag.ContinueOnError)
	dayFlagSet.StringVar(&token, "token", "", constants.TokenUsageMessage)
	dayFlagSet.BoolVar(&gard, "on-gard", false, constants.IsOnGardUsageMessage)
	dayFlagSet.BoolVar(&help, "help", false, constants.HelpUsageMessage)
	dayFlagSet.Usage = util.PrintUsage

	rangeFlagSet = flag.NewFlagSet(constants.RangeKey, flag.ContinueOnError)
	rangeFlagSet.StringVar(&token, "token", "", constants.TokenUsageMessage)
	rangeFlagSet.StringVar(&onGard, "on-gard", "", constants.OnGardUsageMessage)
	rangeFlagSet.StringVar(&offGard, "off-gard", "", constants.OffGardUsageMessage)
	rangeFlagSet.BoolVar(&help, "help", false, constants.HelpUsageMessage)
	rangeFlagSet.Usage = util.PrintUsage

	versionFlagSet = flag.NewFlagSet(constants.VersionKey, flag.ContinueOnError)
}
