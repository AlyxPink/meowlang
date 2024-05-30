package parser

import (
	"fmt"
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

func TestParsingFunctionStatements(t *testing.T) {
	input := `meow add(a, b) { claw a + b; }`

	l := lexer.NewLexer(input)
	p := NewParser(l.Tokenize())

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. got=%d", len(program.Statements))
	}

	stmt := program.Statements[0]
	functionStmt, ok := stmt.(*ast.FunctionStatement)
	if !ok {
		t.Fatalf("stmt not *ast.FunctionStatement. got=%T", stmt)
	}

	if functionStmt.Name.Value != "add" {
		t.Fatalf("functionStmt.Name.Value not 'add'. got=%s", functionStmt.Name.Value)
	}

	if len(functionStmt.Parameters) != 2 {
		t.Fatalf("functionStmt.Parameters does not contain 2 parameters. got=%d", len(functionStmt.Parameters))
	}

	testLiteralExpression(t, functionStmt.Parameters[0], "a")
	testLiteralExpression(t, functionStmt.Parameters[1], "b")

	if len(functionStmt.Body.Statements) != 1 {
		t.Fatalf("functionStmt.Body.Statements does not contain 1 statement. got=%d", len(functionStmt.Body.Statements))
	}

	bodyStmt, ok := functionStmt.Body.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Fatalf("functionStmt.Body.Statements[0] not *ast.ReturnStatement. got=%T", functionStmt.Body.Statements[0])
	}

	testInfixExpression(t, bodyStmt.ReturnValue, "a", "+", "b")
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case string:
		return testIdentifier(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value, integ.TokenLiteral())
		return false
	}

	return true
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}

	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}

	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value, ident.TokenLiteral())
		return false
	}

	return true
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.InfixExpression. got=%T(%s)", exp, exp)
		return false
	}

	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}

	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s'. got=%q", operator, opExp.Operator)
		return false
	}

	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}

	return true
}
