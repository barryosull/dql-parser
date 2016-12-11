package parser

import (
	"testing"
);

var mergedNamespaces = []struct{
	origin Namespace
	other Namespace
	expected Namespace
}{
	{Namespace{[]string{""}}, Namespace{[]string{"a", "b"}}, Namespace{[]string{"a"}}},
	{Namespace{[]string{"a", "b"}}, Namespace{[]string{"c", "d"}}, Namespace{[]string{"a", "b"}}},
	{Namespace{[]string{"a", "", "c"}}, Namespace{[]string{"a", "b"}}, Namespace{[]string{"a", "b", "c"}}},
};

func TestMergingNamespaces(t *testing.T) {
	for _, row := range mergedNamespaces {
		actual := row.origin.Merge(row.other);
		if (!actual.Equal(row.expected)) {
			t.Error("Got the following result instead of expected");
			t.Error(actual);
		}
	}
}

var validNamespaces =
	[]Namespace{
		{[]string{"a"}},
		{[]string{"a", "b"}},
		{[]string{"a", "b", "c"}},
	};

func TestCheckingNamespaces(t *testing.T) {

	for _, namespace := range validNamespaces {
		if (namespace.AssertValid() != nil) {
			t.Error("Namespace is invalid")
		}
	}
}

var invalidNamespaces =[]struct {
	namespace Namespace
	errorMessage string
}{
	{Namespace{[]string{""}}, "database not selected"},
	{Namespace{[]string{"a", ""}}, "domain not selected"},
	{Namespace{[]string{"a", "b", ""}}, "context not selected"},
	{Namespace{[]string{"a", "b", "c", ""}}, "aggregate not selected"},
}

func TestInvalidNamespaces(t *testing.T){
	for _, row := range invalidNamespaces {
		err := row.namespace.AssertValid();
		if (err.Error() != row.errorMessage) {
			t.Error("Error message does not match.")
			t.Error("Expected: "+row.errorMessage);
			t.Error("Actual: "+err.Error());
		}
	}
}

func TestPathConstraintsAreMetForTypes(t *testing.T) {

	_, err := NewDatabaseNamespace([]string{""});
	testValidNamespace(err , t);
	_, err = NewDomainNamespace([]string{"", ""});
	testValidNamespace(err , t);
	_, err = NewContextNamespace([]string{"", "", ""});
	testValidNamespace(err , t);
	_, err = NewAggregateNamespace([]string{"", "", "", ""});
	testValidNamespace(err , t);

	_, err = NewDatabaseNamespace([]string{"", ""});
	testInvalidNamespace(err , t);
	_, err = NewDomainNamespace([]string{""});
	testInvalidNamespace(err , t);
	_, err = NewContextNamespace([]string{"", ""});
	testInvalidNamespace(err , t);
	_, err = NewAggregateNamespace([]string{"", "", ""});
	testInvalidNamespace(err , t);
}

func testValidNamespace(error error, t *testing.T) {
	if (error != nil) {
		t.Error("Namespace is not valid");
	}
}

func testInvalidNamespace(error error, t *testing.T) {
	if (error == nil) {
		t.Error("Namespace is valid");
	}
}
