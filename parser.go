package parser

type Parser interface {
	Parse(dql string) ([]Command, error)
}

type Command interface {
	AssertValid() error
}

