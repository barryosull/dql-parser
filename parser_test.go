package parser

import (
	"testing"
	"parser/peg"
)

var statements = []struct {
	dql string;
	ast Command
}{
	{
		"create database 'db';",
		CreateDatabase{"uuid", "db"},
	},{
		"create domain 'dmn' using database 'db';",
		CreateDomain{"uuid", "dmn", NewDatabaseNamespace([]string{"db"})},
	},{
		"create context 'ctx' using database 'db' for domain 'dmn';",
		CreateContext{"uuid", "dmn", NewDomainNamespace([]string{"db", "dmn"})},
	},{
		"<| value 'address' using database 'db' for domain 'dmn' in context 'ctx' |>",
		CreateValue{"uuid", "address", NewContextNamespace([]string{"db", "dmn", "ctx"})},
	},{
		"create aggregate 'ag' using database 'db' for domain 'dmn' in context 'ctx';",
		CreateAggregate{"uuid", "ag", NewContextNamespace([]string{"db", "dmn", "ctx"})},
	},{
		"<| event 'start' using database 'db' for domain 'dmn' in context 'ctx' within aggregate 'ag' |>",
		CreateEvent{"uuid", "start", NewAggregateNamespace([]string{"db", "dmn", "ctx", "ag"})},
	},
};

func TestReturnsAst(t *testing.T) {
	parser := peg.NewParser();
	for _, statement := range statements {
		ast, _ := parser.Parse(statement.dql);
		if (ast != statement.ast) {
			t.Error("AST produced from '"+statement.dql+"' is not valid");
		}
	}
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