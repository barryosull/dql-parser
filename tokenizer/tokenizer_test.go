package tokenizer

import (
	"testing"
	"strconv"
	token "parser/token"
)

var dbStatements = testStatements {
	{
		"create database 'db1';",
		[]token.Token{token.NewToken(token.create, "create", 0), token.NewToken(token.namespaceObject, "database", 7), token.NewToken(token.quotedName, "db1", 17), token.Semicolon(21)},
	}, {
		"create database 'db2' ;",
		[]token.Token{token.NewToken(token.create, "create", 0), token.NewToken(token.namespaceObject, "database", 7), token.NewToken(token.quotedName, "db2", 17), token.Semicolon(22)},
	},
};

func TestCreateDatabase(t *testing.T) {
	dbStatements.test(t);
}

var multipleStatements = testStatements{
	{
		"create database 'db1'; create database 'db1';",
		[]token.Token{tok(token.create, "create"), tok(token.namespaceObject, "database"), tok(token.quotedName, "db1"), semi(), tok(token.create, "create"), tok(token.namespaceObject, "database"), tok(token.quotedName, "db1"), semi()},
	},
}

func TestMultipeStatements(t *testing.T) {
	multipleStatements.test(t);
}

var domainStatements = testStatements{
	{
		"create domain 'dmn' using database 'db';",
		[]token.Token{tok(token.create, "create"), tok(token.namespaceObject, "domain"), tok(token.quotedName, "dmn"), tok(token.usingDatabase, "db"), semi()},

	},
	{
		"create domain 'dmn' using database 'db'",
		[]token.Token{tok(token.create, "create"), tok(token.namespaceObject, "domain"), tok(token.quotedName, "dmn"), tok(token.usingDatabase, "db")},

	},
};

func tok(typ token.TokenType, val string) token.Token {
	return token.Token{typ, val, token.ignoreTokenPos};
}

func semi() token.Token {
	return token.Semicolon(token.ignoreTokenPos);
}

func TestCreateDomain(t *testing.T) {
	domainStatements.test(t);
}


var contextStatements = testStatements {
	{
		"create context 'ctx' using database 'db' for domain 'dmn';",
		[]token.Token{tok(token.create, "create"), tok(token.namespaceObject, "context"), tok(token.quotedName, "ctx"), tok(token.usingDatabase, "db"), tok(token.forDomain, "dmn"), semi()},
	},
};

func TestCreateContext(t *testing.T) {
	contextStatements.test(t);
}

var valueStatements = testStatements {
	{
		"<| value 'address' using database 'db' for domain 'dmn' in context 'ctx' |>",
		[]token.Token{clsOpen(), tok(token.class, "value"), tok(token.quotedName, "address"), tok(token.usingDatabase, "db"), tok(token.forDomain, "dmn"), tok(token.inContext, "ctx"), clsClose()},
	},
}

func clsOpen() Token {
	return token.ClsOpen(token.ignoreTokenPos);
}

func clsClose() Token {
	return token.ClsClose(token.ignoreTokenPos);
}

func TestCreateValue(t *testing.T) {
	valueStatements.test(t);
}

var aggregateStatements = testStatements{
	{
		"create aggregate 'ag' using database 'db' for domain 'dmn' in context 'ctx';",
		[]token.Token{tok(token.create, "create"), tok(token.namespaceObject, "aggregate"),tok(token.quotedName, "ag"), tok(token.usingDatabase, "db"), tok(token.forDomain, "dmn"), tok(token.inContext, "ctx"), semi()},
	},
}

func TestAggregateStatements (t *testing.T) {
	aggregateStatements.test(t)
}


var eventStatements = testStatements{
	{
		"<| event 'start' using database 'db' for domain 'dmn' in context 'ctx' within aggregate 'ag' |>",
		[]token.Token{clsOpen(), tok(token.class, "event"), tok(token.quotedName, "start"), tok(token.usingDatabase, "db"), tok(token.forDomain, "dmn"), tok(token.inContext, "ctx"), tok(token.withinAggregate, "ag"), clsClose()},
	},
}

func TestEventStatements (t *testing.T) {
	eventStatements.test(t)
}

var statementsWithGloballySetNamespaces = testStatements {
	{
		"using database 'db'; create domain 'dmn';",
		[]token.Token{tok(token.usingDatabase, "db"), semi(), tok(token.create, "create"), tok(token.namespaceObject, "domain"), tok(token.quotedName, "dmn"), semi()},
	},
	{
		"for domain 'dmn'; create context 'ctx';",
		[]token.Token{tok(token.forDomain, "dmn"), semi(), tok(token.create, "create"), tok(token.namespaceObject, "context"), tok(token.quotedName, "ctx"), semi()},
	},
	{
		"in context 'ctx'; <| value 'address' |>",
		[]token.Token{tok(token.inContext, "ctx"), semi(), clsOpen(), tok(token.class, "value"), tok(token.quotedName, "address"), clsClose()},
	},
	{
		"within aggregate 'agg'; <| event 'start' |>",
		[]token.Token{tok(token.withinAggregate, "agg"), semi(), clsOpen(), tok(token.class, "event"), tok(token.quotedName, "start"), clsClose()},
	},
};

func TestGloballySetNamespace (t *testing.T) {
	statementsWithGloballySetNamespaces.test(t)
}

var objectTypes = testStatements {
	{
		"<| entity 'ent' |>",
		[]token.Token{clsOpen(), tok(token.class, "entity"), tok(token.quotedName, "ent"), clsClose()},
	},
	{
		"<| entity 'ent' check ( return value != 0;) |>",
		[]token.Token{
			clsOpen(),
			tok(token.class, "entity"),
			tok(token.quotedName, "ent"),

			tok(token.check, "check"),
			tok(token.lparen, "("),

			tok(token.return_, "return"),
			tok(token.identifier, "value"),
			tok(token.not_eq, "!="),
			tok(token.integer, "0"),
			tok(token.semicolon, ";"),

			tok(token.rparen, ")"),
			clsClose(),
		},
	},
	{
		"<| invariant 'invar' |>",
		[]token.Token{clsOpen(), tok(token.class, "invariant"), tok(token.quotedName, "invar"), clsClose()},
	},
	{
		"<| command 'cmd' |>",
		[]token.Token{clsOpen(), tok(token.class, "command"), tok(token.quotedName, "cmd"), clsClose()},
	},
	{
		"<| query 'qry' |>",
		[]token.Token{clsOpen(), tok(token.class, "query"), tok(token.quotedName, "qry"), clsClose()},
	},
	{
		"<| projection 'proj' |>",
		[]token.Token{clsOpen(), tok(token.class, "projection"), tok(token.quotedName, "proj"), clsClose()},
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
		[]token.Token{
			tok(token.usingDatabase, "database1"),
			tok(token.forDomain, "domain1"),
			tok(token.inContext, "context1"),
			tok(token.colon, ":"),
			tok(token.lbrace, "{"),

			tok(token.create, "create"),
			tok(token.namespaceObject, "aggregate"),
			tok(token.quotedName, "aggregate1"),
			tok(token.semicolon, ";"),

			tok(token.usingDatabase, "database2"),
			tok(token.forDomain, "domain2"),
			tok(token.inContext, "context2"),
			tok(token.colon, ":"),
			tok(token.lbrace, "{"),

			tok(token.create, "create"),
			tok(token.namespaceObject, "aggregate"),
			tok(token.quotedName, "aggregate2"),
			tok(token.semicolon, ";"),

			tok(token.rbrace, "}"),
			tok(token.rbrace, "}"),
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
		[]token.Token{
			tok(token.properties, "properties"),
			tok(token.lbrace, "{"),

			tok(token.typeRef, "value\\service_charge"),
			tok(token.identifier, "service_charge"),
			tok(token.assign, "="),
			tok(token.quotedName, "value\\service_charge"),
			tok(token.lparen, "("),
			tok(token.integer, "1"),
			tok(token.rparen, ")"),
			tok(token.semicolon, ";"),

			tok(token.typeRef, "value\\category"),
			tok(token.identifier, "category"),
			tok(token.assign, "="),
			tok(token.lbracked, "["),
			tok(token.rbracket, "]"),
			tok(token.semicolon, ";"),

			tok(token.rbrace, "}"),
		},
	},
	{
		`
		check
		(
			return value != 0;
		)`,
		[]token.Token{
			tok(token.check, "check"),
			tok(token.lparen, "("),

			tok(token.return_, "return"),
			tok(token.identifier, "value"),
			tok(token.not_eq, "!="),
			tok(token.integer, "0"),
			tok(token.semicolon, ";"),

			tok(token.rparen, ")"),
		},
	},
	{
		`
		function doThing()
		{
			a = 2.1;
		}`,
		[]token.Token{
			tok(token.function, "function"),
			tok(token.identifier, "doThing"),
			tok(token.lparen, "("),
			tok(token.rparen, ")"),
			tok(token.lbrace, "{"),
			tok(token.identifier, "a"),
			tok(token.assign, "="),
			tok(token.float, "2.1"),
			tok(token.semicolon, ";"),
			tok(token.rbrace, "}"),
		},
	},
	{
		`
		function doThing2(value\service-charge service_charge, value\category category)
		{

		}`,
		[]token.Token{
			tok(token.function, "function"),
			tok(token.identifier, "doThing2"),
			tok(token.lparen, "("),
			tok(token.typeRef, "value\\service-charge"),
			tok(token.identifier, "service_charge"),
			tok(token.comma, ","),
			tok(token.typeRef, "value\\category"),
			tok(token.identifier, "category"),
			tok(token.rparen, ")"),
			tok(token.lbrace, "{"),
			tok(token.rbrace, "}"),
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
		[]token.Token{
			tok(token.handler, "handler"),
			tok(token.lbrace, "{"),
			tok(token.assertInvariant, "assert  invariant"),
			tok(token.not, "not"),
			tok(token.quotedName, "is-started"),
			tok(token.semicolon, ";"),
			tok(token.identifier, "revision"),
			tok(token.assign, "="),
			tok(token.runQuery, "run query"),
			tok(token.quotedName, "next-revision-number"),
			tok(token.lparen, "("),
			tok(token.identifier, "agency_id"),
			tok(token.comma, ","),
			tok(token.identifier, "quote_number"),
			tok(token.rparen, ")"),
			tok(token.semicolon, ";"),
			tok(token.applyEvent, "apply event"),
			tok(token.quotedName, "started"),
			tok(token.lparen, "("),
			tok(token.identifier, "agency_id"),
			tok(token.comma, ","),
			tok(token.identifier, "brand_id"),
			tok(token.comma, ","),
			tok(token.identifier, "quote_number"),
			tok(token.comma, ","),
			tok(token.identifier, "revision"),
			tok(token.rparen, ")"),
			tok(token.semicolon, ";"),
			tok(token.rbrace, "}"),

		},
	},
	{
		`
		when event 'started'
		{
			agency_id = event->agency_id;
			is_started = true;
		}`,
		[]token.Token{
			tok(token.whenEvent, "started"),
			tok(token.lbrace, "{"),
			tok(token.identifier, "agency_id"),
			tok(token.assign, "="),
			tok(token.identifier, "event"),
			tok(token.arrow, "->"),
			tok(token.identifier, "agency_id"),
			tok(token.semicolon, ";"),
			tok(token.identifier, "is_started"),
			tok(token.assign, "="),
			tok(token.boolean, "true"),
			tok(token.semicolon, ";"),
			tok(token.rbrace, "}"),
		},

	},
};

func TestClassComponents (t *testing.T) {
	classComponents.test(t)
}

var expressions = testStatements{
	{
		`--a
		a++
		a <= b
		b >= a`,
		[]token.Token{
			tok(token.minus, "-"),
			tok(token.minus, "-"),
			tok(token.identifier, "a"),
			tok(token.identifier, "a"),
			tok(token.plus, "+"),
			tok(token.plus, "+"),
			tok(token.identifier, "a"),
			tok(token.ltOrEq, "<="),
			tok(token.identifier, "b"),
			tok(token.identifier, "b"),
			tok(token.gtOrEq, ">="),
			tok(token.identifier, "a"),
		},
	},
	{
		"a + b - c",
		[]token.Token{
			tok(token.identifier, "a"),
			tok(token.plus, "+"),
			tok(token.identifier, "b"),
			tok(token.minus, "-"),
			tok(token.identifier, "c"),
		},
	},
	{
		"a + (a - b)",
		[]token.Token{
			tok(token.identifier, "a"),
			tok(token.plus, "+"),
			tok(token.lparen, "("),
			tok(token.identifier, "a"),
			tok(token.minus, "-"),
			tok(token.identifier, "b"),
			tok(token.rparen, ")"),
		},
	},
	{
		"a->b->c + a->b() - !b and a == b and a < b or a > b ",
		[]token.Token{
			tok(token.identifier, "a"),
			tok(token.arrow, "->"),
			tok(token.identifier, "b"),
			tok(token.arrow, "->"),
			tok(token.identifier, "c"),
			tok(token.plus, "+"),
			tok(token.identifier, "a"),
			tok(token.arrow, "->"),
			tok(token.identifier, "b"),
			tok(token.lparen, "("),
			tok(token.rparen, ")"),
			tok(token.minus, "-"),
			tok(token.bang, "!"),
			tok(token.identifier, "b"),
			tok(token.and, "and"),
			tok(token.identifier, "a"),
			tok(token.eq, "=="),
			tok(token.identifier, "b"),
			tok(token.and, "and"),
			tok(token.identifier, "a"),
			tok(token.lt, "<"),
			tok(token.identifier, "b"),
			tok(token.or, "or"),
			tok(token.identifier, "a"),
			tok(token.gt, ">"),
			tok(token.identifier, "b"),
		},
	},
	{
		"a = andrew",
		[]token.Token {
			tok(token.identifier, "a"),
			tok(token.assign, "="),
			tok(token.identifier, "andrew"),
		},
	},
	{
		"clarkKent = 'value\\isSuperman'(false)",
		[]token.Token{
			tok(token.identifier, "clarkKent"),
			tok(token.assign, "="),
			tok(token.quotedName, "value\\isSuperman"),
			tok(token.lparen, "("),
			tok(token.boolean, "false"),
			tok(token.rparen, ")"),
		},
	},
	{
		`"string value"`,
		[]token.Token{
			tok(token.string_, "string value"),
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
		[]token.Token{
			tok(token.if_, "if"),
			tok(token.lparen, "("),
			tok(token.identifier, "a"),
			tok(token.rparen, ")"),
			tok(token.lbrace, "{"),
			tok(token.identifier, "a"),
			tok(token.semicolon, ";"),
			tok(token.rbrace, "}"),

			tok(token.elseIf, "else if"),
			tok(token.lparen, "("),
			tok(token.identifier, "b"),
			tok(token.rparen, ")"),
			tok(token.lbrace, "{"),
			tok(token.identifier, "a"),
			tok(token.semicolon, ";"),
			tok(token.rbrace, "}"),

			tok(token.else_, "else"),
			tok(token.lbrace, "{"),
			tok(token.identifier, "b"),
			tok(token.semicolon, ";"),
			tok(token.rbrace, "}"),

			tok(token.foreach, "foreach"),
			tok(token.lparen, "("),
			tok(token.identifier, "a"),
			tok(token.arrow, "->"),
			tok(token.identifier, "b"),
			tok(token.lparen, "("),
			tok(token.rparen, ")"),
			tok(token.as, "as"),
			tok(token.identifier, "b"),
			tok(token.strongArrow, "=>"),
			tok(token.identifier, "c"),
			tok(token.rparen, ")"),

			tok(token.lbrace, "{"),
			tok(token.identifier, "a"),
			tok(token.semicolon, ";"),
			tok(token.rbrace, "}"),
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
		[]token.Token{
			tok(token.identifier, "database"),
			tok(token.identifier, "domain"),
			tok(token.identifier, "context"),
			tok(token.identifier, "aggregate"),
			tok(token.identifier, "value"),
			tok(token.identifier, "event"),
			tok(token.identifier, "entity"),
			tok(token.identifier, "command"),
			tok(token.identifier, "projection"),
			tok(token.identifier, "invariant"),
			tok(token.identifier, "query"),
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
		[]token.Token {
			tok(token.identifier, "propertiesA"),
			tok(token.identifier, "checkA"),
			tok(token.identifier, "handlerA"),
			tok(token.identifier, "functionA"),
			tok(token.identifier, "whenA"),
			tok(token.identifier, "andA"),
			tok(token.identifier, "orA"),
			tok(token.identifier, "ifA"),
			tok(token.identifier, "elseA"),
			tok(token.identifier, "returnA"),
			tok(token.identifier, "foreachA"),
			tok(token.identifier, "asA"),
			tok(token.identifier, "createA"),
		},
	},
}

func TestKeywordsInExpressions(t *testing.T) {
	keywordsInExpressions.test(t)
}

type testStatement struct {
	dql string;
	expected []token.Token;
}

type testStatements []testStatement

func (statements testStatements) test(t *testing.T) {

	for _, statement := range statements {
		tokenizer := NewTokenizer(statement.dql);

		var token *token.Token
		var actual []token.Token
		var err *token.Token
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

func compareTokenLists(expected, actual []token.Token, dql string, t *testing.T) {
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
