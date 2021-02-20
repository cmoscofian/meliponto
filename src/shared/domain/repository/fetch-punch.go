package repository

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
)

// The FetchPunchClient type is an interface that defines
// the implementation of all methods on the communication
// layer related to fetching punches from the system.
type FetchPunchClient interface {
	GetPunchByID(token, punchID string) (*entity.PunchResponse, error)
	GetPunchByDateRange(token string, start, end time.Time) (*entity.QueryPunchResponse, error)
}

// The FetchPunchService type is an interface that defines
// the implementation of all methods required on the application
// layer related to fetching punches from the system.
type FetchPunchService interface {
	HandleFetchSingle(ctx *entity.Context, punchID string, password string) (*entity.PunchResponse, error)
	HandleFetchRange(ctx *entity.Context, start, end time.Time, password string) (*entity.QueryPunchResponse, error)
}
