package node

import (
	"github.com/cmoscofian/meliponto/src/shared/chain/domain"
	"github.com/cmoscofian/meliponto/src/shared/usecase"
	"github.com/cmoscofian/meliponto/src/shared/util"
)

type gardHoliday struct {
	next domain.Node
}

// NewGardHoliday returns a new node of the chain
// that evaluates whether or not the given request
// is for a holiday.
func NewGardHoliday() domain.Node {
	return &gardHoliday{}
}

func (n *gardHoliday) SetNext(node domain.Node) {
	n.next = node
}

func (n *gardHoliday) HasNext() bool {
	return n.next != nil
}

func (n *gardHoliday) Evaluate(c domain.Context) domain.Usecase {
	if c.IsGard() && util.IsHoliday(c.GetDate()) {
		return usecase.GardHoliday
	}

	if n.HasNext() {
		return n.next.Evaluate(c)
	}

	return nil
}
