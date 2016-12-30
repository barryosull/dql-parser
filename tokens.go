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
	namespaceObject = "namespaceObject"
	quotedName = "quotedName"
	usingDatabase = "using database"
	forDomain = "forDomain"
	inContext = "inContext"
	withinAggregate = "within aggregate"
	class = "class"
	classOpen = "<|"
	classClose = "|>"
	eof = "eof"

	//Keywords - actions
	create = "create"

	//DQL Keywords - Objects
	database = "database"
	domain = "domain"
	context = "context"
	aggregate = "aggregate"
	value = "value"
	event = "event"
	entity = "entity"
	command = "command"
	projection = "projection"
	invariant = "invariant"
	query = "query"

	// Class components
	properties = "properties"
	check = "check"
	handler = "handler"
	function = "function"
	whenEvent = "when event"


	// Command Handler statements
	assertInvariant = "assert invariant"
	not = "not"
	runQuery = "run query"
	applyEvent = "apply event"

	// Operators
	assign   = "="
	plus     = "+"
	minus    = "-"
	bang     = "!"
	asterisk = "*"
	slash    = "/"
	arrow 	 = "->"
	strongArrow = "=>"
	and 	 = "and"
	or 	 = "or"
	lt = "<"
	gt = ">"
	eq     = "=="
	not_eq = "!="

	// Delimiters
	comma     = ","
	semicolon = ";"
	colon     = ":"

	lparen   = "("
	rparen   = ")"
	lbrace   = "{"
	rbrace   = "}"
	lbracked = "["
	rbracket = "]"

	//Types
	number = "number"
	typeRef = "type reference"
	identifier = "identifier"
	boolean = "boolean"

	//Statements
	if_ = "if"
	elseIf = "else if"
	else_ = "else"
	return_ = "return"
	foreach = "foreach"
	as = "as"
)

func Apos(pos int) Token {
	return NewToken(semicolon, ";", pos);
}

func ClsOpen(pos int) Token {
	return NewToken(classOpen, "<|", pos);
}

func ClsClose(pos int) Token {
	return NewToken(classClose, "|>", pos);
}