package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

type custom struct{}

func (u *custom) Get(ctx *entities.Context, date time.Time) []*entities.PunchRequest {
	punches := make([]*entities.PunchRequest, 0)

	var message string = "Outros"

	if v, ok := ctx.Values["msg"]; ok {
		message = v
	}

	custom := &entities.PunchRequest{
		Date:          date.Format(constant.DateLayout),
		Allowance:     nil,
		Justification: "outros",
		Message:       message,
		Punch: &entities.PunchField{
			Time: date.Format(constant.TimeLayout),
		},
		Type: entities.RegularPunch,
	}

	punches = append(punches, custom)

	return punches
}
