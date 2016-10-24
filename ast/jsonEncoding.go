package ast

import (
	"encoding/json"
	"log"
	"fmt"
	"reflect"
)

type genericAst struct {
	Kind string
	Exp json.RawMessage
}

func (a Ast) MarshalJSON() ([]byte, error) {
	kind := reflect.TypeOf(a.Exp).String();
	return json.Marshal(&struct {
		Kind string
		Exp  Exp
	}{
		Kind: kind,
		Exp: a.Exp,
	})
}

func (a *Ast) UnmarshalJSON(data []byte) error {

	gAst := new(genericAst);
	err := json.Unmarshal(data, gAst);
	if err != nil {
		log.Fatalln("error:", err)
	}

	var dst interface{};
	switch gAst.Kind {
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
		log.Fatalln("error: Unknown Kind of", gAst.Kind)
	}

	json.Unmarshal(gAst.Exp, dst)
	a.Exp = dst.(Exp);
	return nil
}
