package service

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/chain/builder"
	"github.com/cmoscofian/meliponto/src/shared/chain/domain"
	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
	"github.com/cmoscofian/meliponto/src/shared/domain/repository"
)

type usecases struct {
	n domain.Node
}

// NewUsecase returns an implementation of the
// GenerateBodys service given a valid node.
func NewUsecase(n domain.Node) repository.GenerateBodys {
	return &usecases{n}
}

func (u usecases) HandleRegularBody(ctx *entity.Context, date time.Time) []*entity.PunchRequest {
	chainBuilder := builder.NewChainContext()
	context := chainBuilder.
		SetContext(ctx).
		SetDate(date).
		SetEndDate(nil).
		SetForced(false).
		SetGard(false).
		SetStartDate(nil).
		Build()

	return u.n.Evaluate(context).Get(ctx, date)
}

func (u usecases) HandleGardBody(ctx *entity.Context, date time.Time, start, end *time.Time) []*entity.PunchRequest {
	chainBuilder := builder.NewChainContext()
	context := chainBuilder.
		SetContext(ctx).
		SetDate(date).
		SetEndDate(end).
		SetForced(false).
		SetGard(true).
		SetStartDate(start).
		Build()

	return u.n.Evaluate(context).Get(ctx, date)
}

func (u usecases) HandleOffTimeBody(ctx *entity.Context, date, hour time.Time) []*entity.PunchRequest {
	chainBuilder := builder.NewChainContext()
	context := chainBuilder.
		SetContext(ctx).
		SetDate(date).
		SetEndDate(nil).
		SetForced(true).
		SetGard(false).
		SetStartDate(nil).
		Build()

	return u.n.Evaluate(context).Get(ctx, date)
}
