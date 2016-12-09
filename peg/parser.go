package peg

import (
	"parser/peg/namespace"
	"parser"
);

type Peg struct {

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
		command.Check();
	}

	return commands;
}

func (p *Peg) ReferenceError(commandID string, reference string) string {
	return "";
}
