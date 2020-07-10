package constants

// Define command keys
const (
	ConfigKey   string = "config"
	GetTokenKey string = "get-token"
	SingleKey   string = "single"
	DayKey      string = "day"
	RangeKey    string = "range"
	VersionKey  string = "version"
)

// Usage messages for the flags defined
const (
	ReUsageMessage       string = "Sets your employee registration within your company on the config file"
	CompanyUsageMessage  string = "Sets your company id number registered within Ahgora into the config file"
	GenerateUsageMessage string = "Generates a brand new config file from scratch <Will override current config file>"
	DateUsageMessage     string = "The date this punch will register <REQUIRED>"
	TimeUsageMessage     string = "The time this punch will register <REQUIRED>"
	MessageUsageMessage  string = "Message used to justify this punch [OPTIONAL]"
	TokenUsageMessage    string = "When provided will be used to authenticate instead of prompting for your password [OPTIONAL]"
	IsOnGardUsageMessage string = "If present will apply full on-gard punches according to config [OPTIONAL]"
	OnGardUsageMessage   string = "Date of which on-gard starts (must be provided with --off-gard) [OPTIONAL]"
	OffGardUsageMessage  string = "Date of which on-gard ends (must be provided with --on-gard) [OPTIONAL]"
	HelpUsageMessage     string = "Displays the help message for the command"
)
