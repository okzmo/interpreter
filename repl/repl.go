package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/okzmo/interpreter/lexer"
	"github.com/okzmo/interpreter/token"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	const PROMPT = ">> "

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
