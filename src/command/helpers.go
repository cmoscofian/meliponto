package command

import (
	"time"

	"github.com/cmoscofian/meliponto/src/context"
	"github.com/cmoscofian/meliponto/src/usecase"
	"github.com/cmoscofian/meliponto/src/util"
)

func dailyCheck(ctx *context.Configuration, day time.Time, bodys *[][]byte, isOnGard bool) error {
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

		if isOnGard {
			for _, v := range ctx.Gard.Hours.Weekday {
				body, err := usecase.NewOnGard(ctx, day, v.Start, v.End)
				if err != nil {
					return err
				}
				*bodys = append(*bodys, body)
			}
		}
	}

	if isOnGard {
		if util.IsSaturday(day) {
			for _, v := range ctx.Gard.Hours.Saturday {
				body, err := usecase.NewOnGard(ctx, day, v.Start, v.End)
				if err != nil {
					return err
				}
				*bodys = append(*bodys, body)
			}

			return nil
		}

		if util.IsSunday(day) {
			for _, v := range ctx.Gard.Hours.Sunday {
				body, err := usecase.NewOnGard(ctx, day, v.Start, v.End)
				if err != nil {
					return err
				}
				*bodys = append(*bodys, body)
			}

			return nil
		}

		if util.IsWeekHoliday(day) {
			for _, v := range ctx.Gard.Hours.Holiday {
				body, err := usecase.NewOnGard(ctx, day, v.Start, v.End)
				if err != nil {
					return err
				}
				*bodys = append(*bodys, body)
			}

			return nil
		}
	}

	return nil
}

func dailyCheckOnGard(ctx *context.Configuration, day, start, end time.Time, bodys *[][]byte) error {
	if util.IsSameDay(day, start) {
		for _, v := range ctx.Gard.Hours.Begin {
			body, err := usecase.NewOnGard(ctx, day, v.Start, v.End)
			if err != nil {
				return err
			}
			*bodys = append(*bodys, body)
		}

		return nil
	}

	if util.IsSameDay(day, end) {
		for _, v := range ctx.Gard.Hours.Finish {
			body, err := usecase.NewOnGard(ctx, day, v.Start, v.End)
			if err != nil {
				return err
			}
			*bodys = append(*bodys, body)
		}

		return nil
	}

	if util.IsRegularDay(day) {
		for _, v := range ctx.Gard.Hours.Weekday {
			body, err := usecase.NewOnGard(ctx, day, v.Start, v.End)
			if err != nil {
				return err
			}
			*bodys = append(*bodys, body)
		}

		return nil
	}

	if util.IsWeekHoliday(day) {
		for _, v := range ctx.Gard.Hours.Holiday {
			body, err := usecase.NewOnGard(ctx, day, v.Start, v.End)
			if err != nil {
				return err
			}
			*bodys = append(*bodys, body)
		}

		return nil
	}

	if util.IsSaturday(day) {
		for _, v := range ctx.Gard.Hours.Saturday {
			body, err := usecase.NewOnGard(ctx, day, v.Start, v.End)
			if err != nil {
				return err
			}
			*bodys = append(*bodys, body)
		}

		return nil
	}

	if util.IsSunday(day) {
		for _, v := range ctx.Gard.Hours.Sunday {
			body, err := usecase.NewOnGard(ctx, day, v.Start, v.End)
			if err != nil {
				return err
			}
			*bodys = append(*bodys, body)
		}

		return nil
	}

	return nil
}
