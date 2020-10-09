package node

import (
	"github.com/cmoscofian/meliponto/src/shared/chain/domain"
	"github.com/cmoscofian/meliponto/src/shared/usecase"
	"github.com/cmoscofian/meliponto/src/shared/util"
)

type gardSunday struct {
	next domain.Node
}

// NewGardSunday returns a new node of the chain
// that evaluates whether or not the given request
// is for a sunday.
func NewGardSunday() domain.Node {
	return &gardSunday{}
}

func (n *gardSunday) SetNext(node domain.Node) {
	n.next = node
}

func (n *gardSunday) HasNext() bool {
	return n.next != nil
}

func (n *gardSunday) Evaluate(c domain.Context) domain.Usecase {
	if c.IsGard() && util.IsSunday(c.GetDate()) {
		return usecase.GardSunday
	}

	if n.HasNext() {
		return n.next.Evaluate(c)
	}

	return nil
}
