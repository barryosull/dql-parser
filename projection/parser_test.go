package projection

import (
	"testing"
)

var projections = []string{
	`<| domain projection 'quote' |>`,
	`<| aggregate projection 'quote'
	 	properties {}
	|>`,
	`<| domain projection 'quote'
	 	properties {}
	 	when event 'started'
		{

		}
	|>`,
	`<| aggregate projection 'quote'

		properties
		{
			value\uuid agency_id;
			value\uuid brand_id;
			value\boolean is_started = false;
			value\boolean is_completed = false;
			value\integer quote_number;

			value\price total;
			value\service-charge additional_charge;
			value\deposit deposit;
			value\balance-due balance_due;
			value\service-charge service_charge;

			index items = [];
			index passengers = [];
			value\uuid quote_owner;
		}

		when event 'started'
		{

		}

		when event 'item-added'
		{

		}

		when event 'item-edited'
		{

		}

		when event 'item-removed'
		{

		}

		when event 'passenger-added'
		{

		}

		when event 'passenger-removed'
		{

		}

		when event 'owner-assigned'
		{

		}

		when event 'owner-removed'
		{

		}

		when event 'fees-set'
		{
		}

		when event 'completed'
		{

		}
	|>`,
};

func TestProjections(t *testing.T) {
	assertCanParse(projections, t);
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
