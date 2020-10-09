package node

import (
	"github.com/cmoscofian/meliponto/src/shared/chain/domain"
	"github.com/cmoscofian/meliponto/src/shared/usecase"
	"github.com/cmoscofian/meliponto/src/shared/util"
)

type gardSaturday struct {
	next domain.Node
}

// NewGardSaturday returns a new node of the chain
// that evaluates whether or not the given request
// is for a saturday.
func NewGardSaturday() domain.Node {
	return &gardSaturday{}
}

func (n *gardSaturday) SetNext(node domain.Node) {
	n.next = node
}

func (n *gardSaturday) HasNext() bool {
	return n.next != nil
}

func (n *gardSaturday) Evaluate(c domain.Context) domain.Usecase {
	if c.IsGard() && util.IsSaturday(c.GetDate()) {
		return usecase.GardSaturday
	}

	if n.HasNext() {
		return n.next.Evaluate(c)
	}

	return nil
}
