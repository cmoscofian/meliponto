package entity

import (
	"encoding/json"
	"io"
)

// The Context represents the top-level structure
// of the context passed along throghout
// the app.
type Context struct {
	UserID    string            `json:"user_id"`
	CompanyID string            `json:"company_id"`
	Gard      *GardField        `json:"gard"`
	Default   *DefaultField     `json:"default"`
	Holidays  []string          `json:"holidays"`
	Values    map[string]string `json:"-"`
}

// SetCompanyID updates the "CompanyID" field of the context entity
// based on a companyID string argument provided and a Writer
// to be used as destination.
func (c *Context) SetCompanyID(companyID string, w io.Writer) error {
	c.CompanyID = companyID
	return c.update(w)
}

// SetUserID updates the "UserID" field of the context entity
// based on an userID string argument provided and a Writer
// to be used as destination.
func (c *Context) SetUserID(userID string, w io.Writer) error {
	c.UserID = userID
	return c.update(w)
}

// SetValue sets a string value on the context to be used across
// the entire transaction.
func (c *Context) SetValue(key, value string) {
	c.Values[key] = value
}

// GetValue retrieves a string value set on the context
// transaction.
func (c Context) GetValue(key string) string {
	value := c.Values[key]
	return value
}

func (c *Context) update(w io.Writer) error {
	bs, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}

	if _, err := w.Write(bs); err != nil {
		return err
	}

	return nil
}

// The GardField represents the mid-level structure of
// the gard information relevant to all gard
// punches.
type GardField struct {
	Messages *GardFieldMessages `json:"messages"`
	Hours    *GardFieldHours    `json:"hours"`
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
	Begin    []*GardFieldHoursRange `json:"begin"`
	Weekday  []*GardFieldHoursRange `json:"weekday"`
	Saturday []*GardFieldHoursRange `json:"saturday"`
	Sunday   []*GardFieldHoursRange `json:"sunday"`
	Holiday  []*GardFieldHoursRange `json:"holiday"`
	Finish   []*GardFieldHoursRange `json:"finish"`
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
	Messages *DefaultFieldConfig `json:"messages"`
	Hours    *DefaultFieldConfig `json:"hours"`
}

// The DefaultFieldConfig represents the bottom-level structure
// of all regular punches (non-gard).
type DefaultFieldConfig struct {
	WorkStart  string `json:"work_start"`
	LunchStart string `json:"lunch_start"`
	LunchEnd   string `json:"lunch_end"`
	WorkEnd    string `json:"work_end"`
}
