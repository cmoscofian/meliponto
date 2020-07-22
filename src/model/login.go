package model

// The LoginRequest implements the request body structure
// for login into from the Ahgora system.
type LoginRequest struct {
	Empresa   string `json:"empresa"`
	Matricula string `json:"matricula"`
	Senha     string `json:"senha"`
	Origin    string `json:"origin,omitempty"`
}

// The StatusType is a custom type to represent every
// possible valid status from the login response status field.
type StatusType string

// Implementation of every possible valid for StatusType
const (
	SuccessStatus StatusType = "success"
	ErrorStatus   StatusType = "error"
)

// The LoginResponse implements the response body structure
// when login into the Ahgora system.
type LoginResponse struct {
	Status     StatusType `json:"r"`
	Token      string     `json:"jwt"`
	Message    string     `json:"text"`
	Validation bool       `json:"captcha"`
}
