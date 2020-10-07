package service

import (
	"fmt"
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/domain/repositories"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

type punch struct {
	ls  repositories.LoginService
	cpc repositories.CreatePunchClient
}

// NewPunch returns an implementation of the
// CreatePunch service given a valid Login service and
// a CreatePunch client implementation of the communication layer.
func NewPunch(l repositories.LoginService, c repositories.CreatePunchClient) repositories.CreatePunchService {
	if l == nil {
		panic(fmt.Sprintf(constant.ServiceError, "login"))
	}

	if c == nil {
		panic(fmt.Sprintf(constant.ClientError, "create-punch"))
	}

	return &punch{l, c}
}

func (p *punch) HandleCreateRegularPunch(ctx *entities.Context, date time.Time, bodys []*entities.PunchRequest) ([]*entities.PunchRequest, error) {
	return nil, nil
}

func (p *punch) HandleCreateGardPunch(ctx *entities.Context, date time.Time, start time.Time, end time.Time, bodys []*entities.PunchRequest) ([]*entities.PunchRequest, error) {
	return nil, nil
}

func (p *punch) HandleCreateOffTimePunch(ctx *entities.Context, date time.Time, time time.Time, bodys []*entities.PunchRequest) ([]*entities.PunchRequest, error) {
	return nil, nil
}
