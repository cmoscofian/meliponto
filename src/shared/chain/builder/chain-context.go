package builder

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/chain/domain"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
)

type chainContext struct {
	context   *entities.Context
	date      time.Time
	startDate *time.Time
	endDate   *time.Time
	gard      bool
	isRange   bool
	force     bool
}

func (c *chainContext) GetContext() *entities.Context {
	return c.context
}

func (c *chainContext) IsRange() bool {
	return c.isRange
}

func (c *chainContext) GetDate() time.Time {
	return c.date
}

func (c *chainContext) GetStartDate() time.Time {
	return *c.startDate
}

func (c *chainContext) GetEndDate() time.Time {
	return *c.endDate
}

func (c *chainContext) ShouldForce() bool {
	return c.force
}

func (c *chainContext) IsGard() bool {
	return c.gard
}

// The chainContextBuilder type is an implementation
// of the builder pattern to ease the construction
// of the context to be passed on the chain.
type chainContextBuilder struct {
	ctx chainContext
}

// NewChainContext returns a pointer to a
// a chain context builder. That encapsulates
// and ensures no outside interference on the
// context creation.
func NewChainContext() domain.ContextBuilder {
	return &chainContextBuilder{}
}

// SetContext is a setter method to encapsulate
// the logic for setting the Context entity on the
// chain's context.
// It returns a pointer to the ChainContextBuilder type.
func (c *chainContextBuilder) SetContext(ctx *entities.Context) domain.ContextBuilder {
	c.ctx.context = ctx
	return c
}

// SetDate is a setter method to encapsulate
// the logic for setting the "date" to be used
// on the chains evaluation.
// It returns a pointer to the ChainContextBuilder type.
func (c *chainContextBuilder) SetDate(date time.Time) domain.ContextBuilder {
	c.ctx.date = date
	return c
}

// SetStartDate is a setter method to encapsulate
// the logic for setting the "startDate" to be used
// on the chains evaluation.
// It returns a pointer to the ChainContextBuilder type.
func (c *chainContextBuilder) SetStartDate(startDate *time.Time) domain.ContextBuilder {
	c.ctx.startDate = startDate
	return c
}

// SetEndDate is a setter method to encapsulate
// the logic for setting the "startDate" to be used
// on the chains evaluation.
// It returns a pointer to the ChainContextBuilder type.
func (c *chainContextBuilder) SetEndDate(endDate *time.Time) domain.ContextBuilder {
	c.ctx.endDate = endDate
	return c
}

// SetForced is a setter method to encapsulate
// the logic for setting the "force" field to be used
// on the chains evaluation.
// It returns a pointer to the ChainContextBuilder type.
func (c *chainContextBuilder) SetForced(forced bool) domain.ContextBuilder {
	c.ctx.force = forced
	return c
}

// SetGard is a setter method to encapsulate
// the logic for setting the "gard" field to be used
// on the chains evaluation.
// It returns a pointer to the ChainContextBuilder type.
func (c *chainContextBuilder) SetGard(gard bool) domain.ContextBuilder {
	c.ctx.gard = gard
	return c
}

// Build returns an implementation of the
// Context to be used on the chains evaluation.
func (c *chainContextBuilder) Build() domain.Context {
	if c.ctx.startDate != nil && c.ctx.endDate != nil {
		c.ctx.isRange = true
	}

	return &c.ctx
}
