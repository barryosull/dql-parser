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
		//fmt.Println(parsed);
		if (parsed == nil) {
			t.Error("Could not parse "+statement);
		}

	}
}

var identifiers = []string{
	"a",
	"b",
	"sdadsasd",
	"dsfsd12213",
}

func TestIdentifiers(t *testing.T) {
	assertCanParse(identifiers, t);
}

var parenthesisi = []string{
	"(5)",
}

func TestParenthesis(t *testing.T) {
	assertCanParse(parenthesisi, t);
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