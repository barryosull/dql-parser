package namespace

import (
	"testing"
)

type statements struct{
	valid []string
	invalid []string
}

var inlineStatements = statements{
	[]string{
		`using database 'business';`,
		`using database 'business' for domain 'sales';`,
		`using database 'business' for domain 'sales' in context 'quoting';`,
		`using database 'business' for domain 'sales' in context 'quoting' within aggregate 'quote';`,
		`for domain 'sales';`,
		`for domain 'sales' in context 'quoting';`,
		`for domain 'sales' in context 'quoting' within aggregate 'quote';`,
		`in context 'quoting';`,
		`in context 'quoting' within aggregate 'quote';`,
		`for domain 'sales' using database 'business' in context 'quoting';`,
		`within aggregate 'quote' in context 'quoting' for domain 'sales' using database 'business';`,
	},
	[]string{
		`using database '';`,
		`using database 'business' for domain '';`,
	},
};

func TestInlineStatements(t *testing.T) {
	assertCorrectParsing(inlineStatements, t);
}

var createNamespaceTypesWithFullyQualfied = statements{
	[]string {
		`create database 'db';`,
		`create domain 'domain' using database 'database';`,
		`create context 'context' using database 'database' for domain 'domain';`,
		`create aggregate 'aggregate' using database 'database' for domain 'domain' in context 'context';`,
	},
	[]string{},
}

func TestCreateNamespaceTypesWithFullyQualfied(t *testing.T) {
	assertCorrectParsing(createNamespaceTypesWithFullyQualfied, t);
}

var createClassesWithAnWithoutFullyQualfied = statements{
	[]string {
		`<| value 'address' using database 'business' for domain 'sales' in context 'quoting'

		|>`,
		`<| value 'address' for domain 'sales' in context 'quoting'
			properties { string value; }
		|>`,
		`<| value 'address' within aggregate 'quote'
			properties { string value; }
		|>`,
		`<| value 'address'
			properties { string value; }
		|>`,
	},
	[]string{},
}


func TestCreateClassesWithAnWithoutFullyQualfied(t *testing.T) {
	assertCorrectParsing(createClassesWithAnWithoutFullyQualfied, t);
}

var blockStatements = statements{
	[]string{
		`using database 'database' for domain 'domain' in context 'context':
		{
			create aggregate 'aggregate';
		}`,
		`using database 'database' for domain 'domain' in context 'context':
		{
		<| value 'address'
			properties { string value; }
		|>
		}`,
	},
	[]string{},
};

func TestBlockStatements(t *testing.T) {
	assertCorrectParsing(blockStatements, t);
}

func assertCorrectParsing(statements statements, t *testing.T) {
	for _, statement := range statements.valid {
		parsed, _ := Parse("", []byte(statement));
		if (parsed == nil) {
			t.Error("Could not parse " + statement);
		}
	}
	for _, statement := range statements.invalid {
		parsed, _ := Parse("", []byte(statement));
		if (parsed != nil) {
			t.Error("Could parse " + statement);
		}
	}
}