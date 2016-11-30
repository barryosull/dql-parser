package class

import (
	"testing"
)

var classes = []string{
	`properties
	{
		value\service-charge service_charge;
		value\category category;
	}`,
	`properties
	{
		string value;
	}

	check
	(
		return value != 0;
	)`,
	`properties
	{
		string value;
	}

	function doThing(){}
	`,
	`properties{string value;}
	function doThing(){}
	function doThing2(){}
	`,
	`
	properties
	{
		string value;
	}

	function doThing()
	{
		a = 2;
		return (a * 3);
	}

	function doThing2(value\service-charge service_charge, value\category category)
	{

	}`,
};

func TestValues(t *testing.T) {
	assertCanParseStatements(classes, t);
}

func assertCanParseStatements(statements []string, t *testing.T) {
	for _, statement := range statements {
		parsed, _ := Parse("", []byte(statement));
		if (parsed == nil) {
			t.Error("Could not parse " + statement);
		}
	}
}

var differentDeclarationOrder = []string{
	`properties
	{
		string value;
	}`,
};

func TestDifferentDeclarationOrder(t *testing.T) {
	assertCanParseStatements(differentDeclarationOrder, t);
}