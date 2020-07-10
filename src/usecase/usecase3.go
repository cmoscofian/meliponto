package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/context"
)

func NewLunchEnd(ctx *context.Configuration, date time.Time) ([]byte, error) {
	uc := NewUsecase(ctx.Default.Messages.LunchEnd, date, false)
	uc.SetTime(ctx.Default.Hours.LunchEnd)
	return uc.Create()
}
