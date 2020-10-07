package gard

import (
	"errors"

	"github.com/cmoscofian/meliponto/src/shared/chain"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/util"
)

type isHoliday struct {
	Next chain.Node
}

// NewIsHoliday returns a new node of the chain
// that evaluates whether or not the given request
// is for a holiday.
func NewIsHoliday() chain.Node {
	return &isHoliday{}
}

func (i *isHoliday) SetNext(n chain.Node) {
	i.Next = n
}

func (i *isHoliday) Evaluate(c chain.Context) (*entities.PunchRequest, error) {
	date := c.GetDate()
	if util.IsHoliday(date) {
		return &entities.PunchRequest{}, nil
	}

	if i.Next != nil {
		return i.Next.Evaluate(c)
	}

	return nil, errors.New("Should not punch")
}
