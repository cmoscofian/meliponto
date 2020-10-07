package rest

import (
	"fmt"

	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/domain/repositories"
	"github.com/cmoscofian/meliponto/src/shared/infrastructure/restclient"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

type deletePunch struct {
	client restclient.Client
}

// NewDeletePunch returns an implementation of the
// DeletePunch repository given a non-nil rest
// client.
func NewDeletePunch(c restclient.Client) repositories.DeletePunchClient {
	if c == nil {
		panic(constant.ClientError)
	}

	return &deletePunch{c}
}

// DeletePunchByID is responsible for building the request for deleting a punch by it's ID
// given a valid context and making the delete request to the punch URI.
// It returns a pointer to a PunchResponse entity and an error.
func (d *deletePunch) DeletePunchByID(token, punchID string) (*entities.PunchResponse, error) {
	headers := map[string]string{
		"cookie": fmt.Sprintf("qwert-external=%s", token),
	}
	uri := fmt.Sprintf("%s/%s", constant.SitePunchURI, punchID)

	res := new(entities.PunchResponse)
	if err := d.client.Delete(uri, headers, res); err != nil {
		return nil, err
	}

	return res, nil
}
