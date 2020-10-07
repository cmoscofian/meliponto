package chain

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
)

// The Context type is an interface that implements
// all methods required for evaluations on the
// "chain-of-responsability".
type Context interface {
	GetContext() *entities.Context
	GetDate() time.Time
	GetForced() bool
}

// The Node type is an interface that implements
// all methods required for creating a new
// node to be used on the evaluation of the
// "chain-of-responsability"
type Node interface {
	Evaluate(Context) (*entities.PunchRequest, error)
	SetNext(Node)
}
