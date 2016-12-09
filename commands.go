package parser

/******************
  Namespaces
******************/

type Namespace  {
	Check() bool
	Merge(o Namespace) Namespace
	ToArray() []string
}

func merge(origin []string, other []string) []string {
	for i := 0; i < len(origin); i++ {
		if (origin[i] == "") {
			origin[i] = other[i];
		}
	}
	return origin;
}

type DatabaseNamespace struct {
	Database string;
}

func NewDatabaseNamespace(paths []string) *DatabaseNamespace {
	return &DatabaseNamespace{paths[0]};
}

func (n *DatabaseNamespace) Check () bool {
	return n.Database != "";
}

func (n *DatabaseNamespace) Merge(o Namespace) Namespace {
	paths := merge(n.ToArray(), o.ToArray());
	return NewDatabaseNamespace(paths);
}

func (n *DatabaseNamespace) ToArray () []string {
	return []string{n.Database};
}


type DomainNamespace struct {
	Database string;
	Domain string;
}

func NewDomainNamespace(paths []string) DomainNamespace {
	return DomainNamespace{paths[0], paths[1]};
}

func (n *DomainNamespace) Check () bool {
	return n.Database != "" && n.Domain != "";
}

func (n *DomainNamespace) Merge(o Namespace) Namespace {
	paths := merge(n.ToArray(), o.ToArray());
	return NewDomainNamespace(paths);
}

func (n *DomainNamespace) ToArray () []string {
	return []string{n.Database, n.Domain};
}


type ContextNamespace struct {
	Database string;
	Domain string;
	Context string;
}

func NewContextNamespace(paths []string) *ContextNamespace {
	return &ContextNamespace{paths[0], paths[1], paths[2]};
}

func (n *ContextNamespace) Check () bool {
	return n.Database != "" && n.Domain != "" && n.Context != "";
}

func (n *ContextNamespace) Merge(o Namespace) Namespace {
	paths := merge(n.ToArray(), o.ToArray());
	return NewContextNamespace(paths);
}

func (n *ContextNamespace) ToArray () []string {
	return []string{n.Database, n.Domain, n.Context};
}


type AggregateNamespace struct {
	Database string;
	Domain string;
	Context string;
	Aggregate string;
}

func NewAggregateNamespace(paths []string) *AggregateNamespace {
	return &AggregateNamespace{paths[0], paths[1], paths[2], paths[3]};
}

func (n *AggregateNamespace) Check () bool {
	return n.Database != "" && n.Domain != "" && n.Context != "" && n.Aggregate != "";
}

func (n *AggregateNamespace) Merge(o Namespace) Namespace {
	paths := merge(n.ToArray(), o.ToArray());
	return NewAggregateNamespace(paths);
}

func (n *AggregateNamespace) ToArray () []string {
	return []string{n.Database, n.Domain, n.Context, n.Aggregate};
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
	Namespace DatabaseNamespace;
}

func (c *CreateDomain) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o).(DomainNamespace);
}

func (c *CreateDomain) Check () bool {
	return c.Namespace.Check();
}


type CreateContext struct {
	ID string;
	Name string;
	Namespace DomainNamespace;
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
	Namespace ContextNamespace;
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
	Namespace ContextNamespace;
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
	Namespace ContextNamespace;
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
	Namespace AggregateNamespace;
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
	Namespace AggregateNamespace;
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
	Namespace AggregateNamespace;
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
	Namespace AggregateNamespace;
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
	Namespace AggregateNamespace;
}

func (c *CreateQuery) MergeNamespace (o Namespace) {
	c.Namespace = c.Namespace.Merge(o);
}

func (c *CreateQuery) Check () bool {
	return c.Namespace.Check();
}


