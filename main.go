package parser

import (
	//"parser/ast"
	"fmt"
	//"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"os"
	"parser/peg"
)

func main() {
	filepath := os.Args[1]
	astNode, err := peg.ParseFile(filepath)
	if err != nil {
		panic(err)
	}

	fmt.Println("AST: ")

	spew.Dump(astNode)

	fmt.Println("\n")
}
