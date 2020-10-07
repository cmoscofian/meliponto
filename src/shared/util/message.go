package util

import (
	"time"

	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
)

// GetDefaultMessage returns a valid string message. If message provided
// is empty it will check the datetime and return an adequate WorkStart or WorkEnd
// message.
func GetDefaultMessage(ctx *entities.Context, message string, datetime time.Time) string {
	if message == "" {
		if datetime.Hour() <= 12 {
			return ctx.Default.Messages.WorkStart
		}

		return ctx.Default.Messages.WorkEnd
	}

	return message
}
