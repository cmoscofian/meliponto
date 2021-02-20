package rest

import (
	"fmt"
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
	"github.com/cmoscofian/meliponto/src/shared/domain/repository"
	"github.com/cmoscofian/meliponto/src/shared/infrastructure/restclient"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

type fetchPunch struct {
	client restclient.Client
}

// NewFetchPunch returns an implementation of the
// FetchPunch repository with a the rest client
// passed.
func NewFetchPunch(c restclient.Client) repository.FetchPunchClient {
	if c == nil {
		panic(constant.ClientError)
	}

	return &fetchPunch{c}
}

// GetPunchByID is responsible for building the request for fetching a single punch by it's
// ID given a valid context and making the get request to the punch URI.
// It returns a pointer to a PunchResponse entity and an error.
func (f fetchPunch) GetPunchByID(token, punchID string) (*entity.PunchResponse, error) {
	uri := fmt.Sprintf("%s/%s", constant.SitePunchURI, punchID)
	headers := map[string]string{
		"cookie": fmt.Sprintf("qwert-external=%s", token),
	}

	resp := new(entity.PunchResponse)
	if err := f.client.Get(uri, headers, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// GetPunchByDateRange is responsible for building the request for fetching a range of punches
// given a valid context, start and end date and making the get request to the punch URI.
// It returns a pointer to a QueryPunchResponse and an error.
// PS: To get a single day request just use the same date for start and end.
func (f fetchPunch) GetPunchByDateRange(token string, start, end time.Time) (*entity.QueryPunchResponse, error) {
	headers := map[string]string{
		"cookie": fmt.Sprintf("qwert-external=%s", token),
	}
	uri := fmt.Sprintf("%s?inicio=%s&fim=%s", constant.SitePunchURI, start.Format(constant.DateLayout), end.Format(constant.DateLayout))

	resp := new(entity.QueryPunchResponse)
	if err := f.client.Get(uri, headers, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
