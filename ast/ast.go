package ast

type Exp interface {
	e()
}

type Ast struct {
	Exp Exp
}

type ExpBlock struct {
	Exps []Ast
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
	Value Ast
}

type If struct {
	Check      Ast
	Consequent Ast
	Alternate  Ast
}

func (a ExpBlock) e() {}
func (a NullExp) e() {}
func (a TrueLiteral) e() {}
func (a FalseLiteral) e() {}
func (a Return) e() {}
func (a If) e() {}





