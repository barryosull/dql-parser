package tokenizer

import (
	"strings"
	"unicode/utf8"
	tok "parser/token"
)

// lexer holds the state of the scanner.
type lexer struct {
	name  string    // used only for error reports.
	input string    // the string being scanned.
	start int       // start position of this item.
	pos   int       // current position in the input.
	width int       // width of last rune read from input.
	tokens []tok.Token
	error *tok.Error
}

type stateFn func(*lexer) stateFn

var tokenToLexer = map[string]stateFn {}
var easyLexKeywords = []tok.TokenType{}
var easyLexTokens = []tok.TokenType{}

func lex(name, input string) (*lexer) {
	l := &lexer{
		name:  name,
		input: input,
	}

	tokenToLexer = map[string]stateFn{
		"create": lexCreate,
		"using": lexUsingDatabase,
		"for": lexForDomain,
		"in": lexInContext,
		"within": lexWithinAggregate,
		"assert": lexAssertInvariant,
		"run": lexRunQuery,
		"apply": lexApplyEvent,
		"when": lexWhenEvent,
	}

	easyLexKeywords = []tok.TokenType{
		tok.AND,
		tok.OR,
		tok.PROPERTIES,
		tok.CHECK,
		tok.HANDLER,
		tok.FUNCTION,
		tok.IF,
		tok.ELSEIF,
		tok.ELSE,
		tok.FOREACH,
		tok.RETURN,
		tok.AS,
	}

	easyLexTokens = []tok.TokenType{
		tok.STRONGARROW,
		tok.EQ,
		tok.SEMICOLON,
		tok.ASSIGN,
		tok.CLASSCLOSE,
		tok.COLON,
		tok.LBRACE,
		tok.RBRACE,
		tok.LPAREN,
		tok.RPAREN,
		tok.LBRACKET,
		tok.RBRACKET,
		tok.NOTEQ,
		tok.COMMA,
		tok.ARROW,
		tok.PLUS,
		tok.MINUS,
		tok.BANG,
		tok.ASTERISK,
		tok.SLASH,
		tok.LTOREQ,
		tok.GTOREQ,
		tok.LT,
		tok.GT,
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

func (l *lexer) emit(t tok.TokenType) {
	l.tokens = append(l.tokens, tok.Token{t, l.input[l.start:l.pos], l.start});
	l.start = l.pos
}

func (l *lexer) isNextPrefix(prefix string) bool {
	var unlexed = l.input[l.pos:];
	return strings.HasPrefix(unlexed, prefix);
}

//Check of the prefix matches and is not followed immediately by another identifier character
func (l *lexer) isKeyWordAndNotIdentifier(prefix string) bool {
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

func (l *lexer) matchPrefix(expected []string) (string, *tok.Error) {
	for _, prefix := range expected {
		if (l.isNextPrefix(prefix)) {
			return prefix, nil
		}
	}
	l.scanIdentifier()
	found := l.input[l.start:l.pos]
	l.err(strings.Join(expected, ", "), found)
	return found, l.error
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

func (l *lexer) lexQuotedStringAsToken(tokenType tok.TokenType) bool {
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

func (l *lexer) lexAsToken(tokenType tok.TokenType) stateFn {
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

func (l *lexer) err(expected string, found string) stateFn {
	l.error = &tok.Error{l.input, l.start, expected, found}
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



func (l *lexer) isTypeRef() bool {
	if (l.isKeyWordAndNotIdentifier(tok.STRING)) {
		return true;
	}
	if (l.isKeyWordAndNotIdentifier(tok.INTEGER)) {
		return true;
	}
	if (l.isKeyWordAndNotIdentifier(tok.FLOAT)) {
		return true;
	}
	if (l.isKeyWordAndNotIdentifier(tok.BOOLEAN)) {
		return true;
	}
	if (l.isNextPrefix(tok.VALUE+"\\") || l.isNextPrefix(tok.ENTITY+"\\") || l.isNextPrefix(tok.EVENT+"\\") || l.isNextPrefix(tok.COMMAND+"\\") || l.isNextPrefix(tok.INVARIANT+"\\") || l.isNextPrefix(tok.PROJECTION+"\\")|| l.isNextPrefix(tok.QUERY+"\\")) {
		return true
	}
	return false
}

func lexToken(l *lexer) stateFn {

	l.skipWS();

	if (l.peek() == EOF) {
		return nil;
	}

	//Have special lexing rules
	if (l.isTypeRef()) {
		return lexTypeRef
	}
	if (l.isNextPrefix("'")) {
		return lexNSObjectName
	}
	if (l.isNextPrefix("\"")) {
		return lexString
	}
	if (l.isNextPrefix(tok.CLASSOPEN)) {
		return lexClassOpen;
	}

	for token, stateFn := range tokenToLexer {
		if l.isKeyWordAndNotIdentifier(token) {
			return stateFn
		}
	}

	for _, token := range easyLexKeywords {
		if l.isKeyWordAndNotIdentifier(string(token)) {
			return l.lexAsToken(token)
		}
	}

	for _, token := range easyLexTokens {
		if l.isNextPrefix(string(token)) {
			return l.lexAsToken(token)
		}
	}

	if (isDigit(l.peek())) {
		return lexNumber;
	}
	if (isLetter(l.peek())) {
		return lexIdentifier;
	}

	return l.err("keyword", "nothing");
}

func lexCreate(l *lexer) stateFn {
	l.lexAsToken(tok.CREATE);
	return lexNSObjectType
}

func lexNSObjectType(l *lexer) stateFn {
	l.skipWS()
	found, err := l.matchPrefix([]string{tok.DATABASE, tok.DOMAIN, tok.CONTEXT, tok.AGGREGATE})
	if (err != nil) {
		return nil
	}
	l.pos += len(found)
	l.emit(tok.NAMESPACEOBJECT)
	return lexToken
}

func lexNSObjectName(l *lexer) stateFn {
	l.lexQuotedStringAsToken(tok.QUOTEDNAME)
	return lexToken
}

func lexUsingDatabase(l *lexer) stateFn {
	l.skipStr("using")
	l.skipStr(tok.DATABASE)

	if l.lexQuotedStringAsToken(tok.USINGDATABASE) {
		return lexToken
	}
	return nil
}

func lexForDomain (l *lexer) stateFn {
	l.skipStr("for")
	l.skipStr(tok.DOMAIN)

	if l.lexQuotedStringAsToken(tok.FORDOMAIN) {
		return lexToken
	}
	return nil
}

func lexInContext (l *lexer) stateFn {
	l.skipStr("in")
	l.skipStr(tok.CONTEXT)

	if l.lexQuotedStringAsToken(tok.INCONTEXT) {
		return lexToken
	}
	return nil
}

func lexWithinAggregate(l *lexer) stateFn {
	l.skipStr("within")
	l.skipStr(tok.AGGREGATE)

	if l.lexQuotedStringAsToken(tok.WITHINAGGREGATE) {
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
	l.emit(tok.ASSERTINVARIANT)
	l.skipWS()

	if (l.isNextPrefix(tok.NOT)) {
		l.lexAsToken(tok.NOT)
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
	l.emit(tok.RUNQUERY)
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
	l.emit(tok.APPLYEVENT)
	l.skipWS()

	return lexToken
}

func lexWhenEvent(l *lexer) stateFn {
	l.skipStr("when")
	l.skipWS()
	l.skipStr("event")
	l.skipWS()

	l.lexQuotedStringAsToken(tok.WHENEVENT)

	return lexToken
}

func lexClass(l *lexer) stateFn {
	match, _ := l.matchPrefix([]string{tok.VALUE, tok.ENTITY, tok.EVENT, tok.COMMAND, tok.QUERY, tok.INVARIANT, tok.PROJECTION})
	l.pos += len(match)
	l.emit(tok.CLASS)
	return lexToken
}

func lexTypeRef(l *lexer) stateFn {
	for {
		if (contains(whitespace, l.peek())) {
			break;
		}
		l.next();
	}

	l.emit(tok.TYPEREF)
	l.skipWS()
	return lexIdentifier
}

func lexIdentifier(l *lexer) stateFn {
	l.scanIdentifier()
	word := l.input[l.start:l.pos]
	if (word == "true" || word == "false") {
		l.emit(tok.BOOLEAN)
		return lexToken;
	}

	l.emit(tok.IDENTIFIER)
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

	l.emit(tok.STRING)
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
		l.emit(tok.FLOAT)
	} else {
		l.emit(tok.INTEGER)
	}

	return lexToken;
}

func lexClassOpen(l *lexer) stateFn {
	l.lexAsToken(tok.CLASSOPEN);
	l.skipWS()
	return lexClass;
}

func isLetter(ch int) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch int) bool {
	return '0' <= ch && ch <= '9'
}
