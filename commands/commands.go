package commands

/*
import "errors"

type Command interface {
	AssertValid() error
}


///////////////////
  Namespaces
///////////////////

type Namespace struct {
	Paths []string
}

func (n *Namespace) Merge(o Namespace) Namespace {
	paths := merge(n.Paths, o.Paths);
	return Namespace{paths};
}

func (n *Namespace) Equal (o Namespace) bool {
	if (len(n.Paths) != len(o.Paths)) {
		return false;
	}
	for i, path := range n.Paths {
		if (path != o.Paths[i]) {
			return false;
		}
	}
	return true;
}

func merge(origin []string, other []string) []string {
	for i := 0; i < len(origin); i++ {
		if (origin[i] == "") {
			origin[i] = other[i];
		}
	}
	return origin;
}

func checkLength(paths []string, length int) bool {
	return len(paths) == length;
}


func NewNamespace(paths []string, length int) (Namespace, error){
	if (!checkLength(paths, length)) {
		return Namespace{[]string{}}, errors.New("Paths is the wrong length")
	}
	return Namespace{paths}, nil;
}

var pathNames = []string {
	"database",
	"domain",
	"context",
	"aggregate",
};

func (n *Namespace) AssertValid() error {

	for i, path := range n.Paths {
		if (path == "") {
			return errors.New(pathNames[i]+" not selected");
		}
	}
	return nil;
}

func NewDatabaseNamespace(paths []string) (Namespace, error) {
	return NewNamespace(paths, 1);
}

func NewDomainNamespace(paths []string) (Namespace, error) {
	return NewNamespace(paths, 2);
}

func NewContextNamespace(paths []string) (Namespace, error) {
	return NewNamespace(paths, 3);
}

func NewAggregateNamespace(paths []string) (Namespace, error) {
	return NewNamespace(paths, 4);
}


///////////////////
  Commands
///////////////////

type CreateDatabase struct {
	ID string;
	Name string;
}

func NewCreateDatabase(ID string, Name string) *CreateDatabase {
	return &CreateDatabase{ID, Name};
}

func (c *CreateDatabase) AssertValid() error {
	return nil;
}


type CreateDomain struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateDomain) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateDomain) AssertValid() error {
	return c.Namespace.AssertValid();
}


type CreateContext struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateContext) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateContext) AssertValid() error {
	return c.Namespace.AssertValid();
}


type CreateValue struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateValue) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateValue) AssertValid() error {
	return c.Namespace.AssertValid();
}


type CreateEntity struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateEntity) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateEntity) AssertValid() error {
	return c.Namespace.AssertValid();
}


type CreateAggregate struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateAggregate) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateAggregate) AssertValid() error {
	return c.Namespace.AssertValid();
}


type CreateEvent struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateEvent) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateEvent) AssertValid() error {
	return c.Namespace.AssertValid();
}


type CreateCommand struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateCommand) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateCommand) AssertValid() error {
	return c.Namespace.AssertValid();
}


type CreateProjection struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateProjection) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateProjection) AssertValid() error {
	return c.Namespace.AssertValid();
}


type CreateInvariant struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateInvariant) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateInvariant) AssertValid() error {
	return c.Namespace.AssertValid();
}


type CreateQuery struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateQuery) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateQuery) AssertValid() error {
	return c.Namespace.AssertValid();
}

*/
