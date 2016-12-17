package parser

import (
	"testing"
)

type testStatement struct {
	dql string;
	expected []Token;
	error *Token;
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

		if (err == nil && statement.error != nil) {
			t.Error("Expected error, got nothing.")
			t.Error(statement.dql);
			t.Error(err);
		} else if (err != nil && statement.error == nil) {
			t.Error("Got error, expected nothing.")
			t.Error(err);
		}

		if (err != nil && statement.error != nil) {
			if (err.String() != statement.error.String()) {
				t.Error("Errors do not match.")
				t.Error(statement.error);
				t.Error(err);
			}
		}
	}
}

var dbStatements = testStatements {
	{
		"create database 'db1';",
		[]Token{NewToken(create, "create", 0), NewToken(namespaceObject, "database", 7), NewToken(quotedName, "db1", 17), Apos(21)}, nil,
	}, {
		"create database 'db2' ;",
		[]Token{NewToken(create, "create", 0), NewToken(namespaceObject, "database", 7), NewToken(quotedName, "db2", 17), Apos(22)}, nil,
	}, {
		"create dbase 'db2' ",
		[]Token{NewToken(create, "create", 0)}, Err("There was a problem near: \"create \"", 20),
	},
};

func TestCreateDatabase(t *testing.T) {
	dbStatements.test(t);
}

var multipeStatements = testStatements{
	{
		"create database 'db1'; create database 'db1';",
		[]Token{tok(create, "create"), tok(namespaceObject, "database"), tok(quotedName, "db1"), apos(), tok(create, "create"), tok(namespaceObject, "database"), tok(quotedName, "db1"), apos()}, nil,
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
		[]Token{tok(create, "create"), tok(namespaceObject, "domain"), tok(quotedName, "dmn"), tok(usingDatabase, "db"), apos()}, nil,

	},
	{
		"create domain 'dmn' using database 'db'",
		[]Token{tok(create, "create"), tok(namespaceObject, "domain"), tok(quotedName, "dmn"), tok(usingDatabase, "db")}, nil,

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
		[]Token{tok(create, "create"), tok(namespaceObject, "context"), tok(quotedName, "ctx"), tok(usingDatabase, "db"), tok(forDomain, "dmn"), apos()}, nil,
	},
};

func TestCreateContext(t *testing.T) {
	contextStatements.test(t);
}

var valueStatements = testStatements {
	{
		"<| value 'address' using database 'db' for domain 'dmn' in context 'ctx' |>",
		[]Token{clsOpen(), tok(class, "value"), tok(quotedName, "address"), tok(usingDatabase, "db"), tok(forDomain, "dmn"), tok(inContext, "ctx"), clsClose()}, nil,
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
		[]Token{tok(create, "create"), tok(namespaceObject, "aggregate"),tok(quotedName, "ag"), tok(usingDatabase, "db"), tok(forDomain, "dmn"), tok(inContext, "ctx"), apos()}, nil,
	},
}

func TestAggregateStatements (t *testing.T) {
	aggregateStatements.test(t)
}

/*
var eventStatements = []struct {
	dql string;
	ast CreateAggregate
}{
	{
		"<| event 'start' using database 'db' for domain 'dmn' in context 'ctx' within aggregate 'ag' |>",
		CreateEvent{"uuid", "start", NewAggregateNamespace([]string{"db", "dmn", "ctx", "ag"})},
	},
}

/*
var statementsMissingNamespaceVars = []struct {
	dql string;
	error string;
}{
	{"create domain 'dmn';", "database not selected"},
	{"create context 'ctx' using database 'db';", "domain not selected"},
	{"<| value 'address' using database 'db' for domain 'dmn' |>", "context not selected"},
	{"<| event 'start' using database 'db' for domain 'dmn' in context 'ctx' |>", "aggregate not selected"},
};

func TestNamespacesAreValidated(t *testing.T) {
	parser := NewParser();
	for _, statement := range statementsMissingNamespaceVars {
		_, err := parser.Parse(statement.dql);
		if (err != statement.error) {
			t.Error("Invalid statement '"+statement.dql+"' was accepted");
		}
	}
}

var statementsWithGloballySetNamespaces = []struct {
	dql string;

}{
	{"using database 'db'; create domain 'dmn';"},
	{"for domain 'dmn'; create context 'ctx';"},
	{"in context 'ctx'; <| value 'address' |>"},
	{"create aggregate 'agg';"},
	{"within aggregate 'agg'; <| event 'start' |>"},
};

func TestGloballySetNamespaceVarsAreUsedAsDefaults(t *testing.T) {
	parser := NewParser();
	for _, statement := range statementsWithGloballySetNamespaces {
		commands, err := parser.Parse(statement.dql);
		if (err != nil) {
			t.Error(err);
		}
		if (len(commands) != 1) {
			t.Error("GlobalNamespace commands should not return statements,")
		}
	}
}

var tieredNamespaces = []struct {
	dql string;
	namespace1 Namespace;
	namespace2 Namespace;
}{
	{
		`using database 'database1' for domain 'domain1' in context 'context1':{
			create aggregate 'aggregate1';

			using database 'database2' for domain 'domain2' in context 'context2':{
				create aggregate 'aggregate2';
			}
		}`,
		&AggregateNamespace{"database1", "domain1", "context1"},
		&AggregateNamespace{"database2", "domain2", "context2"},
	},{
		`using database 'database1' for domain 'domain1':{
			in context 'context1';
			create aggregate 'aggregate1';

			using domain 'domain2' in context 'context2':{
				create aggregate 'aggregate2';
			}
		}`,
		&AggregateNamespace{"database1", "domain1", "context1"},
		&AggregateNamespace{"database1", "domain2", "context2"},
	},
}

func TestTakesNamespaceVarsFromCurrentNamespace(t *testing.T) {
	parser := NewParser();
	for _, statement := range tieredNamespaces {
		commands, err := parser.Parse(statement.dql);
		if (commands[0].Namespace != statement.namespace1) {
			t.Error(err);
		}
		if (commands[1].Namespace != statement.namespace2) {
			t.Error(err);
		}
	}
}
*/