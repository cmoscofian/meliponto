package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

type regular struct{}

func (u *regular) Get(ctx *entities.Context, date time.Time) []*entities.PunchRequest {
	punches := make([]*entities.PunchRequest, 0)

	workStart := &entities.PunchRequest{
		Date:          date.Format(constant.DateLayout),
		Allowance:     nil,
		Justification: "outros",
		Message:       ctx.Default.Messages.WorkStart,
		Punch: &entities.PunchField{
			Time: ctx.Default.Hours.WorkStart,
		},
		Type: entities.RegularPunch,
	}

	lunchStart := &entities.PunchRequest{
		Date:          date.Format(constant.DateLayout),
		Allowance:     nil,
		Justification: "outros",
		Message:       ctx.Default.Messages.LunchStart,
		Punch: &entities.PunchField{
			Time: ctx.Default.Hours.LunchStart,
		},
		Type: entities.RegularPunch,
	}

	lunchEnd := &entities.PunchRequest{
		Date:          date.Format(constant.DateLayout),
		Allowance:     nil,
		Justification: "outros",
		Message:       ctx.Default.Messages.LunchEnd,
		Punch: &entities.PunchField{
			Time: ctx.Default.Hours.LunchEnd,
		},
		Type: entities.RegularPunch,
	}

	workEnd := &entities.PunchRequest{
		Date:          date.Format(constant.DateLayout),
		Allowance:     nil,
		Justification: "outros",
		Message:       ctx.Default.Messages.WorkEnd,
		Punch: &entities.PunchField{
			Time: ctx.Default.Hours.WorkEnd,
		},
		Type: entities.RegularPunch,
	}

	punches = append(punches, workStart, lunchStart, lunchEnd, workEnd)

	return punches
}
