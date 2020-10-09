package entities

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

// IsSuccess return a bool validating whether
// or not the response "r" field is successfull.
func (l *LoginResponse) IsSuccess() bool {
	return l.Status == SuccessStatus
}

// HasMessage return a bool validating whether
// or not there is a valid message on the
// "text" field.
func (l *LoginResponse) HasMessage() bool {
	return l.Message != ""
}

// GetToken is a getter method to encapsulate
// the logic for retrieving a "jwt" field token.
func (l *LoginResponse) GetToken() string {
	return l.Token
}

// GetMessage is a getter method to encapsulate
// the logic for retrieving the "text" field value.
func (l *LoginResponse) GetMessage() string {
	return l.Message
}
