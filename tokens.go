package parser

import (
	"sync"
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

func (t *Token) String() string {
	return t.val;
}

type TokenType int

const (
	Create TokenType = iota
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