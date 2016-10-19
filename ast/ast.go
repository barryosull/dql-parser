package ast

import (

)

type Exp interface {
	e()
}

type ExpBlock struct {
	Exps []Exp
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
	Value Exp
}

type If struct {
	Check      Exp
	Consequent Exp
	Alternate  Exp
}

func NewTrueLiteral() TrueLiteral {
	return TrueLiteral{true}
}

func NewFalseLiteral() FalseLiteral {
	return FalseLiteral{false}
}

func (e TrueLiteral) e()  {}
func (e FalseLiteral) e() {}
func (e If) e()           {}
func (e Return) e()       {}
func (e NullExp) e()      {}
func (e ExpBlock) e()     {}


