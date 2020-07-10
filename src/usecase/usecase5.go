package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/context"
)

func NewOnGard(ctx *context.Configuration, date time.Time, start, end string) ([]byte, error) {
	uc := NewUsecase(ctx.Gard.Messages.Default, date, true)
	uc.SetAllowance(start, end)
	return uc.Create()
}
