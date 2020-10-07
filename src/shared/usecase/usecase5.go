package usecase

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
)

// NewOnGard returns a []byte and an error given a valid context and date.
// It implements the NewOnGard usecase based on the default configuration field,
func NewOnGard(ctx *entities.Context, date time.Time, start, end string) ([]byte, error) {
	uc := NewUsecase(ctx.Gard.Messages.Default, date, true)
	uc.SetAllowance(start, end)
	return uc.Create()
}
