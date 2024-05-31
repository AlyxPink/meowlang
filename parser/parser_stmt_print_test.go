package parser

import (
	"testing"

	"github.com/AlyxPink/meowlang/ast"
	"github.com/AlyxPink/meowlang/lexer"
)

func TestPrintStatements(t *testing.T) {
	input := `
purr "Hello world!"
purr 5
`

	l := lexer.NewLexer(input)
	p := NewParser(l.Tokenize())

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 2 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedValue string
	}{
		{"Hello world!"},
		{"5"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testPrintStatement(t, stmt, tt.expectedValue) {
			return
		}
	}
}

func testPrintStatement(t *testing.T, s ast.Statement, value string) bool {
	if s == nil {
		t.Errorf("Assign statement is nil")
		return false
	}

	if s.TokenLiteral() != "purr" {
		t.Errorf("s.TokenLiteral not 'purr'. got=%q", s.TokenLiteral())
		return false
	}

	printStmt, ok := s.(*ast.PrintStatement)
	if !ok {
		t.Errorf("s not *ast.PrintStatement. got=%T", s)
		return false
	}

	if printStmt.Value.TokenLiteral() != value {
		t.Errorf("printStmt.Value.TokenLiteral() not '%s'. got=%s", value, printStmt.Value.TokenLiteral())
		return false
	}

	return true
}
