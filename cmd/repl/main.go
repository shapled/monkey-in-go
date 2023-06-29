package main

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
	"os"
	"os/user"
)

// repl/repl.go

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	Start(os.Stdin, os.Stdout)
}
