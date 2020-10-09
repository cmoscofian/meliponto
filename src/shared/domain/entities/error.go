package entities

// The ErrorResponse implements the default error
// response body structure from the Ahgora system.
type ErrorResponse struct {
	Error   bool   `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
