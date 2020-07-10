package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/context"
)

func NewLunchStart(ctx *context.Configuration, date time.Time) ([]byte, error) {
	uc := NewUsecase(ctx.Default.Messages.LunchStart, date, false)
	uc.SetTime(ctx.Default.Hours.LunchStart)
	return uc.Create()
}
