package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

type gardWeekday struct{}

func (u *gardWeekday) Get(ctx *entity.Context, date time.Time) []*entity.PunchRequest {
	punches := make([]*entity.PunchRequest, 0)

	for _, h := range ctx.Gard.Hours.Weekday {
		punch := &entity.PunchRequest{
			Date: date.Format(constant.DateLayout),
			Allowance: &entity.AllowanceField{
				Begin:  h.Start,
				End:    h.End,
				Period: "especifico",
				Reason: "Sobreaviso",
			},
			Justification: "outros",
			Message:       ctx.Gard.Messages.Default,
			Punch:         nil,
			Type:          entity.AllowancePunch,
		}

		punches = append(punches, punch)
	}

	return punches
}
