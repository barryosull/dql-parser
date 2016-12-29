package parser

import (
	"testing"
	"strconv"
)

var dbStatements = testStatements {
	{
		"create database 'db1';",
		[]Token{NewToken(create, "create", 0), NewToken(namespaceObject, "database", 7), NewToken(quotedName, "db1", 17), Apos(21)},
	}, {
		"create database 'db2' ;",
		[]Token{NewToken(create, "create", 0), NewToken(namespaceObject, "database", 7), NewToken(quotedName, "db2", 17), Apos(22)},
	},
};

func TestCreateDatabase(t *testing.T) {
	dbStatements.test(t);
}

var multipeStatements = testStatements{
	{
		"create database 'db1'; create database 'db1';",
		[]Token{tok(create, "create"), tok(namespaceObject, "database"), tok(quotedName, "db1"), apos(), tok(create, "create"), tok(namespaceObject, "database"), tok(quotedName, "db1"), apos()},
	},
}

func TestMultipeStatements(t *testing.T) {
	multipeStatements.test(t);
}

func compareTokens(a []Token, b []Token) bool {
	if (len(a) != len(b)) {
		return false;
	}
	for i, t := range a {
		if (!t.Compare(b[i])) {
			return false;
		}
	}
	return true;
}

var domainStatements = testStatements{
	{
		"create domain 'dmn' using database 'db';",
		[]Token{tok(create, "create"), tok(namespaceObject, "domain"), tok(quotedName, "dmn"), tok(usingDatabase, "db"), apos()},

	},
	{
		"create domain 'dmn' using database 'db'",
		[]Token{tok(create, "create"), tok(namespaceObject, "domain"), tok(quotedName, "dmn"), tok(usingDatabase, "db")},

	},
};

func tok(typ TokenType, val string) Token {
	return Token{typ, val, ignoreTokenPos};
}

func apos() Token {
	return Apos(ignoreTokenPos);
}

func TestCreateDomain(t *testing.T) {
	domainStatements.test(t);
}


var contextStatements = testStatements {
	{
		"create context 'ctx' using database 'db' for domain 'dmn';",
		[]Token{tok(create, "create"), tok(namespaceObject, "context"), tok(quotedName, "ctx"), tok(usingDatabase, "db"), tok(forDomain, "dmn"), apos()},
	},
};

func TestCreateContext(t *testing.T) {
	contextStatements.test(t);
}

var valueStatements = testStatements {
	{
		"<| value 'address' using database 'db' for domain 'dmn' in context 'ctx' |>",
		[]Token{clsOpen(), tok(class, "value"), tok(quotedName, "address"), tok(usingDatabase, "db"), tok(forDomain, "dmn"), tok(inContext, "ctx"), clsClose()},
	},
}

func clsOpen() Token {
	return ClsOpen(ignoreTokenPos);
}

func clsClose() Token {
	return ClsClose(ignoreTokenPos);
}

func TestCreateValue(t *testing.T) {
	valueStatements.test(t);
}

var aggregateStatements = testStatements{
	{
		"create aggregate 'ag' using database 'db' for domain 'dmn' in context 'ctx';",
		[]Token{tok(create, "create"), tok(namespaceObject, "aggregate"),tok(quotedName, "ag"), tok(usingDatabase, "db"), tok(forDomain, "dmn"), tok(inContext, "ctx"), apos()},
	},
}

func TestAggregateStatements (t *testing.T) {
	aggregateStatements.test(t)
}


var eventStatements = testStatements{
	{
		"<| event 'start' using database 'db' for domain 'dmn' in context 'ctx' within aggregate 'ag' |>",
		[]Token{clsOpen(), tok(class, "event"), tok(quotedName, "start"), tok(usingDatabase, "db"), tok(forDomain, "dmn"), tok(inContext, "ctx"), tok(withinAggregate, "ag"), clsClose()},
	},
}

func TestEventStatements (t *testing.T) {
	eventStatements.test(t)
}

var statementsWithGloballySetNamespaces = testStatements {
	{
		"using database 'db'; create domain 'dmn';",
		[]Token{tok(usingDatabase, "db"), apos(), tok(create, "create"), tok(namespaceObject, "domain"), tok(quotedName, "dmn"), apos()},
	},
	{
		"for domain 'dmn'; create context 'ctx';",
		[]Token{tok(forDomain, "dmn"), apos(), tok(create, "create"), tok(namespaceObject, "context"), tok(quotedName, "ctx"), apos()},
	},
	{
		"in context 'ctx'; <| value 'address' |>",
		[]Token{tok(inContext, "ctx"), apos(), clsOpen(), tok(class, "value"), tok(quotedName, "address"), clsClose()},
	},
	{
		"within aggregate 'agg'; <| event 'start' |>",
		[]Token{tok(withinAggregate, "agg"), apos(), clsOpen(), tok(class, "event"), tok(quotedName, "start"), clsClose()},
	},
};

func TestGloballySetNamespace (t *testing.T) {
	statementsWithGloballySetNamespaces.test(t)
}

var namespaceBlocks= testStatements {
	{
		`using database 'database1' for domain 'domain1' in context 'context1':{
			create aggregate 'aggregate1';

			using database 'database2' for domain 'domain2' in context 'context2':{
				create aggregate 'aggregate2';
			}
		}`,
		[]Token{
			tok(usingDatabase, "database1"),
			tok(forDomain, "domain1"),
			tok(inContext, "context1"),
			tok(colon, ":"),
			tok(lbrace, "{"),

			tok(create, "create"),
			tok(namespaceObject, "aggregate"),
			tok(quotedName, "aggregate1"),
			tok(semicolon, ";"),

			tok(usingDatabase, "database2"),
			tok(forDomain, "domain2"),
			tok(inContext, "context2"),
			tok(colon, ":"),
			tok(lbrace, "{"),

			tok(create, "create"),
			tok(namespaceObject, "aggregate"),
			tok(quotedName, "aggregate2"),
			tok(semicolon, ";"),

			tok(rbrace, "}"),
			tok(rbrace, "}"),
		},
	},
};

func TestNamespaceBlocks (t *testing.T) {
	namespaceBlocks.test(t)
}

var classComponents = testStatements{
	{
		`
		properties
		{
			value\service_charge service_charge = 'value\service_charge'(1);
			value\category category = [];
		}`,
		[]Token{
			tok(properties, "properties"),
			tok(lbrace, "{"),

			tok(typeRef, "value\\service_charge"),
			tok(identifier, "service_charge"),
			tok(assign, "="),
			tok(quotedName, "value\\service_charge"),
			tok(lparen, "("),
			tok(number, "1"),
			tok(rparen, ")"),
			tok(semicolon, ";"),

			tok(typeRef, "value\\category"),
			tok(identifier, "category"),
			tok(assign, "="),
			tok(lbracked, "["),
			tok(rbracket, "]"),
			tok(semicolon, ";"),

			tok(rbrace, "}"),
		},
	},
	{
		`
		check
		(
			return value != 0;
		)`,
		[]Token{
			tok(check, "check"),
			tok(lparen, "("),

			tok(return_, "return"),
			tok(identifier, "value"),
			tok(not_eq, "!="),
			tok(number, "0"),
			tok(semicolon, ";"),

			tok(rparen, ")"),
		},
	},
	{
		`
		function doThing()
		{
			a = 2;
		}`,
		[]Token{
			tok(function, "function"),
			tok(identifier, "doThing"),
			tok(lparen, "("),
			tok(rparen, ")"),
			tok(lbrace, "{"),
			tok(identifier, "a"),
			tok(assign, "="),
			tok(number, "2"),
			tok(semicolon, ";"),
			tok(rbrace, "}"),
		},
	},
	{
		`
		function doThing2(value\service-charge service_charge, value\category category)
		{

		}`,
		[]Token{
			tok(function, "function"),
			tok(identifier, "doThing2"),
			tok(lparen, "("),
			tok(typeRef, "value\\service-charge"),
			tok(identifier, "service_charge"),
			tok(comma, ","),
			tok(typeRef, "value\\category"),
			tok(identifier, "category"),
			tok(rparen, ")"),
			tok(lbrace, "{"),
			tok(rbrace, "}"),
		},
	},
	{
		`
		handler
		{
			assert  invariant not 'is-started';
			revision = run query 'next-revision-number' (agency_id, quote_number);
			apply event 'started' (agency_id, brand_id, quote_number, revision);
		}`,
		[]Token{
			tok(handler, "handler"),
			tok(lbrace, "{"),
			tok(assertInvariant, "assert  invariant"),
			tok(not, "not"),
			tok(quotedName, "is-started"),
			tok(semicolon, ";"),
			tok(identifier, "revision"),
			tok(assign, "="),
			tok(runQuery, "run query"),
			tok(quotedName, "next-revision-number"),
			tok(lparen, "("),
			tok(identifier, "agency_id"),
			tok(comma, ","),
			tok(identifier, "quote_number"),
			tok(rparen, ")"),
			tok(semicolon, ";"),
			tok(applyEvent, "apply event"),
			tok(quotedName, "started"),
			tok(lparen, "("),
			tok(identifier, "agency_id"),
			tok(comma, ","),
			tok(identifier, "brand_id"),
			tok(comma, ","),
			tok(identifier, "quote_number"),
			tok(comma, ","),
			tok(identifier, "revision"),
			tok(rparen, ")"),
			tok(semicolon, ";"),
			tok(rbrace, "}"),

		},
	},
	{
		`
		when event 'started'
		{
			agency_id = event->agency_id;
			is_started = true;
		}`,
		[]Token{
			tok(whenEvent, "started"),
			tok(lbrace, "{"),
			tok(identifier, "agency_id"),
			tok(assign, "="),
			tok(identifier, "event"),
			tok(arrow, "->"),
			tok(identifier, "agency_id"),
			tok(semicolon, ";"),
			tok(identifier, "is_started"),
			tok(assign, "="),
			tok(boolean, "true"),
			tok(semicolon, ";"),
			tok(rbrace, "}"),
		},

	},
};

func TestClassComponents (t *testing.T) {
	classComponents.test(t)
}

var expressions = testStatements{
	{
		"a + b - c",
		[]Token{
			tok(identifier, "a"),
			tok(plus, "+"),
			tok(identifier, "b"),
			tok(minus, "-"),
			tok(identifier, "c"),
		},
	},
	/*
		"a + (a - b)",
		"(a + b) + (a - b)",
		"(a + b) + (a - b) - a->b->c + a->b() - !b + a and b",
		"a->b = 'value\\integer'(1) - ('value\\integer'(1) + b) + (a - b) - a->b->c + a->b() - !b + a and b",
		"a = b + c = c + 24",
		"quote->items->has(item) == true",
		"quote->items->has(item) - 5 == true",
		"quote->is_started == true and quote->is_completed == false",
		"quote->is_started == true",
	}
	*/
};

func TestExpressions(t *testing.T) {
	expressions.test(t)
}


type testStatement struct {
	dql string;
	expected []Token;
}

type testStatements []testStatement

func (statements testStatements) test(t *testing.T) {

	for _, statement := range statements {
		tokenizer := NewTokenizer(statement.dql);

		actual, err := tokenizer.Tokens();

		compareTokenLists(statement.expected, actual, statement.dql, t);

		if (err != nil) {
			t.Error("Got error")
			t.Error(err);
		}
	}
}

func compareTokenLists(expected, actual []Token, dql string, t *testing.T) {
	if (len(expected) != len(actual)) {
		t.Error("Error with AST produced from '"+dql+"'");
		t.Error("Number of tokens are mismtached, expected "+strconv.Itoa(len(expected))+", got "+strconv.Itoa(len(actual)));
	}
	for i, token := range expected {
		if i == len(actual) {
			t.Error("Expected: "+token.String())
			t.Error("Got: nothing")
			return
		}
		if (!token.Compare(actual[i])) {
			t.Error("Expected: "+token.String())
			t.Error("Got: "+actual[i].String())
			return
		}
	}
}
