package model

type LoginRequest struct {
	Empresa   string `json:"empresa"`
	Matricula string `json:"matricula"`
	Senha     string `json:"senha"`
	Origin    string `json:"origin,omitempty"`
}

type StatusType string

const (
	SuccessStatus StatusType = "success"
	ErrorStatus   StatusType = "error"
)

type LoginResponse struct {
	Status     StatusType `json:"r"`
	Token      string     `json:"jwt"`
	Message    string     `json:"text"`
	Validation bool       `json:"captcha"`
}
