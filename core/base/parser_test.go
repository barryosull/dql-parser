package base

import (
	"testing"
)

var classRef = []string{
	"'value\\cup'",
	"'value\\tea-cup'",
	"'value\\tea_cup'",
	"'entity\\cat'",
	"'command\\kill'",
	"'event\\killed'",
	"	'value\\cup'",
	"   'value\\cup'",
};

func TestClassReferences(t *testing.T) {
	assertCanParse(classRef, t);
}

func assertCanParse(statements []string, t *testing.T) {
	for _, statement := range statements {
		parsed, _ := Parse("", []byte(statement));
		if (parsed == nil) {
			t.Error("Could not parse "+statement);
		}
	}
}