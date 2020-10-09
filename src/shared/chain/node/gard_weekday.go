package node

import (
	"github.com/cmoscofian/meliponto/src/shared/chain/domain"
	"github.com/cmoscofian/meliponto/src/shared/usecase"
	"github.com/cmoscofian/meliponto/src/shared/util"
)

type gardWeekday struct {
	next domain.Node
}

// NewGardWeekday returns a new node of the chain
// that evaluates whether or not the given request
// is for regular weekday.
func NewGardWeekday() domain.Node {
	return &gardWeekday{}
}

func (n *gardWeekday) SetNext(node domain.Node) {
	n.next = node
}

func (n *gardWeekday) HasNext() bool {
	return n.next != nil
}

func (n *gardWeekday) Evaluate(c domain.Context) domain.Usecase {
	date := c.GetDate()
	if c.IsGard() && util.IsWeekday(date) && !util.IsWeekHoliday(date) {
		return usecase.GardWeekday
	}

	if n.HasNext() {
		return n.next.Evaluate(c)
	}

	return nil
}
