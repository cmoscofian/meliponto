package builder

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/chain"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
)

type chainContext struct {
	Context *entities.Context
	Date    time.Time
	Forced  bool
}

func (c *chainContext) GetContext() *entities.Context {
	return c.Context
}

func (c *chainContext) GetDate() time.Time {
	return c.Date
}

func (c *chainContext) GetForced() bool {
	return c.Forced
}

// The ChainContextBuilder type is an implementation
// of the builder pattern to ease the construction
// of the context to be passed on the chain.
type ChainContextBuilder struct {
	ctx chainContext
}

// SetContext is a setter method to encapsulate
// the logic for setting the Context entity on the
// chain's context.
// It returns a pointer to the ChainContextBuilder type.
func (c *ChainContextBuilder) SetContext(ctx *entities.Context) *ChainContextBuilder {
	c.ctx.Context = ctx
	return c
}

// SetDate is a setter method to encapsulate
// the logic for setting the "date" to be used
// on the chains evaluation.
// It returns a pointer to the ChainContextBuilder type.
func (c *ChainContextBuilder) SetDate(date time.Time) *ChainContextBuilder {
	c.ctx.Date = date
	return c
}

// SetForced is a setter method to encapsulate
// the logic for setting the "forced" field to be used
// on the chains evaluation.
// It returns a pointer to the ChainContextBuilder type.
func (c *ChainContextBuilder) SetForced(forced bool) *ChainContextBuilder {
	c.ctx.Forced = forced
	return c
}

// Builder returns an implementation of the
// Context to be used on the chains evaluation.
func (c *ChainContextBuilder) Builder() chain.Context {
	return &c.ctx
}
