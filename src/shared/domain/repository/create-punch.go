package repository

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
)

// The CreatePunchClient type is an interface that defines
// the implementation of all methods on the communication
// layer related to creating a new punch in the system.
type CreatePunchClient interface {
	CreatePunch(string, *entity.PunchRequest) (*entity.PunchResponse, error)
}

// The GenerateBodys type is an interface that defines
// the implementation of all methods required on the application
// layer related to creating punches into the system.
type GenerateBodys interface {
	HandleRegularBody(*entity.Context, time.Time) []*entity.PunchRequest
	HandleGardBody(*entity.Context, time.Time, *time.Time, *time.Time) []*entity.PunchRequest
	HandleOffTimeBody(*entity.Context, time.Time, time.Time) []*entity.PunchRequest
}
