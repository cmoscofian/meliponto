package context

type Configuration struct {
	UserID    string       `json:"user_id"`
	CompanyID string       `json:"company_id"`
	Gard      GardField    `json:"gard"`
	Default   DefaultField `json:"default"`
	Holidays  []string     `json:"holidays"`
}

type GardField struct {
	Messages GardFieldMessages `json:"messages"`
	Hours    GardFieldHours    `json:"hours"`
}

type GardFieldMessages struct {
	Default string `json:"default"`
}

type GardFieldHours struct {
	Begin    []GardFieldHoursRange `json:"begin"`
	Weekday  []GardFieldHoursRange `json:"weekday"`
	Saturday []GardFieldHoursRange `json:"saturday"`
	Sunday   []GardFieldHoursRange `json:"sunday"`
	Holiday  []GardFieldHoursRange `json:"holiday"`
	Finish   []GardFieldHoursRange `json:"finish"`
}

type GardFieldHoursRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type DefaultField struct {
	Messages DefaultFieldConfig `json:"messages"`
	Hours    DefaultFieldConfig `json:"hours"`
}

type DefaultFieldConfig struct {
	WorkStart  string `json:"work_start"`
	LunchStart string `json:"lunch_start"`
	LunchEnd   string `json:"lunch_end"`
	WorkEnd    string `json:"work_end"`
}
