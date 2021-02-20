package service

import (
	"fmt"
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
	"github.com/cmoscofian/meliponto/src/shared/domain/repository"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

type fetch struct {
	fc repository.FetchPunchClient
}

// NewFetch returns an implementation of the
// FetchPunch service given a valid Login service implementation
// and a FetchPunch client on the communication layer.
func NewFetch(fc repository.FetchPunchClient) repository.FetchPunchService {
	if fc == nil {
		panic(fmt.Sprintf(constant.ClientError, "fetch-punch"))
	}

	return &fetch{fc}
}

func (f fetch) HandleFetchSingle(ctx *entity.Context, punchID string, token string) (*entity.PunchResponse, error) {
	query, err := f.fc.GetPunchByID(token, punchID)
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (f fetch) HandleFetchRange(ctx *entity.Context, start, end time.Time, token string) (*entity.QueryPunchResponse, error) {
	query, err := f.fc.GetPunchByDateRange(token, start, end)
	if err != nil {
		return nil, err
	}

	return query, nil
}
