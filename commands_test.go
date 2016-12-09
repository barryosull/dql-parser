package parser

import "testing";

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
