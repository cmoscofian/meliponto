package util

import (
	"fmt"
	"strings"

	"github.com/cmoscofian/meliponto/src/shared/domain/entity"
)

// FormatCSVMessage returns a slice of bytes and an error
// based on a valid context and a pointer to a punch query
// response entity.
// It converts the punches query into a valid format that
// will be converted into a report used for automatic validation.
func FormatCSVMessage(ctx *entity.Context, punches []*entity.PunchResponse) ([]byte, error) {
	buff := &strings.Builder{}
	for _, p := range punches {
		if _, err := buff.WriteString(fmt.Sprintf("%s,%s,%s\r\n", p.ID, p.Date, ctx.UserID)); err != nil {
			return nil, err
		}
	}

	return []byte(buff.String()), nil
}
