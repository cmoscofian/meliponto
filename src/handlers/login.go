package handlers

import (
	"encoding/json"
	"errors"

	"github.com/cmoscofian/meliponto/src/context"
	"github.com/cmoscofian/meliponto/src/model"
	"github.com/cmoscofian/meliponto/src/service"
	"github.com/cmoscofian/meliponto/src/util/constants"
)

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

		if login.Message != "" {
			return login.Token, errors.New(login.Message)
		}

		return login.Token, errors.New(constants.InvalidLoginError)

	case err := <-cher:
		return "", err
	}
}
