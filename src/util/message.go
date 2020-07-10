package util

import (
	"time"

	"github.com/cmoscofian/meliponto/src/context"
)

func GetDefaultMessage(ctx *context.Configuration, message string, datetime time.Time) string {
	if message == "" {
		if datetime.Hour() <= 12 {
			return ctx.Default.Messages.WorkStart
		}

		return ctx.Default.Messages.WorkEnd
	}

	return message
}
