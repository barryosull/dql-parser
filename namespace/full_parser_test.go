package namespace

import (
	"testing"
	"io/ioutil"
)

func assertParse(dql string, t *testing.T) {
	parsed, _ := Parse("", []byte(dql));
	if (parsed == nil) {
		t.Error("Could not parse '" + dql + "'");
	}
}

func TestFullDQLContext(t *testing.T) {
	dql, _ :=ioutil.ReadFile("../dql_examples/dynamicres.dql");
	assertParse(string(dql), t);
}



