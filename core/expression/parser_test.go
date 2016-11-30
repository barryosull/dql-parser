package expression

import (
	"testing"
	//"fmt"
)

var literals = []string{
	"true",
	"false",
	"null",
	"\"string\"",
	"1",
	"2.4",
};

func TestLiterals(t *testing.T) {
	assertCanParse(literals, t);
}

func assertCanParse(statements []string, t *testing.T) {
	for _, statement := range statements {
		parsed, _ := Parse("", []byte(statement));
		if (parsed == nil) {
			t.Error("Could not parse "+statement);
		}
	}
}

func assertCannotParse(statements []string, t *testing.T) {
	for _, statement := range statements {
		parsed, _ := Parse("", []byte(statement));
		if (parsed != nil) {
			t.Error("Could parse invalid statement "+statement);
		}
	}
}

var identifiers = []string{
	"a",
	"b",
	"sdadsasd",
	"dsfsd12213",
}

var badIdentifiers = []string{
	"",
	"adds dffsdf",
}

func TestIdentifiers(t *testing.T) {
	assertCanParse(identifiers, t);
	assertCannotParse(badIdentifiers, t);
}


var parenthesisi = []string{
	"(5)",
}

func TestParenthesis(t *testing.T) {
	assertCanParse(parenthesisi, t);
}

var newInstances = []string {
	"'value\\integer'(1)",
	"'value\\integer'",
}

func TestNewInstances(t *testing.T) {
	assertCanParse(newInstances, t);
}

var unary = []string {
	"+a",
	"-a",
	//"a++",
	//"a--",
	"!a",
}

func TestUnary(t *testing.T) {
	assertCanParse(unary, t);
}

var arithmetic = []string {
	"a + b",
	"a - b",
	"a * b",
	"a / b",
	"a % b",
}

func TestArithmetic(t *testing.T) {
	assertCanParse(arithmetic, t);
}

var assignment = []string {
	"a = 1",
	"a->b = c",
	"a = 'value\\integer'(1)",
}

func TestAssignment(t *testing.T) {
	assertCanParse(assignment, t);
}


var logical = []string {
	"a and b",
	"a or b",
}

func TestLogical(t *testing.T) {
	assertCanParse(logical, t);
}

var comparison = []string {
	"a == b",
	"a != b",
	"a < b",
	"a > b",
	"a <= b",
	"a >= b",
	"a === b",
	"a !== b",
}

func TestComparison(t *testing.T) {
	assertCanParse(comparison, t);
}

var objectAccess = []string{
	"a->b",
	"a->b->c",
}

func TestObjectAccess(t *testing.T) {
	assertCanParse(objectAccess, t);
}

var methodCalls = []string{
	"a()",
	"a->b()",
	"a->b->z()",
	"a->b(1)",
	"a->b(1,2,3,4)",
	"a->b(1)->c()",
}

func TestMethodCalls(t *testing.T) {
	assertCanParse(methodCalls, t);
}

var queries = []string {
	`run query 'next-revision-number'`,
	`run query 'next-revision-number'() `,
	`run query 'next-revision-number'(agency_id, quote_number)`,
};

var badQueries = []string{
	`ran querty 'next-revision-number'`,
};

func TestQueries(t *testing.T) {
	assertCanParse(queries, t);
	assertCannotParse(badQueries, t);
}

var compound = []string{
	"a + (a - b)",
	"(a + b) + (a - b)",
	"(a + b) + (a - b) - a->b->c + a->b() - !b + a and b",
	"a->b = 'value\\integer'(1) - ('value\\integer'(1) + b) + (a - b) - a->b->c + a->b() - !b + a and b",
	"a = b + c = c + 24",
}

func TestCompound(t *testing.T) {
	assertCanParse(compound, t);
}


