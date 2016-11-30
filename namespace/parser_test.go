package namespace

import (
	"testing"
)

type statements struct{
	valid []string
	invalid []string
}

func (s *statements) assertParse(t *testing.T) {
	for _, statement := range s.valid {
		parsed, _ := Parse("", []byte(statement));
		if (parsed == nil) {
			t.Error("Could not parse " + statement);
			//t.Error(err);
		}
	}
	for _, statement := range s.invalid {
		parsed, _ := Parse("", []byte(statement));
		if (parsed != nil) {
			t.Error("Could parse " + statement);
		}
	}
}

var inlineStatements = statements{
	[]string{
		`using database 'business';`,
		`using database 'business' for domain 'sales';`,
		`using database 'business' for domain 'sales' in context 'quoting';`,
		`using database 'business' for domain 'sales' in context 'quoting' within aggregate 'quote';`,
		`for domain 'sales';`,
		`for domain 'sales' in context 'quoting';`,
		`for domain 'sales' in context 'quoting' within aggregate 'quote';`,
		`in context 'quoting';`,
		`in context 'quoting' within aggregate 'quote';`,
		`for domain 'sales' using database 'business' in context 'quoting';`,
		`within aggregate 'quote' in context 'quoting' for domain 'sales' using database 'business';`,
	},
	[]string{
		`using database '';`,
		`using database 'business' for domain '';`,
	},
};

func TestInlineStatements(t *testing.T) {
	inlineStatements.assertParse(t);
}

var createNamespaceTypesWithFullyQualfied = statements{
	[]string {
		`create database 'db';`,
		`create domain 'domain' using database 'database';`,
		`create context 'context' using database 'database' for domain 'domain';`,
		`create aggregate 'aggregate' using database 'database' for domain 'domain' in context 'context';`,
	},
	[]string{},
}

func TestCreateNamespaceTypesWithFullyQualfied(t *testing.T) {
	createNamespaceTypesWithFullyQualfied.assertParse(t);
}

var createClassesWithAnWithoutFullyQualfied = statements{
	[]string {
		`<| value 'address' using database 'business' for domain 'sales' in context 'quoting'

		|>`,
		`<| value 'address' for domain 'sales' in context 'quoting'
			properties { string value; }
		|>`,
		`<| value 'address' within aggregate 'quote'
			properties { string value; }
		|>`,
	},
	[]string{},
}


func TestCreateClassesWithAnWithoutFullyQualfied(t *testing.T) {
	createClassesWithAnWithoutFullyQualfied.assertParse(t);
}

var blockStatements = statements{
	[]string{
		`using database 'database' for domain 'domain' in context 'context':
		{
			create aggregate 'aggregate';
		}`,
		`using database 'database' for domain 'domain' in context 'context':
		{
		<| value 'address'
			properties { string value; }
		|>
		}`,
	},
	[]string{},
};

func TestBlockStatements(t *testing.T) {
	blockStatements.assertParse(t);
}

var createValues = statements{
	[]string{
		`<| value 'address'
			properties { string value; }
		|>`,
		`<| value 'address'
			properties { string value; }

			check (
				return a + b;
			)
		|>`,

		`<| value 'address'
			properties { string value; }
			check ( return a + b; )

			function doThing(value\a a, value\b b) {
				return 22;
			}
		|>`,
		`<| value 'address'
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

			}
		|>`,
	},
	[]string{
		`<| value 'address'
			properties { string value; }

			handle (
				return 22;
			)
		|>`,
	},
};

func TestCreateValues(t *testing.T) {
	createValues.assertParse(t);
}

var createEntities = statements{
	[]string{
		`<| entity 'address'
			properties { string value; }
		|>`,
		`<| entity 'address'
			properties { string value; }

			check (
				return a + b;
			)
		|>`,

		`<| entity 'address'
			properties { string value; }
			check ( return a + b; )

			function doThing(value\a a, value\b b) {
				return 22;
			}
		|>`,
		`<| entity 'address'
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

			}
		|>`,
	},
	[]string{
		`<| entity 'address'
			properties { string value; }

			handle (
				return 22;
			)
		|>`,
	},
};

func TestCreateEntities(t *testing.T) {
	createEntities.assertParse(t);
}

var createEvents = statements{
	[]string{
		`<| event 'started'
			properties
			{
				value\uuid agent_id;
				value\uuid agency_id;
				value\uuid brand_id;
				value\integer quote_number;
				value\integer revision;
			}
		|>`,
		`
		<| event 'revision-started'
			properties
			{
				value\uuid old_quote_id;
			}

			check (
				return a + b;
			)
		|>`,
	},
	[]string{
		`
		<| event 'revision-started'
			properties
			{
				value\uuid old_quote_id;
			}

			handle ( )
		|>`,
	},
};

func TestEvents(t *testing.T) {
	createEvents.assertParse(t);
}

var createCommands = statements{
	[]string{
		`<| command 'start'

			properties
			{
				string value;
			}

			handler
			{
				assert invariant not 'is-started';
				quote_number = run query 'next-quote-number' (agency_id);
				apply event 'started' (agent_id, agency_id, brand_id);
			}
		|>`,
		`
		<| command 'start-from-existing'

			properties
			{
				string value;
			}

			handler
			{
				assert invariant not 'is-started';

				revision = run query 'next-revision-number' (agency_id, quote_number);

				apply event 'started' (agent_id, agency_id, brand_id, quote_number, revision);
			}
		|>`,
	},
	[]string{
		`
		<| command 'start' properties
			{ |>`,
	},
};

func TestCommands(t *testing.T) {
	createCommands.assertParse(t);
}

var createProjections = statements{
	[]string{
		`<| aggregate projection 'quote
			properties
			{
				value\uuid agency_id;
				value\uuid brand_id;
			}

			when event 'started'
			{
				agency_id = event.agency_id;
				brand_id = event.brand_id;
				is_started = true;
			}

			when event 'item-added'
			{
				items.add(event.item);
			}
		}`,
	},
	[]string{},
};

func TestProjections(t *testing.T) {
	createProjections.assertParse(t);
}



