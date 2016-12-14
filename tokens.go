package parser

import (
	"sync"
	"fmt"
)

type Token struct {
	typ TokenType
	val string
}

func NewToken(typ TokenType, val string) *Token {
	return &Token{typ, val};
}

func (t *Token) Compare(o *Token) bool{
	return t.typ == o.typ && t.val == o.val;
}

func (i *Token) String() string {
	switch i.typ {
	case EOF:
		return "EOF"
	case Error:
		return i.val
	}
	if len(i.val) > 10 {
		return fmt.Sprintf("%.10q...", i.val)
	}
	return fmt.Sprintf("%q", i.val)
}

type TokenType int

const (
	Error TokenType = iota
	Create
	NamespaceObject
	QuotedName
	UsingDatabase
	ForDomain
	InContext
	WithinAggregate

	Class
	ClassOpen
	ClassClose
	Apostrophe
	EOF
)

func Apos() *Token {
	return NewToken(Apostrophe, ";");
}

func ClsOpen() *Token {
	return NewToken(ClassOpen, "<|");
}

func ClsClose() *Token {
	return NewToken(ClassClose, "|>");
}

type TokenList struct {
	tokens []*Token
}

func (t *TokenList) Append(token *Token) {
	t.tokens = append(t.tokens, token);
}

func (t *TokenList) All() []*Token {
	res := t.tokens;
	t.tokens = make([]*Token, 0);
	return res;
}

var instance *TokenList
var once sync.Once

func GetInstanceTokenList() *TokenList {
	once.Do(func() {
		instance = &TokenList{}
	})
	return instance
}
