package interpreter

import (
	"testing"

	"github.com/AlyxPink/meowlang/lexer"
	"github.com/AlyxPink/meowlang/object"
	"github.com/AlyxPink/meowlang/parser"
)

func TestInterpreter(t *testing.T) {
	input := `lick x = 5;
              lick y = 10;
              claw x + y;`

	l := lexer.NewLexer(input)
	p := parser.NewParser(l.Tokenize())
	program := p.ParseProgram()

	interpreter := NewInterpreter()
	result := interpreter.Interpret(program)

	if result == nil {
		t.Fatalf("Interpret() returned nil")
	}

	if integer, ok := result.(*object.Integer); ok {
		if integer.Value != 15 {
			t.Errorf("result.Value not 15. got=%d", integer.Value)
		}
	} else {
		t.Fatalf("result is not *object.Integer. got=%T (%+v)", result, result)
	}
}
