package ast

import (
	"encoding/json"
	"log"
	"fmt"
	"reflect"
)

func (n Node) MarshalJSON() ([]byte, error) {
	kind := reflect.TypeOf(n.Exp).String();
	return json.Marshal(&struct {
		Kind string
		Exp  Exp
	}{
		Kind: kind,
		Exp: n.Exp,
	})
}

type genericAst struct {
	Kind string
	Exp json.RawMessage
}

func (n *Node) UnmarshalJSON(data []byte) error {

	gAst := new(genericAst);
	err := json.Unmarshal(data, gAst);
	if err != nil {
		log.Fatalln("error:", err)
	}

	var dst interface{} = gAst.exp();

	if (dst == nil) {
		fmt.Println(string(data));
		log.Fatalln("error: Unknown Kind of", gAst.Kind)
	}

	json.Unmarshal(gAst.Exp, dst)
	n.Exp = dst.(Exp);
	return nil
}

var kindHandlers = map[string]func() interface{}{
	"ast.ExpBlock": func() interface{} { return new(ExpBlock) },
	"ast.NullExp": func() interface{} { return new(NullExp) },
	"ast.TrueLiteral": func() interface{} { return new(TrueLiteral) },
	"ast.FalseLiteral": func() interface{} { return new(FalseLiteral) },
	"ast.Return": func() interface{} { return new(Return) },
	"ast.If": func() interface{} { return new(If) },
}

func (a *genericAst) exp() interface{} {
	return kindHandlers[a.Kind]();
}
