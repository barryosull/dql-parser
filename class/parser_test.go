package class

import (
	"testing"
	"strings"
)

var classes = []string{
	`<| {KIND} 'value'
		properties
		{
			value\service-charge service_charge;
		    	value\category category;
		}
	|>`,
	`<| {KIND} 'balanceDue'

		properties
		{
			string value;
		}

		check
		(

		)
	|>`,
	`<| {KIND} 'value'
		properties
		{
			string value;
		}

		function doThing(){}
	|>`,
	`<| {KIND} 'value'
		properties{string value;}
		function doThing(){}
		function doThing2(){}
	|>`,
	`<| {KIND} 'value'
		properties
		{
			string value;
		}

		function doThing()
		{

		}

		function doThing2(value\service-charge service_charge, value\category category)
		{

		}
	|>`,
};

func TestValues(t *testing.T) {
	assertCanParseStatements(classes, t);
}

func assertCanParseStatements(statements []string, t *testing.T) {
	for _, statement := range statements {
		assertCanParseStatement(statement, t);
	}
}

func assertCanParseStatement(statement string, t *testing.T) {
	var kinds = []string{
		"value",
		"entity",
		"command",
		"event",
	};
	for _, kind := range kinds {
		classStatement := strings.Replace(statement, "{KIND}", kind, -1);
		parsed, _ := Parse("", []byte(classStatement));
		if (parsed == nil) {
			t.Error("Could not parse " + classStatement);
		}
	}
}