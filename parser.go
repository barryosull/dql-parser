package parser

type Parser interface {
	Parse(dql string) ([]Command, string)
	ReferenceError(commandID string, reference string) string
}

type Command interface {
	Check() bool
}
