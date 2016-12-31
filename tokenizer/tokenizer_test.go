package tokenizer

import (
	"testing"
	"strconv"
	tok "parser/token"
)

var dbStatements = testStatements {
	{
		"create database 'db1';",
		[]tok.Token{tok.NewToken(tok.CREATE, "create", 0), tok.NewToken(tok.NAMESPACEOBJECT, "database", 7), tok.NewToken(tok.QUOTEDNAME, "db1", 17), tok.Semicolon(21)},
	}, {
		"create database 'db2' ;",
		[]tok.Token{tok.NewToken(tok.CREATE, "create", 0), tok.NewToken(tok.NAMESPACEOBJECT, "database", 7), tok.NewToken(tok.QUOTEDNAME, "db2", 17), tok.Semicolon(22)},
	},
};

func TestCreateDatabase(t *testing.T) {
	dbStatements.test(t);
}

var multipleStatements = testStatements{
	{
		"create database 'db1'; create database 'db1';",
		[]tok.Token{tk(tok.CREATE, "create"), tk(tok.NAMESPACEOBJECT, "database"), tk(tok.QUOTEDNAME, "db1"), semi(), tk(tok.CREATE, "create"), tk(tok.NAMESPACEOBJECT, "database"), tk(tok.QUOTEDNAME, "db1"), semi()},
	},
}

func TestMultipeStatements(t *testing.T) {
	multipleStatements.test(t);
}

var domainStatements = testStatements{
	{
		"create domain 'dmn' using database 'db';",
		[]tok.Token{tk(tok.CREATE, "create"), tk(tok.NAMESPACEOBJECT, "domain"), tk(tok.QUOTEDNAME, "dmn"), tk(tok.USINGDATABASE, "db"), semi()},

	},
	{
		"create domain 'dmn' using database 'db'",
		[]tok.Token{tk(tok.CREATE, "create"), tk(tok.NAMESPACEOBJECT, "domain"), tk(tok.QUOTEDNAME, "dmn"), tk(tok.USINGDATABASE, "db")},

	},
};

func tk(typ tok.TokenType, val string) tok.Token {
	return tok.Token{typ, val, tok.IgnoreTokenPos};
}

func semi() tok.Token {
	return tok.Semicolon(tok.IgnoreTokenPos);
}

func TestCreateDomain(t *testing.T) {
	domainStatements.test(t);
}


var contextStatements = testStatements {
	{
		"create context 'ctx' using database 'db' for domain 'dmn';",
		[]tok.Token{tk(tok.CREATE, "create"), tk(tok.NAMESPACEOBJECT, "context"), tk(tok.QUOTEDNAME, "ctx"), tk(tok.USINGDATABASE, "db"), tk(tok.FORDOMAIN, "dmn"), semi()},
	},
};

func TestCreateContext(t *testing.T) {
	contextStatements.test(t);
}

var valueStatements = testStatements {
	{
		"<| value 'address' using database 'db' for domain 'dmn' in context 'ctx' |>",
		[]tok.Token{clsOpen(), tk(tok.CLASS, "value"), tk(tok.QUOTEDNAME, "address"), tk(tok.USINGDATABASE, "db"), tk(tok.FORDOMAIN, "dmn"), tk(tok.INCONTEXT, "ctx"), clsClose()},
	},
}

func clsOpen() tok.Token {
	return tok.ClsOpen(tok.IgnoreTokenPos);
}

func clsClose() tok.Token {
	return tok.ClsClose(tok.IgnoreTokenPos);
}

func TestCreateValue(t *testing.T) {
	valueStatements.test(t);
}

var aggregateStatements = testStatements{
	{
		"create aggregate 'ag' using database 'db' for domain 'dmn' in context 'ctx';",
		[]tok.Token{tk(tok.CREATE, "create"), tk(tok.NAMESPACEOBJECT, "aggregate"),tk(tok.QUOTEDNAME, "ag"), tk(tok.USINGDATABASE, "db"), tk(tok.FORDOMAIN, "dmn"), tk(tok.INCONTEXT, "ctx"), semi()},
	},
}

func TestAggregateStatements (t *testing.T) {
	aggregateStatements.test(t)
}


var eventStatements = testStatements{
	{
		"<| event 'start' using database 'db' for domain 'dmn' in context 'ctx' within aggregate 'ag' |>",
		[]tok.Token{clsOpen(), tk(tok.CLASS, "event"), tk(tok.QUOTEDNAME, "start"), tk(tok.USINGDATABASE, "db"), tk(tok.FORDOMAIN, "dmn"), tk(tok.INCONTEXT, "ctx"), tk(tok.WITHINAGGREGATE, "ag"), clsClose()},
	},
}

func TestEventStatements (t *testing.T) {
	eventStatements.test(t)
}

var statementsWithGloballySetNamespaces = testStatements {
	{
		"using database 'db'; create domain 'dmn';",
		[]tok.Token{tk(tok.USINGDATABASE, "db"), semi(), tk(tok.CREATE, "create"), tk(tok.NAMESPACEOBJECT, "domain"), tk(tok.QUOTEDNAME, "dmn"), semi()},
	},
	{
		"for domain 'dmn'; create context 'ctx';",
		[]tok.Token{tk(tok.FORDOMAIN, "dmn"), semi(), tk(tok.CREATE, "create"), tk(tok.NAMESPACEOBJECT, "context"), tk(tok.QUOTEDNAME, "ctx"), semi()},
	},
	{
		"in context 'ctx'; <| value 'address' |>",
		[]tok.Token{tk(tok.INCONTEXT, "ctx"), semi(), clsOpen(), tk(tok.CLASS, "value"), tk(tok.QUOTEDNAME, "address"), clsClose()},
	},
	{
		"within aggregate 'agg'; <| event 'start' |>",
		[]tok.Token{tk(tok.WITHINAGGREGATE, "agg"), semi(), clsOpen(), tk(tok.CLASS, "event"), tk(tok.QUOTEDNAME, "start"), clsClose()},
	},
};

func TestGloballySetNamespace (t *testing.T) {
	statementsWithGloballySetNamespaces.test(t)
}

var objectTypes = testStatements {
	{
		"<| entity 'ent' |>",
		[]tok.Token{clsOpen(), tk(tok.CLASS, "entity"), tk(tok.QUOTEDNAME, "ent"), clsClose()},
	},
	{
		"<| entity 'ent' check ( return value != 0;) |>",
		[]tok.Token{
			clsOpen(),
			tk(tok.CLASS, "entity"),
			tk(tok.QUOTEDNAME, "ent"),

			tk(tok.CHECK, "check"),
			tk(tok.LPAREN, "("),

			tk(tok.RETURN, "return"),
			tk(tok.IDENTIFIER, "value"),
			tk(tok.NOTEQ, "!="),
			tk(tok.INTEGER, "0"),
			tk(tok.SEMICOLON, ";"),

			tk(tok.RPAREN, ")"),
			clsClose(),
		},
	},
	{
		"<| invariant 'invar' |>",
		[]tok.Token{clsOpen(), tk(tok.CLASS, "invariant"), tk(tok.QUOTEDNAME, "invar"), clsClose()},
	},
	{
		"<| command 'cmd' |>",
		[]tok.Token{clsOpen(), tk(tok.CLASS, "command"), tk(tok.QUOTEDNAME, "cmd"), clsClose()},
	},
	{
		"<| query 'qry' |>",
		[]tok.Token{clsOpen(), tk(tok.CLASS, "query"), tk(tok.QUOTEDNAME, "qry"), clsClose()},
	},
	{
		"<| projection 'proj' |>",
		[]tok.Token{clsOpen(), tk(tok.CLASS, "projection"), tk(tok.QUOTEDNAME, "proj"), clsClose()},
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
		[]tok.Token{
			tk(tok.USINGDATABASE, "database1"),
			tk(tok.FORDOMAIN, "domain1"),
			tk(tok.INCONTEXT, "context1"),
			tk(tok.COLON, ":"),
			tk(tok.LBRACE, "{"),

			tk(tok.CREATE, "create"),
			tk(tok.NAMESPACEOBJECT, "aggregate"),
			tk(tok.QUOTEDNAME, "aggregate1"),
			tk(tok.SEMICOLON, ";"),

			tk(tok.USINGDATABASE, "database2"),
			tk(tok.FORDOMAIN, "domain2"),
			tk(tok.INCONTEXT, "context2"),
			tk(tok.COLON, ":"),
			tk(tok.LBRACE, "{"),

			tk(tok.CREATE, "create"),
			tk(tok.NAMESPACEOBJECT, "aggregate"),
			tk(tok.QUOTEDNAME, "aggregate2"),
			tk(tok.SEMICOLON, ";"),

			tk(tok.RBRACE, "}"),
			tk(tok.RBRACE, "}"),
		},
	},
};

func TestNamespaceBlocks (t *testing.T) {
	namespaceBlocks.test(t)
}

var CLASSComponents = testStatements{
	{
		`
		properties
		{
			value\service_charge service_charge = 'value\service_charge'(1);
			value\category category = [];
		}`,
		[]tok.Token{
			tk(tok.PROPERTIES, "properties"),
			tk(tok.LBRACE, "{"),

			tk(tok.TYPEREF, "value\\service_charge"),
			tk(tok.IDENTIFIER, "service_charge"),
			tk(tok.ASSIGN, "="),
			tk(tok.QUOTEDNAME, "value\\service_charge"),
			tk(tok.LPAREN, "("),
			tk(tok.INTEGER, "1"),
			tk(tok.RPAREN, ")"),
			tk(tok.SEMICOLON, ";"),

			tk(tok.TYPEREF, "value\\category"),
			tk(tok.IDENTIFIER, "category"),
			tk(tok.ASSIGN, "="),
			tk(tok.LBRACKET, "["),
			tk(tok.RBRACKET, "]"),
			tk(tok.SEMICOLON, ";"),

			tk(tok.RBRACE, "}"),
		},
	},
	{
		`
		check
		(
			return value != 0;
		)`,
		[]tok.Token{
			tk(tok.CHECK, "check"),
			tk(tok.LPAREN, "("),

			tk(tok.RETURN, "return"),
			tk(tok.IDENTIFIER, "value"),
			tk(tok.NOTEQ, "!="),
			tk(tok.INTEGER, "0"),
			tk(tok.SEMICOLON, ";"),

			tk(tok.RPAREN, ")"),
		},
	},
	{
		`
		function doThing()
		{
			a = 2.1;
		}`,
		[]tok.Token{
			tk(tok.FUNCTION, "function"),
			tk(tok.IDENTIFIER, "doThing"),
			tk(tok.LPAREN, "("),
			tk(tok.RPAREN, ")"),
			tk(tok.LBRACE, "{"),
			tk(tok.IDENTIFIER, "a"),
			tk(tok.ASSIGN, "="),
			tk(tok.FLOAT, "2.1"),
			tk(tok.SEMICOLON, ";"),
			tk(tok.RBRACE, "}"),
		},
	},
	{
		`
		function doThing2(value\service-charge service_charge, value\category category, string title, integer int, float flt, boolean bl)
		{

		}`,
		[]tok.Token{
			tk(tok.FUNCTION, "function"),
			tk(tok.IDENTIFIER, "doThing2"),
			tk(tok.LPAREN, "("),
			tk(tok.TYPEREF, "value\\service-charge"),
			tk(tok.IDENTIFIER, "service_charge"),
			tk(tok.COMMA, ","),
			tk(tok.TYPEREF, "value\\category"),
			tk(tok.IDENTIFIER, "category"),
			tk(tok.COMMA, ","),
			tk(tok.TYPEREF, "string"),
			tk(tok.IDENTIFIER, "title"),
			tk(tok.COMMA, ","),
			tk(tok.TYPEREF, "integer"),
			tk(tok.IDENTIFIER, "int"),
			tk(tok.COMMA, ","),
			tk(tok.TYPEREF, "float"),
			tk(tok.IDENTIFIER, "flt"),
			tk(tok.COMMA, ","),
			tk(tok.TYPEREF, "boolean"),
			tk(tok.IDENTIFIER, "bl"),
			tk(tok.RPAREN, ")"),
			tk(tok.LBRACE, "{"),
			tk(tok.RBRACE, "}"),
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
		[]tok.Token{
			tk(tok.HANDLER, "handler"),
			tk(tok.LBRACE, "{"),
			tk(tok.ASSERTINVARIANT, "assert  invariant"),
			tk(tok.NOT, "not"),
			tk(tok.QUOTEDNAME, "is-started"),
			tk(tok.SEMICOLON, ";"),
			tk(tok.IDENTIFIER, "revision"),
			tk(tok.ASSIGN, "="),
			tk(tok.RUNQUERY, "run query"),
			tk(tok.QUOTEDNAME, "next-revision-number"),
			tk(tok.LPAREN, "("),
			tk(tok.IDENTIFIER, "agency_id"),
			tk(tok.COMMA, ","),
			tk(tok.IDENTIFIER, "quote_number"),
			tk(tok.RPAREN, ")"),
			tk(tok.SEMICOLON, ";"),
			tk(tok.APPLYEVENT, "apply event"),
			tk(tok.QUOTEDNAME, "started"),
			tk(tok.LPAREN, "("),
			tk(tok.IDENTIFIER, "agency_id"),
			tk(tok.COMMA, ","),
			tk(tok.IDENTIFIER, "brand_id"),
			tk(tok.COMMA, ","),
			tk(tok.IDENTIFIER, "quote_number"),
			tk(tok.COMMA, ","),
			tk(tok.IDENTIFIER, "revision"),
			tk(tok.RPAREN, ")"),
			tk(tok.SEMICOLON, ";"),
			tk(tok.RBRACE, "}"),

		},
	},
	{
		`
		WHEN event 'started'
		{
			agency_id = event->agency_id;
			is_started = true;
		}`,
		[]tok.Token{
			tk(tok.WHENEVENT, "started"),
			tk(tok.LBRACE, "{"),
			tk(tok.IDENTIFIER, "agency_id"),
			tk(tok.ASSIGN, "="),
			tk(tok.IDENTIFIER, "event"),
			tk(tok.ARROW, "->"),
			tk(tok.IDENTIFIER, "agency_id"),
			tk(tok.SEMICOLON, ";"),
			tk(tok.IDENTIFIER, "is_started"),
			tk(tok.ASSIGN, "="),
			tk(tok.BOOLEAN, "true"),
			tk(tok.SEMICOLON, ";"),
			tk(tok.RBRACE, "}"),
		},

	},
};

func TestClassComponents (t *testing.T) {
	CLASSComponents.test(t)
}

var expressions = testStatements{
	{
		`--a
		a++
		a <= b
		b >= a`,
		[]tok.Token{
			tk(tok.MINUS, "-"),
			tk(tok.MINUS, "-"),
			tk(tok.IDENTIFIER, "a"),
			tk(tok.IDENTIFIER, "a"),
			tk(tok.PLUS, "+"),
			tk(tok.PLUS, "+"),
			tk(tok.IDENTIFIER, "a"),
			tk(tok.LTOREQ, "<="),
			tk(tok.IDENTIFIER, "b"),
			tk(tok.IDENTIFIER, "b"),
			tk(tok.GTOREQ, ">="),
			tk(tok.IDENTIFIER, "a"),
		},
	},
	{
		"a + b - c",
		[]tok.Token{
			tk(tok.IDENTIFIER, "a"),
			tk(tok.PLUS, "+"),
			tk(tok.IDENTIFIER, "b"),
			tk(tok.MINUS, "-"),
			tk(tok.IDENTIFIER, "c"),
		},
	},
	{
		"a + (a - b)",
		[]tok.Token{
			tk(tok.IDENTIFIER, "a"),
			tk(tok.PLUS, "+"),
			tk(tok.LPAREN, "("),
			tk(tok.IDENTIFIER, "a"),
			tk(tok.MINUS, "-"),
			tk(tok.IDENTIFIER, "b"),
			tk(tok.RPAREN, ")"),
		},
	},
	{
		"a->b->c + a->b() - !b and a == b and a < b or a > b ",
		[]tok.Token{
			tk(tok.IDENTIFIER, "a"),
			tk(tok.ARROW, "->"),
			tk(tok.IDENTIFIER, "b"),
			tk(tok.ARROW, "->"),
			tk(tok.IDENTIFIER, "c"),
			tk(tok.PLUS, "+"),
			tk(tok.IDENTIFIER, "a"),
			tk(tok.ARROW, "->"),
			tk(tok.IDENTIFIER, "b"),
			tk(tok.LPAREN, "("),
			tk(tok.RPAREN, ")"),
			tk(tok.MINUS, "-"),
			tk(tok.BANG, "!"),
			tk(tok.IDENTIFIER, "b"),
			tk(tok.AND, "and"),
			tk(tok.IDENTIFIER, "a"),
			tk(tok.EQ, "=="),
			tk(tok.IDENTIFIER, "b"),
			tk(tok.AND, "and"),
			tk(tok.IDENTIFIER, "a"),
			tk(tok.LT, "<"),
			tk(tok.IDENTIFIER, "b"),
			tk(tok.OR, "or"),
			tk(tok.IDENTIFIER, "a"),
			tk(tok.GT, ">"),
			tk(tok.IDENTIFIER, "b"),
		},
	},{
		"a = andrew",
		[]tok.Token {
			tk(tok.IDENTIFIER, "a"),
			tk(tok.ASSIGN, "="),
			tk(tok.IDENTIFIER, "andrew"),
		},
	},{
		"clarkKent = 'value\\isSuperman'(false)",
		[]tok.Token{
			tk(tok.IDENTIFIER, "clarkKent"),
			tk(tok.ASSIGN, "="),
			tk(tok.QUOTEDNAME, "value\\isSuperman"),
			tk(tok.LPAREN, "("),
			tk(tok.BOOLEAN, "false"),
			tk(tok.RPAREN, ")"),
		},
	},{
		`"string value"`,
		[]tok.Token{
			tk(tok.STRING, "string value"),
		},
	},{
		`null`,
		[]tok.Token{
			tk(tok.NULL, "null"),
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
		[]tok.Token{
			tk(tok.IF, "if"),
			tk(tok.LPAREN, "("),
			tk(tok.IDENTIFIER, "a"),
			tk(tok.RPAREN, ")"),
			tk(tok.LBRACE, "{"),
			tk(tok.IDENTIFIER, "a"),
			tk(tok.SEMICOLON, ";"),
			tk(tok.RBRACE, "}"),

			tk(tok.ELSEIF, "else if"),
			tk(tok.LPAREN, "("),
			tk(tok.IDENTIFIER, "b"),
			tk(tok.RPAREN, ")"),
			tk(tok.LBRACE, "{"),
			tk(tok.IDENTIFIER, "a"),
			tk(tok.SEMICOLON, ";"),
			tk(tok.RBRACE, "}"),

			tk(tok.ELSE, "else"),
			tk(tok.LBRACE, "{"),
			tk(tok.IDENTIFIER, "b"),
			tk(tok.SEMICOLON, ";"),
			tk(tok.RBRACE, "}"),

			tk(tok.FOREACH, "foreach"),
			tk(tok.LPAREN, "("),
			tk(tok.IDENTIFIER, "a"),
			tk(tok.ARROW, "->"),
			tk(tok.IDENTIFIER, "b"),
			tk(tok.LPAREN, "("),
			tk(tok.RPAREN, ")"),
			tk(tok.AS, "as"),
			tk(tok.IDENTIFIER, "b"),
			tk(tok.STRONGARROW, "=>"),
			tk(tok.IDENTIFIER, "c"),
			tk(tok.RPAREN, ")"),

			tk(tok.LBRACE, "{"),
			tk(tok.IDENTIFIER, "a"),
			tk(tok.SEMICOLON, ";"),
			tk(tok.RBRACE, "}"),
		},
	},
}

func TestStatements(t *testing.T) {
	statements.test(t)
}

// These keywords should be seen as expressions, NOT keywords, dependent on context
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
		[]tok.Token{
			tk(tok.IDENTIFIER, "database"),
			tk(tok.IDENTIFIER, "domain"),
			tk(tok.IDENTIFIER, "context"),
			tk(tok.IDENTIFIER, "aggregate"),
			tk(tok.IDENTIFIER, "value"),
			tk(tok.IDENTIFIER, "event"),
			tk(tok.IDENTIFIER, "entity"),
			tk(tok.IDENTIFIER, "command"),
			tk(tok.IDENTIFIER, "projection"),
			tk(tok.IDENTIFIER, "invariant"),
			tk(tok.IDENTIFIER, "query"),
		},
	},
}

func TestKeywordsAsExpressions(t *testing.T) {
	keyWordsAsExpressions.test(t)
}

// These keywords can be used in expressions only if they're part of an IDENTIFIER
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
		createA
		nullA`,
		[]tok.Token {
			tk(tok.IDENTIFIER, "propertiesA"),
			tk(tok.IDENTIFIER, "checkA"),
			tk(tok.IDENTIFIER, "handlerA"),
			tk(tok.IDENTIFIER, "functionA"),
			tk(tok.IDENTIFIER, "whenA"),
			tk(tok.IDENTIFIER, "andA"),
			tk(tok.IDENTIFIER, "orA"),
			tk(tok.IDENTIFIER, "ifA"),
			tk(tok.IDENTIFIER, "elseA"),
			tk(tok.IDENTIFIER, "returnA"),
			tk(tok.IDENTIFIER, "foreachA"),
			tk(tok.IDENTIFIER, "asA"),
			tk(tok.IDENTIFIER, "createA"),
			tk(tok.IDENTIFIER, "nullA"),
		},
	},
}

func TestKeywordsInExpressions(t *testing.T) {
	keywordsInExpressions.test(t)
}

var badStatements = []struct{
	dql string
	err tok.Error
}{
	{
		"create dbase",
		tok.Error{"create dbase", 7, "database, domain, context, aggregate", "dbase", },
	}, {
		"create dmain",
		tok.Error{"create dmain", 7, "database, domain, context, aggregate", "dmain", },
	}, {
		"create cntext",
		tok.Error{"create cntext", 7, "database, domain, context, aggregate", "cntext", },
	},{
		"create aggrege",
		tok.Error{"create aggrege", 7, "database, domain, context, aggregate", "aggrege",},
	},{
		"using dbase",
		tok.Error{"using dbase", 6, "database", "dbase",},
	},{
		"for dom",
		tok.Error{"for dom", 4, "domain", "dom",},
	},{
		"in cntext",
		tok.Error{"in cntext", 3, "context", "cntext",},
	},{
		"within agg",
		tok.Error{"within agg", 7, "aggregate", "agg",},
	},{
		"assert invar",
		tok.Error{"assert invar", 0, "assert invariant", "assert invar",},
	},{
		"run qry",
		tok.Error{"run qry", 0, "run query", "run qry",},
	},{
		"apply evt",
		tok.Error{"apply evt", 0, "apply event", "apply evt",},
	},{
		"when evt",
		tok.Error{"when evt", 5, "event", "evt",},
	},{
		"<| valu ",
		tok.Error{"<| valu ", 3, "value, entity, event, command, query, invariant, projection", "valu",},
	},{
		"for domain ",
		tok.Error{"for domain ", 11, "'", "EOF",},
	},{
		"for domain '",
		tok.Error{"for domain '", 12, "'", "EOF",},
	},{
		"<| value ''",
		tok.Error{"<| value ''", 10, "value name", "empty name",},
	},{
		"~",
		tok.Error{"~", 0, "keyword", "~",},
	},
}

func TestBadStatements(t *testing.T){
	for _, statement := range badStatements {
		tokenizer := NewTokenizer(statement.dql);

		var token *tok.Token
		var err *tok.Error
		for {
			token, err = tokenizer.Next()
			if (token == nil) {
				break;
			}
		}
		if (err == nil) {
			t.Error("No error found in DQL statement '"+statement.dql+"'")
			t.Error(tokenizer.Tokens())
		} else if (!err.Equals(statement.err)) {
			t.Error("Error found in DQL statement '"+statement.dql+"' does not match expected")
			t.Error("Expected: "+statement.err.String())
			t.Error("Actual: "+err.String())
		}
	}
}

type testStatement struct {
	dql string;
	expected []tok.Token;
}

type testStatements []testStatement

func (statements testStatements) test(t *testing.T) {

	for _, statement := range statements {
		tokenizer := NewTokenizer(statement.dql);

		var token *tok.Token
		var actual []tok.Token
		var err *tok.Error
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

func compareTokenLists(expected, actual []tok.Token, dql string, t *testing.T) {
	if (len(expected) != len(actual)) {
		t.Error("Error with AST produced from '"+dql+"'");
		t.Error("Number of tokens are mismtached, expected "+strconv.Itoa(len(expected))+", got "+strconv.Itoa(len(actual)));
	}

	for i, token := range expected {
		if i == len(actual) {
			t.Error("Expected: "+token.String())
			t.Error("Got: Nothing")
			return
		}
		if (!token.Compare(actual[i])) {
			t.Error("Expected: "+token.String())
			t.Error("Got: "+actual[i].String())
			return
		}
	}
}
