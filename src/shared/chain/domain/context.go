package domain

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
	GetStartDate() time.Time
	GetEndDate() time.Time
	IsGard() bool
	IsRange() bool
	ShouldForce() bool
}

// The ContextBuilder type is an interface that implements
// all methods required for building the context to be passed
// on to the "chain-of-responsability".
type ContextBuilder interface {
	SetContext(*entities.Context) ContextBuilder
	SetDate(time.Time) ContextBuilder
	SetStartDate(*time.Time) ContextBuilder
	SetEndDate(*time.Time) ContextBuilder
	SetForced(bool) ContextBuilder
	SetGard(bool) ContextBuilder
	Build() Context
}
