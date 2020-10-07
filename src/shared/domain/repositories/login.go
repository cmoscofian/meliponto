package repositories

import "github.com/cmoscofian/meliponto/src/shared/domain/entities"

// The LoginClient type is an interface that defines
// the implementation of all methods on the communication
// layer related to login a user into the system.
type LoginClient interface {
	Login(*entities.Context, *entities.LoginRequest) (*entities.LoginResponse, error)
}

// The LoginService type is an interface that defines
// the implementation of all methods on the application
// layer related to the authentication and session of a user.
type LoginService interface {
	HandleLogin(*entities.Context, string) (string, error)
}
