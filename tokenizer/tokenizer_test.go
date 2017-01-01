package tokenizer

import (
	"testing"
	"strconv"
	tok "parser/token"
)

var dbStatements = testStatements {
	{
		"create database 'db1';",
		[]tok.Token{tok.NewToken(tok.CREATE, "create", 0), tok.NewToken(tok.DATABASE, "database", 7), tok.NewToken(tok.OBJECTNAME, "db1", 17), tok.Semicolon(21)},
	}, {
		"create DATABASE 'db2' ;",
		[]tok.Token{tok.NewToken(tok.CREATE, "create", 0), tok.NewToken(tok.DATABASE, "DATABASE", 7), tok.NewToken(tok.OBJECTNAME, "db2", 17), tok.Semicolon(22)},
	},
};

func TestCreateDatabase(t *testing.T) {
	dbStatements.test(t);
}

var multipleStatements = testStatements{
	{
		"create database 'db1'; create database 'db1';",
		[]tok.Token{tk(tok.CREATE, "create"), tk(tok.DATABASE, "database"), tk(tok.OBJECTNAME, "db1"), semi(), tk(tok.CREATE, "create"), tk(tok.DATABASE, "database"), tk(tok.OBJECTNAME, "db1"), semi()},
	},
}

func TestMultipeStatements(t *testing.T) {
	multipleStatements.test(t);
}

var domainStatements = testStatements{
	{
		"create domain 'dmn' using database 'db';",
		[]tok.Token{tk(tok.CREATE, "create"), tk(tok.DOMAIN, "domain"), tk(tok.OBJECTNAME, "dmn"), tk(tok.USINGDATABASE, "db"), semi()},

	},
	{
		"create domain 'dmn' using database 'db'",
		[]tok.Token{tk(tok.CREATE, "create"), tk(tok.DOMAIN, "domain"), tk(tok.OBJECTNAME, "dmn"), tk(tok.USINGDATABASE, "db")},

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
		[]tok.Token{tk(tok.CREATE, "create"), tk(tok.CONTEXT, "context"), tk(tok.OBJECTNAME, "ctx"), tk(tok.USINGDATABASE, "db"), tk(tok.FORDOMAIN, "dmn"), semi()},
	},
};

func TestCreateContext(t *testing.T) {
	contextStatements.test(t);
}

var valueStatements = testStatements {
	{
		"<| value 'address' using database 'db' for domain 'dmn' in context 'ctx' |>",
		[]tok.Token{clsOpen(), tk(tok.VALUE, "value"), tk(tok.OBJECTNAME, "address"), tk(tok.USINGDATABASE, "db"), tk(tok.FORDOMAIN, "dmn"), tk(tok.INCONTEXT, "ctx"), clsClose()},
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
		[]tok.Token{tk(tok.CREATE, "create"), tk(tok.AGGREGATE, "aggregate"),tk(tok.OBJECTNAME, "ag"), tk(tok.USINGDATABASE, "db"), tk(tok.FORDOMAIN, "dmn"), tk(tok.INCONTEXT, "ctx"), semi()},
	},
}

func TestAggregateStatements (t *testing.T) {
	aggregateStatements.test(t)
}


var eventStatements = testStatements{
	{
		"<| event 'start' using database 'db' for domain 'dmn' in context 'ctx' within aggregate 'ag' |>",
		[]tok.Token{clsOpen(), tk(tok.EVENT, "event"), tk(tok.OBJECTNAME, "start"), tk(tok.USINGDATABASE, "db"), tk(tok.FORDOMAIN, "dmn"), tk(tok.INCONTEXT, "ctx"), tk(tok.WITHINAGGREGATE, "ag"), clsClose()},
	},
}

func TestEventStatements (t *testing.T) {
	eventStatements.test(t)
}

var statementsWithGloballySetNamespaces = testStatements {
	{
		"using database 'db'; create domain 'dmn';",
		[]tok.Token{tk(tok.USINGDATABASE, "db"), semi(), tk(tok.CREATE, "create"), tk(tok.DOMAIN, "domain"), tk(tok.OBJECTNAME, "dmn"), semi()},
	},
	{
		"for domain 'dmn'; create context 'ctx';",
		[]tok.Token{tk(tok.FORDOMAIN, "dmn"), semi(), tk(tok.CREATE, "create"), tk(tok.CONTEXT, "context"), tk(tok.OBJECTNAME, "ctx"), semi()},
	},
	{
		"in context 'ctx'; <| value 'address' |>",
		[]tok.Token{tk(tok.INCONTEXT, "ctx"), semi(), clsOpen(), tk(tok.VALUE, "value"), tk(tok.OBJECTNAME, "address"), clsClose()},
	},
	{
		"within aggregate 'agg'; <| event 'start' |>",
		[]tok.Token{tk(tok.WITHINAGGREGATE, "agg"), semi(), clsOpen(), tk(tok.EVENT, "event"), tk(tok.OBJECTNAME, "start"), clsClose()},
	},
};

func TestGloballySetNamespace (t *testing.T) {
	statementsWithGloballySetNamespaces.test(t)
}

var createObjectTypes = testStatements {
	{
		"<| entity 'ent' |>",
		[]tok.Token{clsOpen(), tk(tok.ENTITY, "entity"), tk(tok.OBJECTNAME, "ent"), clsClose()},
	},
	{
		"<| entity 'ent' CHECK ( return value != 0;) |>",
		[]tok.Token{
			clsOpen(),
			tk(tok.ENTITY, "entity"),
			tk(tok.OBJECTNAME, "ent"),

			tk(tok.CHECK, "CHECK"),
			tk(tok.LPAREN, "("),

			tk(tok.RETURN, "return"),
			tk(tok.IDENT, "value"),
			tk(tok.NOTEQ, "!="),
			tk(tok.INTEGER, "0"),
			tk(tok.SEMICOLON, ";"),

			tk(tok.RPAREN, ")"),
			clsClose(),
		},
	},
	{
		"<| projection 'proj' |>",
		[]tok.Token{clsOpen(), tk(tok.PROJECTION, "projection"), tk(tok.OBJECTNAME, "proj"), clsClose()},
	},
	{
		"<| invariant 'invar' on 'projection\\quote' |>",
		[]tok.Token{clsOpen(), tk(tok.INVARIANT, "invariant"), tk(tok.OBJECTNAME, "invar"), tk(tok.ON, "on"), tk(tok.OBJECTNAME, "projection\\quote"), clsClose()},
	},
	{
		"<| command 'cmd' |>",
		[]tok.Token{clsOpen(), tk(tok.COMMAND, "command"), tk(tok.OBJECTNAME, "cmd"), clsClose()},
	},
	{
		"<| query 'qry' |>",
		[]tok.Token{clsOpen(), tk(tok.QUERY, "query"), tk(tok.OBJECTNAME, "qry"), clsClose()},
	},

}

func TestObjectTypes(t *testing.T) {
	createObjectTypes.test(t)
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
			tk(tok.AGGREGATE, "aggregate"),
			tk(tok.OBJECTNAME, "aggregate1"),
			tk(tok.SEMICOLON, ";"),

			tk(tok.USINGDATABASE, "database2"),
			tk(tok.FORDOMAIN, "domain2"),
			tk(tok.INCONTEXT, "context2"),
			tk(tok.COLON, ":"),
			tk(tok.LBRACE, "{"),

			tk(tok.CREATE, "create"),
			tk(tok.AGGREGATE, "aggregate"),
			tk(tok.OBJECTNAME, "aggregate2"),
			tk(tok.SEMICOLON, ";"),

			tk(tok.RBRACE, "}"),
			tk(tok.RBRACE, "}"),
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
		[]tok.Token{
			tk(tok.PROPERTIES, "properties"),
			tk(tok.LBRACE, "{"),

			tk(tok.OBJECTNAME, "value\\service_charge"),
			tk(tok.IDENT, "service_charge"),
			tk(tok.ASSIGN, "="),
			tk(tok.OBJECTNAME, "value\\service_charge"),
			tk(tok.LPAREN, "("),
			tk(tok.INTEGER, "1"),
			tk(tok.RPAREN, ")"),
			tk(tok.SEMICOLON, ";"),

			tk(tok.OBJECTNAME, "value\\category"),
			tk(tok.IDENT, "category"),
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
			tk(tok.IDENT, "value"),
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
			tk(tok.IDENT, "doThing"),
			tk(tok.LPAREN, "("),
			tk(tok.RPAREN, ")"),
			tk(tok.LBRACE, "{"),
			tk(tok.IDENT, "a"),
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
			tk(tok.IDENT, "doThing2"),
			tk(tok.LPAREN, "("),
			tk(tok.OBJECTNAME, "value\\service-charge"),
			tk(tok.IDENT, "service_charge"),
			tk(tok.COMMA, ","),
			tk(tok.OBJECTNAME, "value\\category"),
			tk(tok.IDENT, "category"),
			tk(tok.COMMA, ","),
			tk(tok.OBJECTNAME, "string"),
			tk(tok.IDENT, "title"),
			tk(tok.COMMA, ","),
			tk(tok.OBJECTNAME, "integer"),
			tk(tok.IDENT, "int"),
			tk(tok.COMMA, ","),
			tk(tok.OBJECTNAME, "float"),
			tk(tok.IDENT, "flt"),
			tk(tok.COMMA, ","),
			tk(tok.OBJECTNAME, "boolean"),
			tk(tok.IDENT, "bl"),
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
			tk(tok.OBJECTNAME, "is-started"),
			tk(tok.SEMICOLON, ";"),
			tk(tok.IDENT, "revision"),
			tk(tok.ASSIGN, "="),
			tk(tok.RUNQUERY, "run query"),
			tk(tok.OBJECTNAME, "next-revision-number"),
			tk(tok.LPAREN, "("),
			tk(tok.IDENT, "agency_id"),
			tk(tok.COMMA, ","),
			tk(tok.IDENT, "quote_number"),
			tk(tok.RPAREN, ")"),
			tk(tok.SEMICOLON, ";"),
			tk(tok.APPLYEVENT, "apply event"),
			tk(tok.OBJECTNAME, "started"),
			tk(tok.LPAREN, "("),
			tk(tok.IDENT, "agency_id"),
			tk(tok.COMMA, ","),
			tk(tok.IDENT, "brand_id"),
			tk(tok.COMMA, ","),
			tk(tok.IDENT, "quote_number"),
			tk(tok.COMMA, ","),
			tk(tok.IDENT, "revision"),
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
			tk(tok.IDENT, "agency_id"),
			tk(tok.ASSIGN, "="),
			tk(tok.IDENT, "event"),
			tk(tok.ARROW, "->"),
			tk(tok.IDENT, "agency_id"),
			tk(tok.SEMICOLON, ";"),
			tk(tok.IDENT, "is_started"),
			tk(tok.ASSIGN, "="),
			tk(tok.BOOLEAN, "true"),
			tk(tok.SEMICOLON, ";"),
			tk(tok.RBRACE, "}"),
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
		[]tok.Token{
			tk(tok.MINUS, "-"),
			tk(tok.MINUS, "-"),
			tk(tok.IDENT, "a"),
			tk(tok.IDENT, "a"),
			tk(tok.PLUS, "+"),
			tk(tok.PLUS, "+"),
			tk(tok.IDENT, "a"),
			tk(tok.LTOREQ, "<="),
			tk(tok.IDENT, "b"),
			tk(tok.IDENT, "b"),
			tk(tok.GTOREQ, ">="),
			tk(tok.IDENT, "a"),
		},
	},
	{
		"a + b - c",
		[]tok.Token{
			tk(tok.IDENT, "a"),
			tk(tok.PLUS, "+"),
			tk(tok.IDENT, "b"),
			tk(tok.MINUS, "-"),
			tk(tok.IDENT, "c"),
		},
	},
	{
		"a + (a - b) % a",
		[]tok.Token{
			tk(tok.IDENT, "a"),
			tk(tok.PLUS, "+"),
			tk(tok.LPAREN, "("),
			tk(tok.IDENT, "a"),
			tk(tok.MINUS, "-"),
			tk(tok.IDENT, "b"),
			tk(tok.RPAREN, ")"),
			tk(tok.REMAINDER, "%"),
			tk(tok.IDENT, "a"),
		},
	},
	{
		"a->b->c + a->b() - !b and a == b and a < b or a > b ",
		[]tok.Token{
			tk(tok.IDENT, "a"),
			tk(tok.ARROW, "->"),
			tk(tok.IDENT, "b"),
			tk(tok.ARROW, "->"),
			tk(tok.IDENT, "c"),
			tk(tok.PLUS, "+"),
			tk(tok.IDENT, "a"),
			tk(tok.ARROW, "->"),
			tk(tok.IDENT, "b"),
			tk(tok.LPAREN, "("),
			tk(tok.RPAREN, ")"),
			tk(tok.MINUS, "-"),
			tk(tok.BANG, "!"),
			tk(tok.IDENT, "b"),
			tk(tok.AND, "and"),
			tk(tok.IDENT, "a"),
			tk(tok.EQ, "=="),
			tk(tok.IDENT, "b"),
			tk(tok.AND, "and"),
			tk(tok.IDENT, "a"),
			tk(tok.LT, "<"),
			tk(tok.IDENT, "b"),
			tk(tok.OR, "or"),
			tk(tok.IDENT, "a"),
			tk(tok.GT, ">"),
			tk(tok.IDENT, "b"),
		},
	},{
		"a = andrew",
		[]tok.Token {
			tk(tok.IDENT, "a"),
			tk(tok.ASSIGN, "="),
			tk(tok.IDENT, "andrew"),
		},
	},{
		"clarkKent = 'value\\isSuperman'(false)",
		[]tok.Token{
			tk(tok.IDENT, "clarkKent"),
			tk(tok.ASSIGN, "="),
			tk(tok.OBJECTNAME, "value\\isSuperman"),
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
			tk(tok.IDENT, "a"),
			tk(tok.RPAREN, ")"),
			tk(tok.LBRACE, "{"),
			tk(tok.IDENT, "a"),
			tk(tok.SEMICOLON, ";"),
			tk(tok.RBRACE, "}"),

			tk(tok.ELSEIF, "else if"),
			tk(tok.LPAREN, "("),
			tk(tok.IDENT, "b"),
			tk(tok.RPAREN, ")"),
			tk(tok.LBRACE, "{"),
			tk(tok.IDENT, "a"),
			tk(tok.SEMICOLON, ";"),
			tk(tok.RBRACE, "}"),

			tk(tok.ELSE, "else"),
			tk(tok.LBRACE, "{"),
			tk(tok.IDENT, "b"),
			tk(tok.SEMICOLON, ";"),
			tk(tok.RBRACE, "}"),

			tk(tok.FOREACH, "foreach"),
			tk(tok.LPAREN, "("),
			tk(tok.IDENT, "a"),
			tk(tok.ARROW, "->"),
			tk(tok.IDENT, "b"),
			tk(tok.LPAREN, "("),
			tk(tok.RPAREN, ")"),
			tk(tok.AS, "as"),
			tk(tok.IDENT, "b"),
			tk(tok.STRONGARROW, "=>"),
			tk(tok.IDENT, "c"),
			tk(tok.RPAREN, ")"),

			tk(tok.LBRACE, "{"),
			tk(tok.IDENT, "a"),
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
			tk(tok.IDENT, "database"),
			tk(tok.IDENT, "domain"),
			tk(tok.IDENT, "context"),
			tk(tok.IDENT, "aggregate"),
			tk(tok.IDENT, "value"),
			tk(tok.IDENT, "event"),
			tk(tok.IDENT, "entity"),
			tk(tok.IDENT, "command"),
			tk(tok.IDENT, "projection"),
			tk(tok.IDENT, "invariant"),
			tk(tok.IDENT, "query"),
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
			tk(tok.IDENT, "propertiesA"),
			tk(tok.IDENT, "checkA"),
			tk(tok.IDENT, "handlerA"),
			tk(tok.IDENT, "functionA"),
			tk(tok.IDENT, "whenA"),
			tk(tok.IDENT, "andA"),
			tk(tok.IDENT, "orA"),
			tk(tok.IDENT, "ifA"),
			tk(tok.IDENT, "elseA"),
			tk(tok.IDENT, "returnA"),
			tk(tok.IDENT, "foreachA"),
			tk(tok.IDENT, "asA"),
			tk(tok.IDENT, "createA"),
			tk(tok.IDENT, "nullA"),
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
