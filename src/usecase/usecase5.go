package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/context"
)

// NewOnGard returns a []byte and an error given a valid context and date.
// It implements the NewOnGard usecase based on the default configuration field,
func NewOnGard(ctx *context.Configuration, date time.Time, start, end string) ([]byte, error) {
	uc := NewUsecase(ctx.Gard.Messages.Default, date, true)
	uc.SetAllowance(start, end)
	return uc.Create()
}
