package ast

type Exp interface {
	e()
}

type TrueLiteral struct {
	value bool
}

type FalseLiteral struct {
	value bool
}

type If struct {
	check Exp
	consequent Exp
	alternate Exp
}

func NewIf(check Exp, consequent Exp, alternate Exp) If{
	return If {
		check: check,
		consequent: consequent,
		alternate: alternate,
	};
}

func NewTrueLiteral() TrueLiteral {
	return TrueLiteral {true};
}

func NewFalseLiteral() FalseLiteral {
	return FalseLiteral {false};
}

func (e TrueLiteral) e() {}
func (e FalseLiteral) e() {}
func (e If) e() {}


