package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/context"
)

// NewWorkStart returns a []byte and an error given a valid context and date.
// It implements the WorkStart usecase based on the default configuration field.
func NewWorkStart(ctx *context.Configuration, date time.Time) ([]byte, error) {
	uc := NewUsecase(ctx.Default.Messages.WorkStart, date, false)
	uc.SetTime(ctx.Default.Hours.WorkStart)
	return uc.Create()
}
