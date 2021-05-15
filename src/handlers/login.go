package handlers

import (
	"encoding/json"
	"errors"

	"github.com/cmoscofian/meliponto/src/context"
	"github.com/cmoscofian/meliponto/src/model"
	"github.com/cmoscofian/meliponto/src/service"
	"github.com/cmoscofian/meliponto/src/util/constants"
)

// HandleLogin returns a valid JWT token and an error provided a valid context.
// It communicates with all the other sytems via channels ([]byte channel and error channel)
func HandleLogin(ctx *context.Configuration, chbs chan []byte, cher chan error) (string, error) {
	go service.Login(ctx, chbs, cher)

	select {
	case response := <-chbs:
		var login model.LoginResponse
		if err := json.Unmarshal(response, &login); err != nil {
			return login.Token, err
		}

		if login.Status == model.SuccessStatus {
			return login.Token, nil
		}

		if login.Validation {
			return login.Token, errors.New("Captcha is required. You must login on browser first")
		}

		if login.Message != "" {
			return login.Token, errors.New(login.Message)
		}

		return login.Token, errors.New(constants.InvalidLoginError)

	case err := <-cher:
		return "", err
	}
}
