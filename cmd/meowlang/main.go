package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/AlyxPink/meowlang/interpreter"
	"github.com/AlyxPink/meowlang/lexer"
	"github.com/AlyxPink/meowlang/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: meowlang <filename>")
		return
	}

	filename := os.Args[1]
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	l := lexer.NewLexer(string(content))
	tokens := l.Tokenize()

	p := parser.NewParser(tokens)
	ast := p.ParseProgram()

	var buf bytes.Buffer
	i := interpreter.NewInterpreterWithOutput(&buf)
	i.Interpret(ast)

	fmt.Print(buf.String())
}
