package parser

/******************
  Namespaces
******************/

type Namespace struct {
	Paths []string
}

func (n *Namespace) Merge(o *Namespace) Namespace {
	paths := merge(n.Paths, o.Paths);
	return Namespace(paths);
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

func fillNils(paths []string, fill string) []string{
	for i,path := range paths {
		if (path == nil) {
			paths[i] = fill;
		}
	}
	return paths;
}

func preparePaths(paths []string, length int) []string{
	if (!checkLength(paths, length)) {
		panic("Paths is the wrong length");
	}
	return fillNils(paths, "");
}

func (n *Namespace) Check() bool {
	for _, path := range n.Paths {
		if (path == "") {
			panic("Missing namespace parameter");
		}
	}
	return true;
}

func NewDatabaseNamespace(paths []string) Namespace {
	return Namespace{preparePaths(paths, 1)};
}

func NewDomainNamespace(paths []string) Namespace {
	return Namespace{preparePaths(paths, 2)};
}

func NewContextNamespace(paths []string) Namespace {
	return Namespace{preparePaths(paths, 3)};
}

func NewAggregateNamespace(paths []string) Namespace {
	return Namespace{preparePaths(paths, 4)};
}


/******************
  Commands
******************/

type CreateDatabase struct {
	ID string;
	Name string;
}

func (c *CreateDatabase) Check () bool {
	return true;
}


type CreateDomain struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateDomain) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateDomain) Check () bool {
	return c.Namespace.Check();
}


type CreateContext struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateContext) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateContext) Check () bool {
	return c.Namespace.Check();
}


type CreateValue struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateValue) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateValue) Check () bool {
	return c.Namespace.Check();
}


type CreateEntity struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateEntity) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateEntity) Check () bool {
	return c.Namespace.Check();
}


type CreateAggregate struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateAggregate) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateAggregate) Check () bool {
	return c.Namespace.Check();
}


type CreateEvent struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateEvent) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateEvent) Check () bool {
	return c.Namespace.Check();
}


type CreateCommand struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateCommand) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateCommand) Check () bool {
	return c.Namespace.Check();
}


type CreateProjection struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateProjection) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateProjection) Check () bool {
	return c.Namespace.Check();
}


type CreateInvariant struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateInvariant) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateInvariant) Check () bool {
	return c.Namespace.Check();
}


type CreateQuery struct {
	ID string;
	Name string;
	Namespace Namespace;
}

func (c *CreateQuery) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateQuery) Check () bool {
	return c.Namespace.Check();
}


