package gard

import (
	"errors"

	"github.com/cmoscofian/meliponto/src/shared/chain"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
	"github.com/cmoscofian/meliponto/src/shared/util"
)

type isSaturday struct {
	Next chain.Node
}

// NewIsSaturday returns a new node of the chain
// that evaluates whether or not the given request
// is for a saturday.
func NewIsSaturday() chain.Node {
	return &isSaturday{}
}

func (i *isSaturday) SetNext(n chain.Node) {
	i.Next = n
}

func (i *isSaturday) Evaluate(c chain.Context) (*entities.PunchRequest, error) {
	date := c.GetDate()
	if util.IsSaturday(date) {
		return &entities.PunchRequest{}, nil
	}

	if i.Next != nil {
		return i.Next.Evaluate(c)
	}

	return nil, errors.New("Should not punch")
}
