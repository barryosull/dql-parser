package invariant

import (
	"testing"
)

var queries = []string{
	`<| query 'next-revision-number' on 'projections\quote-numbers'
		handler
		{

		}
	|>`,
	`<| query 'next-revision-number' on 'projections\quote-numbers'

		properties
		{
			value\uuid agency_id;
			value\integer quote_number;
		}

		handler
		{

		}
	|>`,
};

func TestQueries(t *testing.T) {
	assertCanParse(queries, t);
}

func assertCanParse(statements []string, t *testing.T) {
	for _, statement := range statements {
		Debug(true);
		parsed, _ := Parse("", []byte(statement));
		if (parsed == nil) {
			t.Error("Could not parse "+statement);
		}

	}
}
