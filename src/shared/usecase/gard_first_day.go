package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/chain/domain"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/util"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

type gardFirstDay struct {
	parent domain.Usecase
}

func (u *gardFirstDay) Get(ctx *entities.Context, date time.Time) []*entities.PunchRequest {
	punches := make([]*entities.PunchRequest, 0)

	for _, h := range ctx.Gard.Hours.Finish {
		punch := &entities.PunchRequest{
			Date: date.Format(constant.DateLayout),
			Allowance: &entities.AllowanceField{
				Begin:  h.Start,
				End:    h.End,
				Period: "especifico",
				Reason: "Sobreaviso",
			},
			Justification: "outros",
			Message:       ctx.Gard.Messages.Default,
			Punch:         nil,
			Type:          entities.AllowancePunch,
		}

		punches = append(punches, punch)
	}

	if util.IsWeekday(date) && !util.IsHoliday(date) {
		parent := u.parent.Get(ctx, date)
		punches = append(punches, parent...)
	}

	return punches
}
