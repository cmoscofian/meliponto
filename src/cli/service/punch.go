package service

import (
	"fmt"
	"time"

	"github.com/cmoscofian/meliponto/src/shared/chain/builder"
	"github.com/cmoscofian/meliponto/src/shared/chain/domain"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/domain/repositories"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

type punch struct {
	ls  repositories.LoginService
	cpc repositories.CreatePunchClient
	n   domain.Node
}

// NewPunch returns an implementation of the
// CreatePunch service given a valid Login service and
// a CreatePunch client implementation of the communication layer.
func NewPunch(l repositories.LoginService, c repositories.CreatePunchClient, n domain.Node) repositories.CreatePunchService {
	if l == nil {
		panic(fmt.Sprintf(constant.ServiceError, "login"))
	}

	if c == nil {
		panic(fmt.Sprintf(constant.ClientError, "create-punch"))
	}

	return &punch{l, c, n}
}

func (p *punch) HandleCreateRegularPunch(ctx *entities.Context, date time.Time, bodys []*entities.PunchRequest) []*entities.PunchRequest {
	chainBuilder := builder.NewChainContext()
	context := chainBuilder.
		SetContext(ctx).
		SetDate(date).
		SetEndDate(nil).
		SetForced(false).
		SetGard(false).
		SetStartDate(nil).
		Build()

	return p.n.Evaluate(context).Get(ctx, date)
}

func (p *punch) HandleCreateGardPunch(ctx *entities.Context, date time.Time, start, end *time.Time, bodys []*entities.PunchRequest) []*entities.PunchRequest {
	chainBuilder := builder.NewChainContext()
	context := chainBuilder.
		SetContext(ctx).
		SetDate(date).
		SetEndDate(end).
		SetForced(false).
		SetGard(true).
		SetStartDate(start).
		Build()

	return p.n.Evaluate(context).Get(ctx, date)
}

func (p *punch) HandleCreateOffTimePunch(ctx *entities.Context, date, hour time.Time, bodys []*entities.PunchRequest) []*entities.PunchRequest {
	chainBuilder := builder.NewChainContext()
	context := chainBuilder.
		SetContext(ctx).
		SetDate(date).
		SetEndDate(nil).
		SetForced(true).
		SetGard(true).
		SetStartDate(nil).
		Build()

	return p.n.Evaluate(context).Get(ctx, date)
}
