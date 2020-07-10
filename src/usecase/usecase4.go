package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/context"
)

func NewWorkEnd(ctx *context.Configuration, date time.Time) ([]byte, error) {
	uc := NewUsecase(ctx.Default.Messages.WorkEnd, date, false)
	uc.SetTime(ctx.Default.Hours.WorkEnd)
	return uc.Create()
}
