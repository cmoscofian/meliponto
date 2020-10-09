package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

type gardHoliday struct{}

func (u *gardHoliday) Get(ctx *entities.Context, date time.Time) []*entities.PunchRequest {
	punches := make([]*entities.PunchRequest, 0)

	for _, h := range ctx.Gard.Hours.Holiday {
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

	return punches
}
