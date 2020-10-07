package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
)

// NewLunchStart returns a []byte and an error given a valid context and date.
// It implements the LunchStart usecase based on the default configuration field.
func NewLunchStart(ctx *entities.Context, date time.Time) ([]byte, error) {
	uc := NewUsecase(ctx.Default.Messages.LunchStart, date, false)
	uc.SetTime(ctx.Default.Hours.LunchStart)
	return uc.Create()
}
