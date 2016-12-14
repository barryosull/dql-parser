package parser

import (
	"strings"
	"unicode/utf8"
)


// lexer holds the state of the scanner.
type lexer struct {
	name  string    // used only for error reports.
	input string    // the string being scanned.
	start int       // start position of this item.
	pos   int       // current position in the input.
	width int       // width of last rune read from input.
	tokens []Token
}

type stateFn func(*lexer) stateFn

func lex(name, input string) (*lexer) {
	l := &lexer{
		name:  name,
		input: input,
	}
	l.run()
	return l
}

func (l *lexer) run() {
	for state := lexCreate; state != nil; {
		state = state(l)
	}
}

func (l *lexer) emit(t TokenType) {
	append(l.tokens, Token{t, l.input[l.start:l.pos]});
	l.start = l.pos
}

const eof = -1

func lexCreate(l *lexer) stateFn {
	ignoreWS(l);
	l.pos += len(create)
	l.emit(Create)
	return lexNSObjectType
}

func lexNSObjectType(l *lexer) stateFn {
	ignoreWS(l);

	//Check for type

	l.pos += len(create)
	l.emit(Create)
	return lexNSObjectType
}

func ignoreWS(l *lexer) {
	for {
		switch r := l.next(); {
		case r == " " || r == '\n' || r == "\r" || r == "\t":
			l.ignore()
		default :
			break
		}
	}
}


// next returns the next rune in the input.
func (l *lexer) next() (rune int) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	rune, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return rune
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) peek() int {
	rune := l.next()
	l.backup()
	return rune
}

// accept consumes the next rune if it's from the valid set.
func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}
// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

// error returns an error token and terminates the scan
// by passing back a nil pointer that will be the next
// state, terminating l.run.
func (l *lexer) errorf(format string, args ...interface{})
	stateFn {
		l.items <- item{
		itemError,
		fmt.Sprintf(format, args...),
	}
	return nil
}

const (
	create = "create"

)