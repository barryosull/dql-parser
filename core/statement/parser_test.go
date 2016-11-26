package statement

import (
	"testing"
)

var returns = []string{
	"return true;",
	"return a + b;",
}

func TestReturns(t *testing.T) {
	assertCanParse(returns, t);
}

var blocks = []string {
	"a = 22;\nb = 5 + 1;\n c= b - a;",
}

func TestBlocks(t *testing.T) {
	assertCanParse(blocks, t);
}

func assertCanParse(statements []string, t *testing.T) {
	for _, statement := range statements {
		Debug(true);
		parsed, _ := Parse("", []byte(statement));
		if (parsed == nil) {
			t.Error("Could not parse "+statement);
		}

	}
}