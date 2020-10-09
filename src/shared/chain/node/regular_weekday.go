package node

import (
	"github.com/cmoscofian/meliponto/src/shared/chain/domain"
	"github.com/cmoscofian/meliponto/src/shared/usecase"
	"github.com/cmoscofian/meliponto/src/shared/util"
)

type regularWeekday struct {
	next domain.Node
}

// NewRegularWeekday returns a new node of the chain
// that evaluates whether or not the given request
// is for regular weekday.
func NewRegularWeekday() domain.Node {
	return &regularWeekday{}
}

func (n *regularWeekday) SetNext(node domain.Node) {
	n.next = node
}

func (n *regularWeekday) HasNext() bool {
	return n.next != nil
}

func (n *regularWeekday) Evaluate(c domain.Context) domain.Usecase {
	date := c.GetDate()
	if !c.IsGard() && util.IsWeekday(date) && !util.IsWeekHoliday(date) {
		return usecase.Regular
	}

	if n.HasNext() {
		return n.next.Evaluate(c)
	}

	return nil
}
