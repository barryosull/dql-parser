package ast

import (
	"reflect"
	"encoding/json"
	"fmt"
	"log"
)

type Exp interface {
	e()
}

type Ast struct {
	Kind string
	Exp Exp
}

func NewAst(e Exp) Ast {
	ast := Ast{};
	ast.Kind = reflect.TypeOf(e).String();
	ast.Exp = e;
	return ast;
}

type JsonAst struct {
	Kind string
	Exp json.RawMessage
}

func (a *Ast) UnmarshalJSON(data []byte) error {

	//Turn into generic Ast
	jsonAst := new(JsonAst);
	err := json.Unmarshal(data, jsonAst);
	if err != nil {
		log.Fatalln("error:", err)
	}

	var dst interface{};
	switch jsonAst.Kind {
	case "ast.ExpBlock":
		dst = new(ExpBlock)
	case "ast.NullExp":
		dst = new(NullExp)
	case "ast.TrueLiteral":
		dst = new(TrueLiteral)
	case "ast.FalseLiteral":
		dst = new(FalseLiteral)
	case "ast.Return":
		dst = new(Return)
	case "ast.If":
		dst = new(If)
	}

	if (dst == nil) {
		fmt.Println(string(data));
		log.Fatalln("error: Unknown Kind of", jsonAst.Kind)
	}

	json.Unmarshal(jsonAst.Exp, dst)

	a.Kind = jsonAst.Kind;
	a.Exp = dst.(Exp);
	return nil
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





