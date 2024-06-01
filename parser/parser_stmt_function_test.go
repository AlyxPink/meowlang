package parser

import (
	"testing"

	"github.com/AlyxPink/meowlang/ast"
	"github.com/AlyxPink/meowlang/lexer"
)

func TestParsingFunctionStatements(t *testing.T) {
	input := `
	meow add(a, b) {
		purr 10
		claw a + b
	}`

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

	if len(functionStmt.Body.Statements) != 2 {
		t.Fatalf("functionStmt.Body.Statements does not contain 2 statement. got=%d", len(functionStmt.Body.Statements))
	}

	bodyPrintStmt, ok := functionStmt.Body.Statements[0].(*ast.PrintStatement)
	if !ok {
		t.Fatalf("functionStmt.Body.Statements[0] not *ast.PrintStatement. got=%T", functionStmt.Body.Statements[0])
	}
	testPrintStatement(t, bodyPrintStmt, "10")

	bodyReturnStmt, ok := functionStmt.Body.Statements[1].(*ast.ReturnStatement)
	if !ok {
		t.Fatalf("functionStmt.Body.Statements[0] not *ast.ReturnStatement. got=%T", functionStmt.Body.Statements[1])
	}
	testInfixExpression(t, bodyReturnStmt.ReturnValue, "a", "+", "b")
}
