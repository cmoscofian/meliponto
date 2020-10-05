package model

// The PunchType is a custom type all valid type values
// for the punch request and response type field.
type PunchType string

// Implementation of every possible valid for the PunchType
const (
	AllowancePunch PunchType = "abono"
	RegularPunch   PunchType = "addPunch"
)

// The PunchField implements the field that defines
// the time for a valid default punch
type PunchField struct {
	Time string `json:"hora"`
}

// The AllowanceField defines the structure of the field
// that is responsible for setting the hours for a allowance
// punch (abono).
type AllowanceField struct {
	Period string `json:"periodo"`
	Reason string `json:"motivo"`
	Begin  string `json:"inicio"`
	End    string `json:"termino"`
}

// The PunchRequest implements the request body structure
// for punching a new entrance into from the Ahgora system.
type PunchRequest struct {
	Date          string          `json:"referencia"`
	Type          PunchType       `json:"tipo"`
	Justification string          `json:"justificativa"`
	Message       string          `json:"mensagem"`
	Punch         *PunchField     `json:"addPunch,omitempty"`
	Allowance     *AllowanceField `json:"abono,omitempty"`
}

// The ResponseField implements the field from which
// the previous punches current state is.
type ResponseField struct {
	Message  string `json:"texto"`
	DateTime string `json:"datahora"`
}

// The PunchResponse implements the response body structure
// from a new punch into from the Ahgora system.
type PunchResponse struct {
	ID        string          `json:"id"`
	Date      string          `json:"referencia"`
	Type      PunchType       `json:"tipo"`
	Message   string          `json:"mensagem"`
	DateTime  string          `json:"datahora"`
	State     string          `json:"estado"`
	AddPunch  *PunchField     `json:"addPunch,omitempty"`
	Allowance *AllowanceField `json:"abono,omitempty"`
	Response  *ResponseField  `json:"resposta,omitempty"`
}

// IsAllowance returns a boolean value indicating
// whether or not the punch response type is of value
// "abono".
func (p *PunchResponse) IsAllowance() bool {
	return p.Type == AllowancePunch
}

// IsRegular returns a boolean value indicating
// whether or not the punch response type is of value
// "addPunch".
func (p *PunchResponse) IsRegular() bool {
	return p.Type == RegularPunch
}

// The Query implements the response body structure
// from fetching previous punches.
type Query struct {
	Total int              `json:"total"`
	Data  []*PunchResponse `json:"data"`
}

// HasData returns a boolean value indicating
// whether or not the query response has any
// valid data.
func (q *Query) HasData() bool {
	return q.Total >= 0
}

// GetAllowance returns a slice of pointers to
// PunchResponse entities that are of type
// "abono".
func (q *Query) GetAllowance() []*PunchResponse {
	punches := make([]*PunchResponse, 0)
	for _, p := range q.Data {
		if p.IsAllowance() {
			punches = append(punches, p)
		}
	}

	return punches
}

// GetRegular returns a slice of pointers to
// PunchResponse entities that are of type
// "addPunch".
func (q *Query) GetRegular() []*PunchResponse {
	punches := make([]*PunchResponse, 0)
	for _, p := range q.Data {
		if p.IsRegular() {
			punches = append(punches, p)
		}
	}

	return punches
}
