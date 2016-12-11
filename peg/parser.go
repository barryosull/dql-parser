package peg

import (
	"parser/peg/namespace"
	"parser"
);

type Peg struct {

}

type IdGenerator interface {
	Generate() string
}

func (p *Peg) Parse (dql string) ([]parser.Command, string) {

	commandInterfaces, err := namespace.Parse("", []byte(dql));
	if (commandInterfaces == nil) {
		panic(err);
	}

	commands, ok := commandInterfaces.([]parser.Command)
	if (!ok) {
		panic("Parser did not return commands");
	}

	for _, command := range commands {
		err := command.AssertValid();
		if (err != nil) {

		}
	}

	return commands;
}

func (p *Peg) ReferenceError(commandID string, reference string) string {
	return "";
}

func NewParser() parser.Parser {
	return Peg{};
}
