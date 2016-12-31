package token

import "testing"

func TestToString(t *testing.T) {
	tok := NewToken(STRING, "value", 10)

	expected := "Token(string, \"value\", 10)"

	if (tok.String() != expected) {
		t.Error("Token string does not match expected")
		t.Error("Expected: "+expected);
		t.Error("Got: "+tok.String());
	}
}