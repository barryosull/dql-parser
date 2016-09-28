package main

import (
	"fmt"
	"os"
	"parser/peg"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	filepath := os.Args[1]
	ast, err := peg.ParseFile(filepath)
	if err != nil {
		panic(err);
	}

	fmt.Println("AST: ");
	spew.Dump(ast);
	fmt.Println("\n");
}

