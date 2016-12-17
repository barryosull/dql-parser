package parser

import (
	"fmt"
)

type Token struct {
	typ TokenType
	val string
	pos int
}

func NewToken(typ TokenType, val string, pos int) Token {
	return Token{typ, val, pos};
}

const ignoreTokenPos = 1

func (t *Token) Compare(o Token) bool{
	if (t.pos == ignoreTokenPos || o.pos == ignoreTokenPos) {
		return t.typ == o.typ && t.val == o.val;
	}
	return t.typ == o.typ && t.val == o.val && t.pos == o.pos;
}

func (i *Token) String() string {
	switch i.typ {
	case eof:
		return "EOF"
	case err:
		return i.val
	}
	if len(i.val) > 10 {
		return fmt.Sprintf("char %q, %q, %.10q...", i.pos, i.typ, i.val)
	}
	return fmt.Sprintf("char %q, %q, %q", i.pos, i.typ, i.val)
}

type TokenType string

const (
	err TokenType = "err"
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

func Apos(pos int) Token {
	return NewToken(apostrophe, ";", pos);
}

func Err(e string, pos int) *Token {
	t := NewToken(err, e, pos);
	return &t
}

func ClsOpen(pos int) Token {
	return NewToken(classOpen, "<|", pos);
}

func ClsClose(pos int) Token {
	return NewToken(classClose, "|>", pos);
}