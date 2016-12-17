package parser

import (
	"strings"
	"unicode/utf8"
	"fmt"
)

// lexer holds the state of the scanner.
type lexer struct {
	name  string    // used only for error reports.
	input string    // the string being scanned.
	start int       // start position of this item.
	pos   int       // current position in the input.
	width int       // width of last rune read from input.
	tokens []Token
	error *Token
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
	return
}

func (l *lexer) emit(t TokenType) {
	l.tokens = append(l.tokens, Token{t, l.input[l.start:l.pos], l.start});
	l.start = l.pos
}

func (l *lexer) hasNextPrefix (prefix string) bool {
	var unlexed = l.input[l.pos:];
	return strings.HasPrefix(unlexed, prefix);
}

func (l *lexer) parsed () string {
	start := 0;
	if (l.pos - 40 > 0) {
		start = l.pos - 40
	}
	return l.input[start:l.pos];
}

const EOF = -1

func lexToken(l *lexer) stateFn {
	ignoreWS(l);
	if (l.hasNextPrefix("create")) {
		return lexCreate
	}
	if (l.hasNextPrefix("database")) {
		return lexNSObjectType
	}
	if (l.hasNextPrefix("domain")) {
		return lexNSObjectType
	}
	if (l.hasNextPrefix("context")) {
		return lexNSObjectType
	}
	if (l.hasNextPrefix("'")) {
		return lexNSObjectName
	}
	if (l.hasNextPrefix("using")) {
		return lexUsingDatabase
	}
	if (l.hasNextPrefix("for")) {
		return lexForDomain
	}
	if (l.hasNextPrefix(";")) {
		return lexApostrophe
	}

	return l.errorf("There was a problem near: %q", l.parsed());

	return nil
}

func lexCreate(l *lexer) stateFn {
	l.pos += len(create)
	l.emit(create)
	return lexNSObjectType
}

func lexNSObjectType(l *lexer) stateFn {
	if (l.hasNextPrefix("database")) {
		l.pos += len("database")
		l.emit(namespaceObject)
	}
	if (l.hasNextPrefix("domain")) {
		l.pos += len("domain")
		l.emit(namespaceObject)
	}
	if (l.hasNextPrefix("context")) {
		l.pos += len("context")
		l.emit(namespaceObject)
	}
	return lexToken
}

func lexNSObjectName(l *lexer) stateFn {
	if (l.next() == '\'') {
		l.ignore();
		for {
			r := l.next();
			if (r == '\'') {
				l.backup()
				l.emit(quotedName)
				l.next()
				l.ignore()
				return lexToken
			}
		}
	}
	return lexToken
}

func lexUsingDatabase(l *lexer) stateFn {
	l.pos += len("using")
	ignoreWS(l);

	l.pos += len("database")
	ignoreWS(l);

	if (l.next() == '\'') {
		l.ignore();
		for {
			r := l.next();
			if (r == '\'') {
				l.backup()
				l.emit(usingDatabase)
				l.next()
				l.ignore()
				return lexToken
			}
		}
	}
	return nil
}

func lexForDomain (l *lexer) stateFn {
	l.pos += len("for")
	ignoreWS(l);

	l.pos += len("domain")
	ignoreWS(l);

	if (l.next() == '\'') {
		l.ignore();
		for {
			r := l.next();
			if (r == '\'') {
				l.backup()
				l.emit(forDomain)
				l.next()
				l.ignore()
				return lexToken
			}
		}
	}
	return nil
}

func lexApostrophe(l *lexer) stateFn {
	l.next();
	l.emit(apostrophe);
	return nil;
}

func ignoreWS(l *lexer) {
	for {
		switch r := l.next(); {
		case r == ' ' || r == '\n' || r == '\r' || r == '\t':
			l.ignore()
		default :
			l.backup()
			return
		}
	}
}


// next returns the next rune in the input.
func (l *lexer) next() (rn int) {
	if l.pos >= len(l.input) {
		l.width = 0
		return EOF
	}
	var r rune;
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return int(r);
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) peek() int {
	r := l.next()
	l.backup()
	return r
}

// accept consumes the next rune if it's from the valid set.
func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, rune(l.next())) >= 0 {
		return true
	}
	l.backup()
	return false
}
// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, rune(l.next())) >= 0 {
	}
	l.backup()
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	errToken := Token{err, fmt.Sprintf(format, args...), l.start}
	l.error = &errToken
	return nil
}