package model

type PunchType string

const (
	AllowancePunch PunchType = "abono"
	RegularPunch   PunchType = "addPunch"
)

type PunchField struct {
	Time string `json:"hora"`
}

type AllowanceField struct {
	Period string `json:"periodo"`
	Reason string `json:"motivo"`
	Begin  string `json:"inicio"`
	End    string `json:"termino"`
}

type PunchRequest struct {
	Date          string          `json:"referencia"`
	Type          PunchType       `json:"tipo"`
	Justification string          `json:"justificativa"`
	Message       string          `json:"mensagem"`
	Punch         *PunchField     `json:"addPunch,omitempty"`
	Allowance     *AllowanceField `json:"abono,omitempty"`
}

type ResponseField struct {
	Message  string `json:"texto"`
	DateTime string `json:"datahora"`
}

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

type Query struct {
	Total int              `json:"total"`
	Data  []*PunchResponse `json:"data"`
}
