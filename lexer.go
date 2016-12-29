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
	for state := lexToken; state != nil; {
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

func (l *lexer) matchingPrefix (prefixes []string) (string, bool) {
	for _, prefix := range prefixes {
		if (l.hasNextPrefix(prefix)) {
			return prefix, true
		}
	}
	return "", false
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
	l.skipWS();
	if (l.peek() == EOF) {
		return nil;
	}
	if (l.hasNextPrefix(create)) {
		return lexCreate
	}
	if (l.hasNextPrefix(value)) {
		return lexClass
	}
	if (l.hasNextPrefix(event)) {
		return lexClass
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
	if (l.hasNextPrefix("in")) {
		return lexInContext
	}
	if (l.hasNextPrefix("within")) {
		return lexWithinAggregate
	}
	if (l.hasNextPrefix(apostrophe)) {
		return lexApostrophe
	}
	if (l.hasNextPrefix(classOpen)) {
		return lexClassOpener
	}
	if (l.hasNextPrefix(classClose)) {
		return lexClassCloser
	}
	if (l.hasNextPrefix(namespaceBlockOpen)) {
		return lexNamespaceBlockOpener
	}
	if (l.hasNextPrefix(namespaceBlockClose)) {
		return lexNamespaceBlockCloser
	}

	return l.err();
}

func lexCreate(l *lexer) stateFn {
	l.pos += len(create)
	l.emit(create)
	return lexNSObjectType
}


func lexNSObjectType(l *lexer) stateFn {
	l.skipWS()
	typ, match := l.matchingPrefix([]string{database, domain, context, aggregate})
	if (!match) {
		return l.err()
	}
	l.pos += len(typ)
	l.emit(namespaceObject)
	return lexToken
}

func lexNSObjectName(l *lexer) stateFn {
	l.lexQuotedStringAsToken(quotedName)
	return lexToken
}

func lexUsingDatabase(l *lexer) stateFn {
	l.skipStr("using")
	l.skipStr(database)

	if l.lexQuotedStringAsToken(usingDatabase) {
		return lexToken
	}
	return nil
}

func lexForDomain (l *lexer) stateFn {
	l.skipStr("for")
	l.skipStr(domain)

	if l.lexQuotedStringAsToken(forDomain) {
		return lexToken
	}
	return nil
}

func lexInContext (l *lexer) stateFn {
	l.skipStr("in")
	l.skipStr(context)

	if l.lexQuotedStringAsToken(inContext) {
		return lexToken
	}
	return nil
}

func lexWithinAggregate(l *lexer) stateFn {
	l.skipStr("within")
	l.skipStr(aggregate)

	if l.lexQuotedStringAsToken(withinAggregate) {
		return lexToken
	}
	return nil
}


func lexApostrophe(l *lexer) stateFn {
	l.next();
	l.emit(apostrophe);
	return lexToken;
}

func lexClassOpener(l *lexer) stateFn {
	l.pos += len(classOpen)
	l.emit(classOpen)
	return lexToken
}

func lexClassCloser(l *lexer) stateFn {
	l.pos += len(classClose)
	l.emit(classClose)
	return lexToken
}

func lexNamespaceBlockOpener(l *lexer) stateFn {
	l.pos += len(namespaceBlockOpen)
	l.emit(namespaceBlockOpen)
	return lexToken
}

func lexNamespaceBlockCloser(l *lexer) stateFn {
	l.pos += len(namespaceBlockClose)
	l.emit(namespaceBlockClose)
	return lexToken
}

func lexClass(l *lexer) stateFn {
	l.pos += len(class)
	l.emit(class)
	return lexToken
}

func (l *lexer) skipWS() {
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

func (l *lexer) lexQuotedStringAsToken(tokenType TokenType) bool {
	l.skipWS()
	if (l.next() == '\'') {
		l.ignore();
		for {
			if (l.peek() == '\'') {
				l.emit(tokenType)
				l.skip()
				return true;
			}
			l.next();
		}
	}
	return false;
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

func (l *lexer) skip() {
	l.next()
	l.ignore()
}

func (l *lexer) skipStr(string string) {
	l.pos += len(string)
	l.skipWS()
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

func (l *lexer) err() stateFn {
	format := "There was a problem near: %q"
	errToken := Token{err, fmt.Sprintf(format, l.parsed()), l.start}
	l.error = &errToken
	return nil
}