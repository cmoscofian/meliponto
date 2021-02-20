package repository

import "github.com/cmoscofian/meliponto/src/shared/domain/entity"

// The DeletePunchClient type is an interface that defines
// the implementation of all methods on the communication
// layer related to deleting a punch from the system.
type DeletePunchClient interface {
	DeletePunchByID(token string, punchID string) (*entity.PunchResponse, error)
}

// The DeletePunchService type is an interface that defines
// the implementation of all methods required on the application
// layer related to deleting punches from the system.
type DeletePunchService interface {
	HandleDeletePunch(ctx *entity.Context, punchID string)
}
