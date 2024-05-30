package interpreter

import (
	"fmt"

	"github.com/AlyxPink/meowlang/ast"
	"github.com/AlyxPink/meowlang/object"
)

type Interpreter struct {
	env *object.Environment
}

func NewInterpreter() *Interpreter {
	return &Interpreter{env: object.NewEnvironment()}
}

func (i *Interpreter) Interpret(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return i.evalProgram(node)
	case *ast.AssignStatement:
		return i.evalAssignStatement(node)
	case *ast.FunctionStatement:
		return i.evalFunctionStatement(node)
	case *ast.ReturnStatement:
		return i.evalReturnStatement(node)
	case *ast.Identifier:
		return i.evalIdentifier(node)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.InfixExpression:
		return i.evalInfixExpression(node)
	}
	return nil
}

func (i *Interpreter) evalProgram(program *ast.Program) object.Object {
	var result object.Object
	for _, stmt := range program.Statements {
		result = i.Interpret(stmt)
	}
	return result
}

func (i *Interpreter) evalAssignStatement(stmt *ast.AssignStatement) object.Object {
	val := i.Interpret(stmt.Value)
	if val != nil {
		i.env.Set(stmt.Name.Value, val)
	}
	return val
}

func (i *Interpreter) evalFunctionStatement(stmt *ast.FunctionStatement) object.Object {
	// For now, just print the function definition for demonstration
	fmt.Printf("Function %s with parameters %v and body %v\n", stmt.Name.Value, stmt.Parameters, stmt.Body)
	return nil
}

func (i *Interpreter) evalReturnStatement(stmt *ast.ReturnStatement) object.Object {
	val := i.Interpret(stmt.ReturnValue)
	return val
}

func (i *Interpreter) evalIdentifier(node *ast.Identifier) object.Object {
	if val, ok := i.env.Get(node.Value); ok {
		return val
	}
	return nil
}

func (i *Interpreter) evalInfixExpression(exp *ast.InfixExpression) object.Object {
	left := i.Interpret(exp.Left)
	right := i.Interpret(exp.Right)

	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		return i.evalIntegerInfixExpression(exp.Operator, left, right)
	}

	return nil
}

func (i *Interpreter) evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	default:
		return nil
	}
}
