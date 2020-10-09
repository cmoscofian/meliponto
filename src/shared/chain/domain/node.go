package domain

// The Node type is an interface that implements
// all methods required for creating a new
// node to be used on the evaluation of the
// "chain-of-responsability"
type Node interface {
	Evaluate(Context) Usecase
	HasNext() bool
	SetNext(Node)
}
