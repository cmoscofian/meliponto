package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

type custom struct{}

func (u *custom) Get(ctx *entity.Context, date time.Time) []*entity.PunchRequest {
	punches := make([]*entity.PunchRequest, 0)

	var message string = "Outros"

	if v, ok := ctx.Values["msg"]; ok {
		message = v
	}

	custom := &entity.PunchRequest{
		Date:          date.Format(constant.DateLayout),
		Allowance:     nil,
		Justification: "outros",
		Message:       message,
		Punch: &entity.PunchField{
			Time: date.Format(constant.TimeLayout),
		},
		Type: entity.RegularPunch,
	}

	punches = append(punches, custom)

	return punches
}
