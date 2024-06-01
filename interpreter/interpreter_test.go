package interpreter

import (
	"bytes"
	"testing"

	"github.com/AlyxPink/meowlang/lexer"
	"github.com/AlyxPink/meowlang/parser"
)

func interpret(input string) string {
	l := lexer.NewLexer(input)
	p := parser.NewParser(l.Tokenize())
	program := p.ParseProgram()

	var out bytes.Buffer
	interpreter := NewInterpreterWithOutput(&out)
	interpreter.Interpret(program)

	return out.String()
}

func TestInterpreter_IntegerArithmetic(t *testing.T) {
	input := `purr 1 + 2`
	expectedOutput := "3\n"
	output := interpret(input)

	if output != expectedOutput {
		t.Errorf("expected output %q, got %q", expectedOutput, output)
	}
}

func TestInterpreter_StringConcatenation(t *testing.T) {
	input := `purr "Hello" + " world"`
	expectedOutput := "Hello world\n"
	output := interpret(input)

	if output != expectedOutput {
		t.Errorf("expected output %q, got %q", expectedOutput, output)
	}
}

func TestInterpreter_VariableAssignment(t *testing.T) {
	input := `lick x = 42; purr x;`
	expectedOutput := "42\n"
	output := interpret(input)

	if output != expectedOutput {
		t.Errorf("expected output %q, got %q", expectedOutput, output)
	}
}

func TestInterpreter_PrintStatement(t *testing.T) {
	input := `purr 123`
	expectedOutput := "123\n"
	output := interpret(input)

	if output != expectedOutput {
		t.Errorf("expected output %q, got %q", expectedOutput, output)
	}
}

func TestInterpreter_FunctionDefinition(t *testing.T) {
	input := `
    meow double(a) {
        claw a * 2;
    }
    purr double(5);`
	expectedOutput := "10\n"
	output := interpret(input)

	if output != expectedOutput {
		t.Errorf("expected output %q, got %q", expectedOutput, output)
	}
}

func TestInterpreter_FunctionCall(t *testing.T) {
	input := `
    meow addTen(a) {
        claw a + 10;
    }
    purr addTen(10);`
	expectedOutput := "20\n"
	output := interpret(input)

	if output != expectedOutput {
		t.Errorf("expected output %q, got %q", expectedOutput, output)
	}
}
