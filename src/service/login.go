package service

import (
	"encoding/json"

	"github.com/cmoscofian/meliponto/src/context"
	"github.com/cmoscofian/meliponto/src/model"
	"github.com/cmoscofian/meliponto/src/util"
	"github.com/cmoscofian/meliponto/src/util/constants"
)

// Login is responsible for building the request for the login given a valid context
// and making the post request to the login URI.
// It communicates with all the other sytems via channels ([]byte channel and error channel)
func Login(ctx *context.Configuration, chbs chan<- []byte, cher chan<- error) {
	password, err := util.GetPassword()
	if err != nil {
		cher <- err
		return
	}

	loginBody := &model.LoginRequest{
		Empresa:   ctx.CompanyID,
		Matricula: ctx.UserID,
		Origin:    "portal",
		Senha:     string(password),
	}

	jsonBody, err := json.Marshal(loginBody)
	if err != nil {
		cher <- err
		return
	}

	Post(constants.SiteLoginURI, nil, jsonBody, chbs, cher)
}
