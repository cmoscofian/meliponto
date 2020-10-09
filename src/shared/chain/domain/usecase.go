package domain

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
)

// The Usecase type is an interface that implements
// all methods required for creating a new usecase
// to be returned by a node on the chain.
type Usecase interface {
	Get(*entities.Context, time.Time) []*entities.PunchRequest
}
