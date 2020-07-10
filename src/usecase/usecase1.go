package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/context"
)

func NewWorkStart(ctx *context.Configuration, date time.Time) ([]byte, error) {
	uc := NewUsecase(ctx.Default.Messages.WorkStart, date, false)
	uc.SetTime(ctx.Default.Hours.WorkStart)
	return uc.Create()
}
