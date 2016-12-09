package namespace

type ReferenceChecker interface {
	Add(reference Reference);
	AssertExists(reference Reference);
}

type Reference  struct {
	database string
	domain string
	context string
	kind string
	className string
}

type Inmemory struct {
	references []string
}

func (i *Inmemory) Add (reference Reference) string {
	if (contains(i.references, reference)) {
		return "Class reference '"+reference+"' already defined, cannot redefine";
	}
}

func contains(s []string, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (i *Inmemory) AssertExists (reference Reference) {

}

func NewInMemoryReferenceChecker() *Inmemory {
	return &Inmemory{};
}