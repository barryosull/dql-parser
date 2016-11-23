package invariant

import (
	"testing"
)

var invariants = []string{
	`<| invariant 'is-editable' on 'projections\quote' |>`,
	`<| invariant 'is-editable' on 'projections\quote'

		check
		(

		)
	|>`,
	`<| invariant 'item-exists' on 'projections\quote'

		properties
		{
			entity\item item;
		}

		check
		(

		)
	|>`,
};

func TestInvariants(t *testing.T) {
	assertCanParse(invariants, t);
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
