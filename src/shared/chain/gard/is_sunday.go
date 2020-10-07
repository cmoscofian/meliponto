package gard

import (
	"errors"

	"github.com/cmoscofian/meliponto/src/shared/chain"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/util"
)

type isSunday struct {
	Next chain.Node
}

// NewIsSunday returns a new node of the chain
// that evaluates whether or not the given request
// is for a sunday.
func NewIsSunday() chain.Node {
	return &isSunday{}
}

func (i *isSunday) SetNext(n chain.Node) {
	i.Next = n
}

func (i *isSunday) Evaluate(c chain.Context) (*entities.PunchRequest, error) {
	date := c.GetDate()
	if util.IsSunday(date) {
		return &entities.PunchRequest{}, nil
	}

	if i.Next != nil {
		return i.Next.Evaluate(c)
	}

	return nil, errors.New("Should not punch")
}
