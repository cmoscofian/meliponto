package context

// The Configuration represents the top-level structure
// of the context passed along throghout
// the app.
type Configuration struct {
	UserID    string       `json:"user_id"`
	CompanyID string       `json:"company_id"`
	Gard      GardField    `json:"gard"`
	Default   DefaultField `json:"default"`
	Holidays  []string     `json:"holidays"`
}

// The GardField represents the mid-level structure of
// the gard information relevant to all gard
// punches.
type GardField struct {
	Messages GardFieldMessages `json:"messages"`
	Hours    GardFieldHours    `json:"hours"`
}

// The GardFieldMessages represents the message info
// level of all gard punches.
type GardFieldMessages struct {
	Default string `json:"default"`
}

// The GardFieldHours represents all possible different
// cenarios for a full week of gard punches.
// It gives flexibility to have different punches in specific days.
type GardFieldHours struct {
	Begin    []GardFieldHoursRange `json:"begin"`
	Weekday  []GardFieldHoursRange `json:"weekday"`
	Saturday []GardFieldHoursRange `json:"saturday"`
	Sunday   []GardFieldHoursRange `json:"sunday"`
	Holiday  []GardFieldHoursRange `json:"holiday"`
	Finish   []GardFieldHoursRange `json:"finish"`
}

// The GardFieldHoursRange represents the range in which
// a gard punch can be done into.
type GardFieldHoursRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// The DefaultField represents the mid-level structure
// of all regular punches (non-gard).
type DefaultField struct {
	Messages DefaultFieldConfig `json:"messages"`
	Hours    DefaultFieldConfig `json:"hours"`
}

// The DefaultFieldConfig represents the bottom-level structure
// of all regular punches (non-gard).
type DefaultFieldConfig struct {
	WorkStart  string `json:"work_start"`
	LunchStart string `json:"lunch_start"`
	LunchEnd   string `json:"lunch_end"`
	WorkEnd    string `json:"work_end"`
}
