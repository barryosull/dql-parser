package expression

import "testing"

var literals = []string{
	"true",
	"false",
	"null",
	//"\"string\"",
	"1",
	"2.456",
};

func TestLiterals(t *testing.T) {
	for _, literal := range literals {
		parsed, _ := Parse("", []byte(literal));
		if (parsed == nil) {
			t.Error("Could not parse "+literal);
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
	for _, identifier := range identifiers {
		parsed, _ := Parse("", []byte(identifier));
		if (parsed == nil) {
			t.Error("Could not parse "+identifier);
		}
	}
}
