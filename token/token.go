package token

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

const IgnoreTokenPos = -1

func (t *Token) Compare(o Token) bool{
	if (t.pos == IgnoreTokenPos || o.pos == IgnoreTokenPos) {
		return t.typ == o.typ && t.val == o.val;
	}
	return t.typ == o.typ && t.val == o.val && t.pos == o.pos;
}

func (i *Token) String() string {
	switch i.typ {
	case EOF:
		return "EOF"
	case ERR:
		return i.val
	}
	if len(i.val) > 10 {
		return fmt.Sprintf("char %q, %q, %.10q...", i.pos, i.typ, i.val)
	}
	return fmt.Sprintf("char %q, %q, %q", i.pos, i.typ, i.val)
}

type TokenType string

const (
	ERR TokenType = "err"

	NAMESPACEOBJECT = "namespaceObject"
	QUOTEDNAME = "quotedName"
	USINGDATABASE = "using database"
	FORDOMAIN = "forDomain"
	INCONTEXT = "inContext"
	WITHINAGGREGATE = "within aggregate"
	CLASS = "class"
	CLASSOPEN = "<|"
	CLASSCLOSE = "|>"
	EOF = "eof"

	//Keywords - actions
	CREATE = "create"

	//DQL Keywords - Objects
	DATABASE = "database"
	DOMAIN = "domain"
	CONTEXT = "context"
	AGGREGATE = "aggregate"
	VALUE = "value"
	EVENT = "event"
	ENTITY = "entity"
	COMMAND = "command"
	PROJECTION = "projection"
	INVARIANT = "invariant"
	QUERY = "query"

	// Class components
	PROPERTIES = "properties"
	CHECK = "check"
	HANDLER = "handler"
	FUNCTION = "function"
	WHENEVENT = "when event"

	// Command Handler statements
	ASSERTINVARIANT = "assert invariant"
	NOT = "not"
	RUNQUERY = "run query"
	APPLYEVENT = "apply event"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	ARROW 	 = "->"
	STRONGARROW = "=>"
	AND 	 = "and"
	OR 	 = "or"
	LT = "<"
	GT = ">"
	EQ     = "=="
	NOTEQ = "!="
	LTOREQ = "<="
	GTOREQ = ">="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	//Types
	INTEGER = "integer"
	FLOAT = "float"
	BOOLEAN = "boolean"
	STRING = "string"
	TYPEREF = "type reference"

	IDENTIFIER = "identifier"

	//Statements
	IF = "if"
	ELSEIF = "else if"
	ELSE = "else"
	RETURN = "return"
	FOREACH = "foreach"
	AS = "as"
)

func Semicolon(pos int) Token {
	return NewToken(SEMICOLON, ";", pos);
}

func ClsOpen(pos int) Token {
	return NewToken(CLASSOPEN, "<|", pos);
}

func ClsClose(pos int) Token {
	return NewToken(CLASSCLOSE, "|>", pos);
}