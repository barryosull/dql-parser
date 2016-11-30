package components

import (
	"testing"
)

var components = []string{
	`
	properties
	{
		value\service-charge service_charge;
		value\category category;
	}`,
	`
	check
	(
		return value != 0;
	)`,
	`
	function doThing()
	{
		a = 2;
		return (a * 3);
	}`,
	`
	function doThing2(value\service-charge service_charge, value\category category)
	{

	}`,
	`
	function doThing(){

	}`,
};

var invalidComponents = []string {
	`prumperties {}`,
}

func TestClassComponents(t *testing.T) {
	assertCanParseStatements(components, t);
	assertCannotParseStatements(invalidComponents, t);
}

var listOfcomponents = []string{
	`
	properties
	{
		string value;
	}

	check
	(
		return value != 0;
	)`,
	`
	properties
	{
		string value;
	}

	function doThing(){}
	`,
	`
	properties{string value;}
		function doThing(){}
		function doThing2(){}
	`,
	`
	properties
	{
		string value;
	}

	check
	(
		return value != 0;
	)

	function doThing()
	{
		a = 2;
		return (a * 3);
	}

	function doThing2(value\service-charge service_charge, value\category category)
	{

	}`,
};

func TestListOfComponents(t *testing.T) {
	assertCanParseStatements(listOfcomponents, t);
}

func assertCanParseStatements(statements []string, t *testing.T) {
	for _, statement := range statements {
		parsed, _ := Parse("", []byte(statement));
		if (parsed == nil) {
			t.Error("Could not parse " + statement);
		}
	}
}

func assertCannotParseStatements(statements []string, t *testing.T) {
	for _, statement := range statements {
		parsed, _ := Parse("", []byte(statement));
		if (parsed != nil) {
			t.Error("Could parse " + statement);
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