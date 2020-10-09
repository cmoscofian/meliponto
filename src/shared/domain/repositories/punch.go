package repositories

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
)

// The FetchPunchClient type is an interface that defines
// the implementation of all methods on the communication
// layer related to fetching punches from the system.
type FetchPunchClient interface {
	GetPunchByID(token string, punchID string) (*entities.PunchResponse, error)
	GetPunchByDateRange(string, time.Time, time.Time) (*entities.QueryPunchResponse, error)
}

// The CreatePunchClient type is an interface that defines
// the implementation of all methods on the communication
// layer related to creating a new punch in the system.
type CreatePunchClient interface {
	CreatePunch(string, *entities.PunchRequest) (*entities.PunchResponse, error)
}

// The DeletePunchClient type is an interface that defines
// the implementation of all methods on the communication
// layer related to deleting a punch from the system.
type DeletePunchClient interface {
	DeletePunchByID(token string, punchID string) (*entities.PunchResponse, error)
}

// The FetchPunchService type is an interface that defines
// the implementation of all methods required on the application
// layer related to fetching punches from the system.
type FetchPunchService interface {
	HandleFetchRange(token string, start, end time.Time) (*entities.QueryPunchResponse, error)
}

// The CreatePunchService type is an interface that defines
// the implementation of all methods required on the application
// layer related to creating punches into the system.
type CreatePunchService interface {
	// HandleCreatePunch(*entities.Context, *time.Time, *time.Time) error
	// HandleCreateRegularPunch(*entities.Context, time.Time, []*entities.PunchRequest) []*entities.PunchRequest
	// HandleCreateGardPunch(*entities.Context, time.Time, *time.Time, *time.Time, []*entities.PunchRequest) []*entities.PunchRequest
	// HandleCreateOffTimePunch(*entities.Context, time.Time, time.Time, []*entities.PunchRequest) []*entities.PunchRequest
}
