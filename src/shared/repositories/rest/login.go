package rest

import (
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/domain/repositories"
	"github.com/cmoscofian/meliponto/src/shared/infrastructure/restclient"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

type login struct {
	client restclient.Client
}

// NewLogin returns an implementation of the
// Login repository with a the rest client
// passed.
func NewLogin(c restclient.Client) repositories.LoginClient {
	if c == nil {
		panic(constant.ClientError)
	}

	return &login{c}
}

// Login is responsible for building the request for the login given a valid context
// and request making the post request to the login URI.
// It returns a pointer to a LoginResponse entity and an error.
func (l *login) Login(ctx *entities.Context, req *entities.LoginRequest) (*entities.LoginResponse, error) {
	res := new(entities.LoginResponse)
	if err := l.client.Post(constant.SiteLoginURI, nil, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
