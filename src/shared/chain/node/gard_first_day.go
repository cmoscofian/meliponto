package node

import (
	"github.com/cmoscofian/meliponto/src/shared/chain/domain"
	"github.com/cmoscofian/meliponto/src/shared/usecase"
	"github.com/cmoscofian/meliponto/src/shared/util"
)

type gardFirstDay struct {
	next domain.Node
}

// NewGardFirstDay returns a new node of the chain
// that evaluates whether or not the given request
// is the first day of gard.
func NewGardFirstDay() domain.Node {
	return &gardFirstDay{}
}

func (n *gardFirstDay) SetNext(node domain.Node) {
	n.next = node
}

func (n *gardFirstDay) HasNext() bool {
	return n.next != nil
}

func (n *gardFirstDay) Evaluate(c domain.Context) domain.Usecase {
	if c.IsGard() && c.IsRange() && util.IsSameDay(c.GetDate(), c.GetStartDate()) {
		return usecase.GardFirstDay
	}

	if n.HasNext() {
		return n.next.Evaluate(c)
	}

	return nil
}
