package parser

import (
	"testing"
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
			tok(apostrophe, ";"),

			tok(usingDatabase, "database2"),
			tok(forDomain, "domain2"),
			tok(inContext, "context2"),
			tok(colon, ":"),
			tok(lbrace, "{"),

			tok(create, "create"),
			tok(namespaceObject, "aggregate"),
			tok(quotedName, "aggregate2"),
			tok(apostrophe, ";"),

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
			tok(apostrophe, ";"),

			tok(typeRef, "value\\category"),
			tok(identifier, "category"),
			tok(assign, "="),
			tok(lbracked, "["),
			tok(lbracked, "]"),
			tok(apostrophe, ";"),

			tok(rbrace, "}"),
		},
	},
	/*
	`
	check
	(
		return value != 0;
	)`,
	`
	function doThing()
	{
		a = 2;
	}`,
	`
	function doThing2(value\service-charge service_charge, value\category category)
	{

	}`,
	`
	handler
	{
		a = b + c;
		assert invariant not 'is-started';
		revision = run query 'next-revision-number' (agency_id, quote_number);
		apply event 'started' (agent_id, agency_id, brand_id, quote_number, revision);
	}`,
	`
	when event 'started'
	{
		agency_id = event->agency_id;
		brand_id = event->brand_id;
		is_started = true;
	}`,
	*/
};

func TestClassComponents (t *testing.T) {
	classComponents.test(t)
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

		if (!compareTokens(statement.expected, actual)) {
			t.Error("AST produced from '"+statement.dql+"' is not valid");
			t.Error(statement.expected);
			t.Error(actual);
		}

		if (err != nil) {
			t.Error("Got error")
			t.Error(err);
		}
	}
}