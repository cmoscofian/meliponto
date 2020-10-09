package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/chain/domain"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/util"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

type gardWeekday struct {
	parent domain.Usecase
}

func (u *gardWeekday) Get(ctx *entities.Context, date time.Time) []*entities.PunchRequest {
	punches := make([]*entities.PunchRequest, 0)

	for _, h := range ctx.Gard.Hours.Weekday {
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

	if util.IsWeekday(date) {
		parent := u.parent.Get(ctx, date)
		punches = append(punches, parent...)
	}

	return punches
}
