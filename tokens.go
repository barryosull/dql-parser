package parser

import (
	"fmt"
)

type Token struct {
	typ TokenType
	val string
}

func NewToken(typ TokenType, val string) Token {
	return Token{typ, val};
}

func (t *Token) Compare(o Token) bool{
	return t.typ == o.typ && t.val == o.val;
}

func (i *Token) String() string {
	switch i.typ {
	case eof:
		return "EOF"
	case error:
		return i.val
	}
	if len(i.val) > 10 {
		return fmt.Sprintf("%q, %.10q...", i.typ, i.val)
	}
	return fmt.Sprintf("%q, %q", i.typ, i.val)
}

type TokenType string

const (
	error TokenType = "error"
	create = "create"
	namespaceObject = "namespaceObject"
	quotedName = "quotedName"
	usingDatabase = "usingDatabase"
	forDomain = "forDomain"
	inContext = "inContext"
	withinAggregate = "withinAggregate"
	class = "class"
	classOpen = "classOpen"
	classClose = "classClose"
	apostrophe = "apostrophe"
	eof = "eof"
)

func Apos() Token {
	return NewToken(apostrophe, ";");
}

func ClsOpen() Token {
	return NewToken(classOpen, "<|");
}

func ClsClose() Token {
	return NewToken(classClose, "|>");
}