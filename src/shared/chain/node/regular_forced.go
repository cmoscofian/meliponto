package node

import (
	"github.com/cmoscofian/meliponto/src/shared/chain/domain"
	"github.com/cmoscofian/meliponto/src/shared/usecase"
)

type regularForced struct {
	next domain.Node
}

// NewRegularForced returns a new node of the chain
// that evaluates whether or not the given request
// should be forced.
func NewRegularForced() domain.Node {
	return &regularForced{}
}

func (n *regularForced) SetNext(node domain.Node) {
	n.next = node
}

func (n *regularForced) HasNext() bool {
	return n.next != nil
}

func (n *regularForced) Evaluate(c domain.Context) domain.Usecase {
	if !c.IsGard() && c.ShouldForce() {
		return usecase.Custom
	}

	if n.HasNext() {
		return n.next.Evaluate(c)
	}

	return nil
}
