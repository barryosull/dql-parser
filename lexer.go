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

func (l *lexer) isNextPrefix(prefix string) bool {
	var unlexed = l.input[l.pos:];
	return strings.HasPrefix(unlexed, prefix);
}

//Check of the prefix matches and is not followed immediately by another identifier character
func (l *lexer) isNextKeyword(prefix string) bool {
	hasPrefix := l.isNextPrefix(prefix)
	if (!hasPrefix) {
		return false
	}
	nextRune, _ := l.runeAtPos(l.pos + len(prefix));
	if (isDigit(nextRune) || isLetter(nextRune)) {
		return false
	}
	return true
}

func (l *lexer) matchingPrefix (prefixes []string) (string, bool) {
	for _, prefix := range prefixes {
		if (l.isNextPrefix(prefix)) {
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

var whitespace = []int{' ', '\n', '\r', '\t'}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (l *lexer) skipWS() {
	for {
		switch r := l.next(); {
		case contains(whitespace, r):
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

func (l *lexer) lexAsToken(tokenType TokenType) stateFn {
	l.pos += len(tokenType)
	l.emit(tokenType)
	return lexToken
}


// next returns the next rune in the input.
func (l *lexer) next() int {
	r, width := l.runeAtPos(l.pos)
	if r == EOF {
		l.width = 0
		return r
	}
	l.width = width
	l.pos += l.width
	return r
}

func (l *lexer) runeAtPos(pos int) (rn int, width int) {
	if pos >= len(l.input) {
		rn = EOF
		width = 0
		return
	}
	var r rune;
	r, width = utf8.DecodeRuneInString(l.input[pos:])
	rn = int(r)
	return
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

//Scan until next space or non identifier character
func (l *lexer) scanIdentifier() {
	for {
		if (!isLetter(l.peek()) && !isDigit(l.peek())) {
			break;
		}
		l.next();
	}
}

const EOF = -1

func lexToken(l *lexer) stateFn {
	l.skipWS();

	if (l.peek() == EOF) {
		return nil;
	}

	//Have special lexing rules
	if (l.isNextKeyword(create)) {
		return lexCreate
	}
	if (l.isNextPrefix("'")) {
		return lexNSObjectName
	}
	if (l.isNextPrefix("\"")) {
		return lexString
	}
	if (l.isNextKeyword("using")) {
		return lexUsingDatabase
	}
	if (l.isNextKeyword("for")) {
		return lexForDomain
	}
	if (l.isNextKeyword("in")) {
		return lexInContext
	}
	if (l.isNextKeyword("within")) {
		return lexWithinAggregate
	}
	if (l.isNextPrefix(value+"\\") || l.isNextPrefix(event+"\\")) {
		return lexTypeRef
	}
	if (l.isNextKeyword("assert")) {
		return lexAssertInvariant
	}
	if (l.isNextKeyword("run")) {
		return lexRunQuery
	}
	if (l.isNextKeyword("apply")) {
		return lexApplyEvent
	}
	if (l.isNextKeyword("when")) {
		return lexWhenEvent
	}
	if (l.isNextKeyword(and)) {
		return l.lexAsToken(and)
	}
	if (l.isNextKeyword(or)) {
		return l.lexAsToken(or)
	}

	// No special cases, just lex and move on
	if (l.isNextKeyword(properties)) {
		return l.lexAsToken(properties);
	}
	if (l.isNextKeyword(check)) {
		return l.lexAsToken(check)
	}
	if (l.isNextKeyword(handler)) {
		return l.lexAsToken(handler)
	}
	if (l.isNextKeyword(function)) {
		return l.lexAsToken(function)
	}
	if (l.isNextKeyword(if_)) {
		return l.lexAsToken(if_)
	}
	if (l.isNextKeyword(elseIf)) {
		return l.lexAsToken(elseIf)
	}
	if (l.isNextKeyword(else_)) {
		return l.lexAsToken(else_)
	}
	if (l.isNextKeyword(foreach)) {
		return l.lexAsToken(foreach)
	}
	if (l.isNextKeyword(return_)) {
		return l.lexAsToken(return_)
	}
	if (l.isNextKeyword(as)) {
		return l.lexAsToken(as)
	}

	if (l.isNextPrefix(strongArrow)) {
		return l.lexAsToken(strongArrow)
	}
	if (l.isNextPrefix(eq)) {
		return l.lexAsToken(eq)
	}
	if (l.isNextPrefix(semicolon)) {
		return l.lexAsToken(semicolon);
	}
	if (l.isNextPrefix(assign)) {
		return l.lexAsToken(assign);
	}
	if (l.isNextPrefix(classOpen)) {
		return lexClassOpen;
	}
	if (l.isNextPrefix(classClose)) {
		return l.lexAsToken(classClose);
	}
	if (l.isNextPrefix(colon)) {
		return l.lexAsToken(colon);
	}
	if (l.isNextPrefix(lbrace)) {
		return l.lexAsToken(lbrace);
	}
	if (l.isNextPrefix(rbrace)) {
		return l.lexAsToken(rbrace);
	}
	if (l.isNextPrefix(lparen)) {
		return l.lexAsToken(lparen);
	}
	if (l.isNextPrefix(rparen)) {
		return l.lexAsToken(rparen);
	}
	if (l.isNextPrefix(lbracked)) {
		return l.lexAsToken(lbracked);
	}
	if (l.isNextPrefix(rbracket)) {
		return l.lexAsToken(rbracket);
	}
	if (l.isNextPrefix(lparen)) {
		return l.lexAsToken(lparen)
	}
	if (l.isNextPrefix(rparen)) {
		return l.lexAsToken(rparen)
	}
	if (l.isNextPrefix(not_eq)) {
		return l.lexAsToken(not_eq)
	}
	if (l.isNextPrefix(comma)) {
		return l.lexAsToken(comma)
	}
	if (l.isNextPrefix(arrow)) {
		return l.lexAsToken(arrow)
	}
	if (l.isNextPrefix(plus)) {
		return l.lexAsToken(plus)
	}
	if (l.isNextPrefix(minus)) {
		return l.lexAsToken(minus)
	}
	if (l.isNextPrefix(bang)) {
		return l.lexAsToken(bang)
	}
	if (l.isNextPrefix(asterisk)) {
		return l.lexAsToken(asterisk)
	}
	if (l.isNextPrefix(slash)) {
		return l.lexAsToken(slash)
	}
	if (l.isNextPrefix(lt)) {
		return l.lexAsToken(lt)
	}
	if (l.isNextPrefix(gt)) {
		return l.lexAsToken(gt)
	}

	if (isDigit(l.peek())) {
		return lexNumber;
	}
	if (isLetter(l.peek())) {
		return lexIdentifier;
	}

	return l.err();
}

func lexCreate(l *lexer) stateFn {
	l.lexAsToken(create);
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

func lexAssertInvariant(l *lexer) stateFn {
	l.pos += len("assert")
	for {
		if !contains(whitespace, l.next()) {
			l.backup()
			break;
		}
	}
	l.pos += len("invariant")
	l.emit(assertInvariant)
	l.skipWS()

	if (l.isNextPrefix(not)) {
		l.lexAsToken(not)
	}

	return lexToken
}

func lexRunQuery(l *lexer) stateFn {
	l.pos += len("run")
	for {
		if !contains(whitespace, l.next()) {
			l.backup()
			break;
		}
	}
	l.pos += len("query")
	l.emit(runQuery)
	l.skipWS()

	return lexToken
}

func lexApplyEvent(l *lexer) stateFn {
	l.pos += len("apply")
	for {
		if !contains(whitespace, l.next()) {
			l.backup()
			break;
		}
	}
	l.pos += len("event")
	l.emit(applyEvent)
	l.skipWS()

	return lexToken
}

func lexWhenEvent(l *lexer) stateFn {
	l.skipStr("when")
	l.skipWS()
	l.skipStr("event")
	l.skipWS()

	l.lexQuotedStringAsToken(whenEvent)

	return lexToken
}

func lexClass(l *lexer) stateFn {
	match, _ := l.matchingPrefix([]string{value, entity, event, command, query, invariant, projection})
	l.pos += len(match)
	l.emit(class)
	return lexToken
}

func lexTypeRef(l *lexer) stateFn {
	for {
		if (contains(whitespace, l.peek())) {
			break;
		}
		l.next();
	}

	l.emit(typeRef)
	l.skipWS()
	return lexIdentifier
}

func lexIdentifier(l *lexer) stateFn {
	l.scanIdentifier()
	word := l.input[l.start:l.pos]
	if (word == "true" || word == "false") {
		l.emit(boolean)
		return lexToken;
	}

	l.emit(identifier)
	return lexToken;
}

func lexString(l *lexer) stateFn {
	l.skip()
	for {
		if l.next() == '"' {
			l.backup()
			break;
		}
	}

	l.emit(string_)
	l.skip()

	return lexToken
}

func lexNumber(l *lexer) stateFn {
	hasDot := false;
	for {
		if (l.peek() == '.') {
			hasDot = true;
		}
		if (!isDigit(l.peek()) && l.peek() != '.') {
			break;
		}
		l.next();
	}
	if (hasDot) {
		l.emit(float)
	} else {
		l.emit(integer)
	}

	return lexToken;
}

func lexClassOpen(l *lexer) stateFn {
	l.lexAsToken(classOpen);
	l.skipWS()
	return lexClass;
}

func isLetter(ch int) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch int) bool {
	return '0' <= ch && ch <= '9'
}
