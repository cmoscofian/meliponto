package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/context"
)

// NewLunchEnd returns a []byte and an error given a valid context and date.
// It implements the LunchEnd usecase based on the default configuration field.
func NewLunchEnd(ctx *context.Configuration, date time.Time) ([]byte, error) {
	uc := NewUsecase(ctx.Default.Messages.LunchEnd, date, false)
	uc.SetTime(ctx.Default.Hours.LunchEnd)
	return uc.Create()
}
