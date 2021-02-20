package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

type regular struct{}

func (u *regular) Get(ctx *entity.Context, date time.Time) []*entity.PunchRequest {
	punches := make([]*entity.PunchRequest, 0)

	workStart := &entity.PunchRequest{
		Date:          date.Format(constant.DateLayout),
		Allowance:     nil,
		Justification: "outros",
		Message:       ctx.Default.Messages.WorkStart,
		Punch: &entity.PunchField{
			Time: ctx.Default.Hours.WorkStart,
		},
		Type: entity.RegularPunch,
	}

	lunchStart := &entity.PunchRequest{
		Date:          date.Format(constant.DateLayout),
		Allowance:     nil,
		Justification: "outros",
		Message:       ctx.Default.Messages.LunchStart,
		Punch: &entity.PunchField{
			Time: ctx.Default.Hours.LunchStart,
		},
		Type: entity.RegularPunch,
	}

	lunchEnd := &entity.PunchRequest{
		Date:          date.Format(constant.DateLayout),
		Allowance:     nil,
		Justification: "outros",
		Message:       ctx.Default.Messages.LunchEnd,
		Punch: &entity.PunchField{
			Time: ctx.Default.Hours.LunchEnd,
		},
		Type: entity.RegularPunch,
	}

	workEnd := &entity.PunchRequest{
		Date:          date.Format(constant.DateLayout),
		Allowance:     nil,
		Justification: "outros",
		Message:       ctx.Default.Messages.WorkEnd,
		Punch: &entity.PunchField{
			Time: ctx.Default.Hours.WorkEnd,
		},
		Type: entity.RegularPunch,
	}

	punches = append(punches, workStart, lunchStart, lunchEnd, workEnd)

	return punches
}
