package service

import (
	"errors"
	"fmt"

	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/domain/repositories"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

type login struct {
	lc repositories.LoginClient
}

// NewLogin returns an implementation of the
// Login service given a valid LoginClient implementation
// of the communication layer.
func NewLogin(lc repositories.LoginClient) repositories.LoginService {
	if lc == nil {
		panic(fmt.Sprintf(constant.ClientError, "login"))
	}

	return &login{lc}
}

// HandleLogin returns a string token and an error. It is
// responsible for all the application layer logic
// regarding authentication given a valid
// context and password.
func (l *login) HandleLogin(ctx *entities.Context, password string) (string, error) {
	req := &entities.LoginRequest{
		Empresa:   ctx.CompanyID,
		Matricula: ctx.UserID,
		Origin:    "portal",
		Senha:     password,
	}

	resp, err := l.lc.Login(ctx, req)
	if err != nil {
		return "", err
	}

	if resp.IsSuccess() {
		return resp.GetToken(), nil
	}

	if resp.HasMessage() {
		return "", errors.New(resp.GetMessage())
	}

	return "", errors.New(constant.InvalidLoginError)
}
