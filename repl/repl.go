package repl

import (
	"bufio"
	"fmt"
	"io"
	"parser/tokenizer"
	"parser/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {

	io.WriteString(out, LOGO)

	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if (line == "exit") {
			io.WriteString(out, "See ya!\n")
			break;
		}

		t := tokenizer.NewTokenizer(line)

		tokens, err := t.Tokens()

		if (err != nil) {
			printTokenizerErrors(out, err)
		} else {
			for _, tok := range tokens {
				io.WriteString(out, tok.String()+"\n")
			}
		}
		io.WriteString(out, "\n")
	}
}

const LOGO = `    ____  ____    __
   / __ \/ __ \  / /
  / / / / / / / / /
 / /_/ / /_/ / / /___
/_____/\___\_\/_____/

`

func printTokenizerErrors(out io.Writer, err *token.Error) {
	io.WriteString(out, "Tokenization Error!\n")
	io.WriteString(out, err.String()+"\n")
}
