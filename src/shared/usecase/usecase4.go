package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
)

// NewWorkEnd returns a []byte and an error given a valid context and date.
// It implements the WorkEnd usecase based on the default configuration field.
func NewWorkEnd(ctx *entities.Context, date time.Time) ([]byte, error) {
	uc := NewUsecase(ctx.Default.Messages.WorkEnd, date, false)
	uc.SetTime(ctx.Default.Hours.WorkEnd)
	return uc.Create()
}
