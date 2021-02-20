package rest

import (
	"fmt"

	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
	"github.com/cmoscofian/meliponto/src/shared/domain/repository"
	"github.com/cmoscofian/meliponto/src/shared/infrastructure/restclient"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

type createPunch struct {
	client restclient.Client
}

// NewCreatePunch returns an implementation of the
// FetchPunch repository with a the rest client
// passed.
func NewCreatePunch(c restclient.Client) repository.CreatePunchClient {
	if c == nil {
		panic(constant.ClientError)
	}

	return &createPunch{c}
}

// CreatePunch is responsible for building the request for the punch given a valid context
// and body and making the post request to the login URI.
// It returns a pointer to a PunchResponse entity and an error.
func (c createPunch) CreatePunch(token string, req *entity.PunchRequest) (*entity.PunchResponse, error) {
	headers := map[string]string{
		"cookie": fmt.Sprintf("qwert-external=%s", token),
	}

	res := new(entity.PunchResponse)
	if err := c.client.Post(constant.SitePunchURI, headers, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
