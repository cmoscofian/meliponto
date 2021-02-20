package service

import (
	"errors"
	"fmt"

	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
	"github.com/cmoscofian/meliponto/src/shared/domain/repository"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

type login struct {
	lc repository.LoginClient
}

// NewLogin returns an implementation of the
// Login service given a valid LoginClient implementation
// of the communication layer.
func NewLogin(lc repository.LoginClient) repository.LoginService {
	if lc == nil {
		panic(fmt.Sprintf(constant.ClientError, "login"))
	}

	return &login{lc}
}

// HandleLogin returns a string token and an error. It is
// responsible for all the application layer logic
// regarding authentication given a valid
// context and password.
func (l login) HandleLogin(ctx *entity.Context, password []byte) (string, error) {
	req := &entity.LoginRequest{
		Empresa:   ctx.CompanyID,
		Matricula: ctx.UserID,
		Origin:    "portal",
		Senha:     string(password),
	}

	resp, err := l.lc.Login(req)
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
