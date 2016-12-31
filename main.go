package main

import (
	"parser/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}

