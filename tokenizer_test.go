package parser

import (
	"testing"
	"strconv"
)

var dbStatements = testStatements {
	{
		"create database 'db1';",
		[]Token{NewToken(create, "create", 0), NewToken(namespaceObject, "database", 7), NewToken(quotedName, "db1", 17), Semicolon(21)},
	}, {
		"create database 'db2' ;",
		[]Token{NewToken(create, "create", 0), NewToken(namespaceObject, "database", 7), NewToken(quotedName, "db2", 17), Semicolon(22)},
	},
};

func TestCreateDatabase(t *testing.T) {
	dbStatements.test(t);
}

var multipleStatements = testStatements{
	{
		"create database 'db1'; create database 'db1';",
		[]Token{tok(create, "create"), tok(namespaceObject, "database"), tok(quotedName, "db1"), semi(), tok(create, "create"), tok(namespaceObject, "database"), tok(quotedName, "db1"), semi()},
	},
}

func TestMultipeStatements(t *testing.T) {
	multipleStatements.test(t);
}

var domainStatements = testStatements{
	{
		"create domain 'dmn' using database 'db';",
		[]Token{tok(create, "create"), tok(namespaceObject, "domain"), tok(quotedName, "dmn"), tok(usingDatabase, "db"), semi()},

	},
	{
		"create domain 'dmn' using database 'db'",
		[]Token{tok(create, "create"), tok(namespaceObject, "domain"), tok(quotedName, "dmn"), tok(usingDatabase, "db")},

	},
};

func tok(typ TokenType, val string) Token {
	return Token{typ, val, ignoreTokenPos};
}

func semi() Token {
	return Semicolon(ignoreTokenPos);
}

func TestCreateDomain(t *testing.T) {
	domainStatements.test(t);
}


var contextStatements = testStatements {
	{
		"create context 'ctx' using database 'db' for domain 'dmn';",
		[]Token{tok(create, "create"), tok(namespaceObject, "context"), tok(quotedName, "ctx"), tok(usingDatabase, "db"), tok(forDomain, "dmn"), semi()},
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
		[]Token{tok(create, "create"), tok(namespaceObject, "aggregate"),tok(quotedName, "ag"), tok(usingDatabase, "db"), tok(forDomain, "dmn"), tok(inContext, "ctx"), semi()},
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
		[]Token{tok(usingDatabase, "db"), semi(), tok(create, "create"), tok(namespaceObject, "domain"), tok(quotedName, "dmn"), semi()},
	},
	{
		"for domain 'dmn'; create context 'ctx';",
		[]Token{tok(forDomain, "dmn"), semi(), tok(create, "create"), tok(namespaceObject, "context"), tok(quotedName, "ctx"), semi()},
	},
	{
		"in context 'ctx'; <| value 'address' |>",
		[]Token{tok(inContext, "ctx"), semi(), clsOpen(), tok(class, "value"), tok(quotedName, "address"), clsClose()},
	},
	{
		"within aggregate 'agg'; <| event 'start' |>",
		[]Token{tok(withinAggregate, "agg"), semi(), clsOpen(), tok(class, "event"), tok(quotedName, "start"), clsClose()},
	},
};

func TestGloballySetNamespace (t *testing.T) {
	statementsWithGloballySetNamespaces.test(t)
}

var objectTypes = testStatements {
	{
		"<| entity 'ent' |>",
		[]Token{clsOpen(), tok(class, "entity"), tok(quotedName, "ent"), clsClose()},
	},
	{
		"<| entity 'ent' check ( return value != 0;) |>",
		[]Token{
			clsOpen(),
			tok(class, "entity"),
			tok(quotedName, "ent"),

			tok(check, "check"),
			tok(lparen, "("),

			tok(return_, "return"),
			tok(identifier, "value"),
			tok(not_eq, "!="),
			tok(integer, "0"),
			tok(semicolon, ";"),

			tok(rparen, ")"),
			clsClose(),
		},
	},
	{
		"<| invariant 'invar' |>",
		[]Token{clsOpen(), tok(class, "invariant"), tok(quotedName, "invar"), clsClose()},
	},
	{
		"<| command 'cmd' |>",
		[]Token{clsOpen(), tok(class, "command"), tok(quotedName, "cmd"), clsClose()},
	},
	{
		"<| query 'qry' |>",
		[]Token{clsOpen(), tok(class, "query"), tok(quotedName, "qry"), clsClose()},
	},
	{
		"<| projection 'proj' |>",
		[]Token{clsOpen(), tok(class, "projection"), tok(quotedName, "proj"), clsClose()},
	},

}

func TestObjectTypes(t *testing.T) {
	objectTypes.test(t)
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
			tok(integer, "1"),
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
			tok(integer, "0"),
			tok(semicolon, ";"),

			tok(rparen, ")"),
		},
	},
	{
		`
		function doThing()
		{
			a = 2.1;
		}`,
		[]Token{
			tok(function, "function"),
			tok(identifier, "doThing"),
			tok(lparen, "("),
			tok(rparen, ")"),
			tok(lbrace, "{"),
			tok(identifier, "a"),
			tok(assign, "="),
			tok(float, "2.1"),
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
	{
		"a + (a - b)",
		[]Token{
			tok(identifier, "a"),
			tok(plus, "+"),
			tok(lparen, "("),
			tok(identifier, "a"),
			tok(minus, "-"),
			tok(identifier, "b"),
			tok(rparen, ")"),
		},
	},
	{
		"a->b->c + a->b() - !b and a == b and a < b or a > b ",
		[]Token{
			tok(identifier, "a"),
			tok(arrow, "->"),
			tok(identifier, "b"),
			tok(arrow, "->"),
			tok(identifier, "c"),
			tok(plus, "+"),
			tok(identifier, "a"),
			tok(arrow, "->"),
			tok(identifier, "b"),
			tok(lparen, "("),
			tok(rparen, ")"),
			tok(minus, "-"),
			tok(bang, "!"),
			tok(identifier, "b"),
			tok(and, "and"),
			tok(identifier, "a"),
			tok(eq, "=="),
			tok(identifier, "b"),
			tok(and, "and"),
			tok(identifier, "a"),
			tok(lt, "<"),
			tok(identifier, "b"),
			tok(or, "or"),
			tok(identifier, "a"),
			tok(gt, ">"),
			tok(identifier, "b"),
		},
	},
	{
		"a = andrew",
		[]Token {
			tok(identifier, "a"),
			tok(assign, "="),
			tok(identifier, "andrew"),
		},
	},
	{
		"clarkKent = 'value\\isSuperman'(false)",
		[]Token{
			tok(identifier, "clarkKent"),
			tok(assign, "="),
			tok(quotedName, "value\\isSuperman"),
			tok(lparen, "("),
			tok(boolean, "false"),
			tok(rparen, ")"),
		},
	},
	{
		`"string value"`,
		[]Token{
			tok(string_, "string value"),
		},
	},
};

func TestExpressions(t *testing.T) {
	expressions.test(t)
}


var statements = testStatements{
	{
		`if (a) {
			a;
		} else if (b) {
			a;
		} else {
			b;
		}
		foreach (a->b() as b=>c) {
			a;
		}`,
		[]Token{
			tok(if_, "if"),
			tok(lparen, "("),
			tok(identifier, "a"),
			tok(rparen, ")"),
			tok(lbrace, "{"),
			tok(identifier, "a"),
			tok(semicolon, ";"),
			tok(rbrace, "}"),

			tok(elseIf, "else if"),
			tok(lparen, "("),
			tok(identifier, "b"),
			tok(rparen, ")"),
			tok(lbrace, "{"),
			tok(identifier, "a"),
			tok(semicolon, ";"),
			tok(rbrace, "}"),

			tok(else_, "else"),
			tok(lbrace, "{"),
			tok(identifier, "b"),
			tok(semicolon, ";"),
			tok(rbrace, "}"),

			tok(foreach, "foreach"),
			tok(lparen, "("),
			tok(identifier, "a"),
			tok(arrow, "->"),
			tok(identifier, "b"),
			tok(lparen, "("),
			tok(rparen, ")"),
			tok(as, "as"),
			tok(identifier, "b"),
			tok(strongArrow, "=>"),
			tok(identifier, "c"),
			tok(rparen, ")"),

			tok(lbrace, "{"),
			tok(identifier, "a"),
			tok(semicolon, ";"),
			tok(rbrace, "}"),
		},
	},
}

func TestStatements(t *testing.T) {
	statements.test(t)
}

// These keywords should be seen as expressions, not keywords, dependent on context
var keyWordsAsExpressions = testStatements{
	{
		`
		database
		domain
		context
		aggregate
		value
		event
		entity
		command
		projection
		invariant
		query
		`,
		[]Token{
			tok(identifier, "database"),
			tok(identifier, "domain"),
			tok(identifier, "context"),
			tok(identifier, "aggregate"),
			tok(identifier, "value"),
			tok(identifier, "event"),
			tok(identifier, "entity"),
			tok(identifier, "command"),
			tok(identifier, "projection"),
			tok(identifier, "invariant"),
			tok(identifier, "query"),
		},
	},
}

func TestKeywordsAsExpressions(t *testing.T) {
	keyWordsAsExpressions.test(t)
}

// These keywords can be used in expressions only if they're part of an identifier
var keywordsInExpressions = testStatements {
	{
		`
		propertiesA
		checkA
		handlerA
		functionA
		whenA
		andA
		orA
		ifA
		elseA
		returnA
		foreachA
		asA
		createA`,
		[]Token {
			tok(identifier, "propertiesA"),
			tok(identifier, "checkA"),
			tok(identifier, "handlerA"),
			tok(identifier, "functionA"),
			tok(identifier, "whenA"),
			tok(identifier, "andA"),
			tok(identifier, "orA"),
			tok(identifier, "ifA"),
			tok(identifier, "elseA"),
			tok(identifier, "returnA"),
			tok(identifier, "foreachA"),
			tok(identifier, "asA"),
			tok(identifier, "createA"),
		},
	},
}

func TestKeywordsInExpressions(t *testing.T) {
	keywordsInExpressions.test(t)
}

type testStatement struct {
	dql string;
	expected []Token;
}

type testStatements []testStatement

func (statements testStatements) test(t *testing.T) {

	for _, statement := range statements {
		tokenizer := NewTokenizer(statement.dql);

		var token *Token
		var actual []Token
		var err *Token
		for {
			token, err = tokenizer.Next()
			if (token == nil) {
				break;
			}
			actual = append(actual, *token)
		}

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
