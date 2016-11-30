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

var ifs = []string{
	`if (a) {
		a;
	}`,
	`if (a) {
		b;
		return c;
	} else {
		d;
		return e;
	}`,
	`if(a){return b;}`,
	`if(a){}`,
};

func TestIfs(t *testing.T) {
	assertCanParse(ifs, t);
}

var forEachs = []string {
	`foreach (a as b) {
		a;
	}`,
	`foreach (a as b=>c) {
		a;
	}`,
	`foreach (a->b() as b=>c) {
		a;
	}`,
}

var badForeachs = []string {
	`foreach (a as b,c) {
		a;
	}`,
}

func TestForeachs(t *testing.T) {
	assertCanParse(forEachs, t);
	assertCannotParse(badForeachs, t);
}

var blocks = []string {
	"a = 22;\nb = 5 + 1;\n c= b - a;",
}

func TestBlocks(t *testing.T) {
	assertCanParse(blocks, t);
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