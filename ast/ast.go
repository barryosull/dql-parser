package ast

type Ast struct {
	Kind string
}

type Exp interface {
	e()
}

type ExpBlock struct {
	Ast
	Exps []Exp
}

type NullExp struct {
	Ast
}

type TrueLiteral struct {
	Ast
	Value bool
}

type FalseLiteral struct {
	Ast
	Value bool
}

type Return struct {
	Ast
	Value Exp
}

type If struct {
	Ast
	Check      Exp
	Consequent Exp
	Alternate  Exp
}

func (a Ast) e() {}




