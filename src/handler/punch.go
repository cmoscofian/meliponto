package handler

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/usecase"
	"github.com/cmoscofian/meliponto/src/shared/util"
)

// HandlePunch is responsible for creating evey UC regular punch body and optionally
// UC gard punch bodies. After creating the UC's it append it to the pointer of []byte provided.
func HandlePunch(ctx *entities.Context, day time.Time, bodys *[][]byte, isOnGard bool) error {
	if util.IsRegularDay(day) {
		wsBody, err := usecase.NewWorkStart(ctx, day)
		if err != nil {
			return err
		}

		lsBody, err := usecase.NewLunchStart(ctx, day)
		if err != nil {
			return err
		}

		leBody, err := usecase.NewLunchEnd(ctx, day)
		if err != nil {
			return err
		}

		weBody, err := usecase.NewWorkEnd(ctx, day)
		if err != nil {
			return err
		}

		*bodys = append(*bodys, wsBody)
		*bodys = append(*bodys, lsBody)
		*bodys = append(*bodys, leBody)
		*bodys = append(*bodys, weBody)
	}

	if isOnGard {
		return HandleOnGardPunch(ctx, false, day, day, day, bodys)
	}

	return nil
}

// HandleOnGardPunch is responsible for creating evey UC gard punch bodies. After creating the UC's it append it to the pointer of []byte provided.
func HandleOnGardPunch(ctx *entities.Context, isFull bool, day, start, end time.Time, bodys *[][]byte) error {
	hours := getGardHoursFromContext(ctx, isFull, day, start, end)
	for _, h := range hours {
		body, err := usecase.NewOnGard(ctx, day, h.Start, h.End)
		if err != nil {
			return err
		}

		*bodys = append(*bodys, body)
	}

	return nil
}

func getGardHoursFromContext(ctx *entities.Context, isFull bool, day, start, end time.Time) []*entities.GardFieldHoursRange {
	if util.IsSaturday(day) {
		return ctx.Gard.Hours.Saturday
	}

	if util.IsSunday(day) {
		return ctx.Gard.Hours.Sunday
	}

	if util.IsWeekHoliday(day) {
		return ctx.Gard.Hours.Holiday
	}

	if isFull && util.IsSameDay(day, start) {
		return ctx.Gard.Hours.Begin
	}

	if isFull && util.IsSameDay(day, end) {
		return ctx.Gard.Hours.Finish
	}

	return ctx.Gard.Hours.Weekday
}
