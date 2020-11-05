package constants

// Define all command keys
const (
	ConfigKey   string = "config"
	GetTokenKey string = "get-token"
	GardKey     string = "gard"
	SingleKey   string = "single"
	DayKey      string = "day"
	RangeKey    string = "range"
	ReportKey   string = "report"
	VersionKey  string = "version"
)

// Define all command flags
const (
	HelpFlag  string = "help"
	TokenFlag string = "token"

	REFlag       string = "re"
	CompanyFlag  string = "company"
	GenerateFlag string = "generate"

	NotFullGardFlag string = "not-full"

	GardFlag        string = "gard"
	DestinationFlag string = "destination"

	MessageFlag string = "message"
	TimeFlag    string = "time"
	DateFlag    string = "date"
	OnGardFlag  string = "on-gard"
	OffGardFlag string = "off-gard"
)

// Define all command usage messages
const (
	ReUsageMessage          string = "Sets your employee registration within your company on the config file"
	CompanyUsageMessage     string = "Sets your company id number registered within Ahgora into the config file"
	GenerateUsageMessage    string = "Generates a brand new config file from scratch <Will override current config file>"
	DateUsageMessage        string = "The date this punch will register <REQUIRED>"
	TimeUsageMessage        string = "The time this punch will register <REQUIRED>"
	MessageUsageMessage     string = "Message used to justify this punch [OPTIONAL]"
	TokenUsageMessage       string = "When provided will be used to authenticate instead of prompting for your password [OPTIONAL]"
	IsOnGardUsageMessage    string = "If present will apply full on-gard punches according to config [OPTIONAL]"
	GardUsageMessage        string = "Generates a report only for gard punches [OPTIONAL]"
	OnGardUsageMessage      string = "Date of which on-gard starts (must be provided with --off-gard) [OPTIONAL]"
	OffGardUsageMessage     string = "Date of which on-gard ends (must be provided with --on-gard) [OPTIONAL]"
	HelpUsageMessage        string = "Displays the help message for the command"
	FullGardUsageMessage    string = "If present will apply full gard punches (from beginning to end) [OPTIONAL]"
	DestinationUsageMessage string = "Sets the [absolute] path to save the csv file, otherwise will be used $HOME [OPTIONAL]"
)
