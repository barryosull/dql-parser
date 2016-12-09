package parser

import "testing";

var mergedNamespaces = []struct{
	origin Namespace
	other Namespace
	result Namespace
}{
	{Namspace{[]string{""}}, NewDomainNamespace([]string{"a", "b"}), NewDatabaseNamespace([]string{"a"})},
};

func TestMergingNamespaces(t *testing.T) {
	for i, row := range mergedNamespaces {
		merged := row.origin.Merge(row.other);
		if (merged != row.result) {
			t.Error("mergedNamespaces ["+i+"] failed to run");
		}
	}
}
