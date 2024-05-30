package parser

import (
	"testing"

	"github.com/AlyxPink/meowlang/ast"
	"github.com/AlyxPink/meowlang/lexer"
)

func TestParsingAssignStatements(t *testing.T) {
	input := `lick x = 5;
              lick y = 15;
              lick foobar = 838383;`

	l := lexer.NewLexer(input)
	p := NewParser(l.Tokenize())

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testAssignStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testAssignStatement(t *testing.T, s ast.Statement, name string) bool {
	if s == nil {
		t.Errorf("Assign statement is nil")
		return false
	}

	if s.TokenLiteral() != "lick" {
		t.Errorf("s.TokenLiteral not 'lick'. got=%q", s.TokenLiteral())
		return false
	}

	assignStmt, ok := s.(*ast.AssignStatement)
	if !ok {
		t.Errorf("s not *ast.AssignStatement. got=%T", s)
		return false
	}

	if assignStmt.Name.Value != name {
		t.Errorf("assignStmt.Name.Value not '%s'. got=%s", name, assignStmt.Name.Value)
		return false
	}

	if assignStmt.Name.TokenLiteral() != name {
		t.Errorf("assignStmt.Name.TokenLiteral() not '%s'. got=%s", name, assignStmt.Name.TokenLiteral())
		return false
	}

	return true
}
