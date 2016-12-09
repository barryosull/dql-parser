package ast

type Exp interface {
	e()
}

type Node struct {
	Exp Exp
}

type ExpBlock struct {
	Exps []Node
}

type NullExp struct {

}

type TrueLiteral struct {
	Value bool
}

type FalseLiteral struct {
	Value bool
}

type Return struct {
	Value Node
}

type If struct {
	Check      Node
	Consequent Node
	Alternate  Node
}

func (a ExpBlock) e() {}
func (a NullExp) e() {}
func (a TrueLiteral) e() {}
func (a FalseLiteral) e() {}
func (a Return) e() {}
func (a If) e() {}





