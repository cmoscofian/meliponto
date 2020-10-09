package node

import (
	"github.com/cmoscofian/meliponto/src/shared/chain/domain"
	"github.com/cmoscofian/meliponto/src/shared/usecase"
	"github.com/cmoscofian/meliponto/src/shared/util"
)

type gardLastDay struct {
	next domain.Node
}

// NewGardLastDay returns a new node of the chain
// that evaluates whether or not the given request
// is the last day of gard.
func NewGardLastDay() domain.Node {
	return &gardLastDay{}
}

func (n *gardLastDay) SetNext(node domain.Node) {
	n.next = node
}

func (n *gardLastDay) HasNext() bool {
	return n.next != nil
}

func (n *gardLastDay) Evaluate(c domain.Context) domain.Usecase {
	if c.IsGard() && c.IsRange() && util.IsSameDay(c.GetDate(), c.GetStartDate()) {
		return usecase.GardLastDay
	}

	if n.HasNext() {
		return n.next.Evaluate(c)
	}

	return nil
}
