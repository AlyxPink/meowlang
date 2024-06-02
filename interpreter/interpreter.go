package interpreter

import (
	"bytes"
	"fmt"

	"github.com/AlyxPink/meowlang/ast"
	"github.com/AlyxPink/meowlang/object"
)

// Interpreter represents the interpreter for the MeowLang programming language.
type Interpreter struct {
	env *object.Environment
	out *bytes.Buffer
}

// NewInterpreter creates a new instance of Interpreter.
func NewInterpreter() *Interpreter {
	return &Interpreter{env: object.NewEnvironment(), out: new(bytes.Buffer)}
}

// NewInterpreterWithOutput creates a new instance of Interpreter with a specified output buffer.
func NewInterpreterWithOutput(out *bytes.Buffer) *Interpreter {
	env := object.NewEnvironment()
	return &Interpreter{env: env, out: out}
}

// NewInterpreterWithEnv creates a new instance of Interpreter with a specified environment.
func NewInterpreterWithEnv(env *object.Environment) *Interpreter {
	return &Interpreter{env: env, out: new(bytes.Buffer)}
}

// Interpret interprets the given AST node and returns the resulting object.
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
	case *ast.PrintStatement:
		return i.evalPrintStatement(node)
	case *ast.BlockStatement:
		return i.evalBlockStatement(node)
	case *ast.CallExpression:
		return i.evalCallExpression(node)
	case *ast.Identifier:
		return i.evalIdentifier(node)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.InfixExpression:
		return i.evalInfixExpression(node)
	}
	return &object.Null{}
}

// evalProgram evaluates the given program node.
func (i *Interpreter) evalProgram(program *ast.Program) object.Object {
	var result object.Object
	for _, stmt := range program.Statements {
		result = i.Interpret(stmt)
	}
	return result
}

// evalAssignStatement evaluates an assignment statement.
func (i *Interpreter) evalAssignStatement(stmt *ast.AssignStatement) object.Object {
	val := i.Interpret(stmt.Value)
	if val != nil {
		i.env.Set(stmt.Name.Value, val)
	}
	return val
}

// evalFunctionStatement evaluates a function definition statement.
func (i *Interpreter) evalFunctionStatement(stmt *ast.FunctionStatement) object.Object {
	params := make([]*object.Identifier, len(stmt.Parameters))
	for index, param := range stmt.Parameters {
		params[index] = &object.Identifier{Name: param.Value}
	}

	body := stmt.Body

	function := &object.Function{
		Parameters: params,
		Body:       body,
		Env:        i.env,
	}

	i.env.Set(stmt.Name.Value, function)

	return function
}

// evalReturnStatement evaluates a return statement.
func (i *Interpreter) evalReturnStatement(stmt *ast.ReturnStatement) object.Object {
	val := i.Interpret(stmt.ReturnValue)
	return val
}

// evalPrintStatement evaluates a print statement.
func (i *Interpreter) evalPrintStatement(stmt *ast.PrintStatement) object.Object {
	val := i.Interpret(stmt.Value)
	if val != nil {
		fmt.Fprintln(i.out, val.Inspect())
	}
	return &object.Null{}
}

// evalBlockStatement evaluates a block of statements.
func (i *Interpreter) evalBlockStatement(block *ast.BlockStatement) object.Object {
	var result object.Object
	for _, stmt := range block.Statements {
		result = i.Interpret(stmt)
	}
	return result
}

// evalCallExpression evaluates a function call expression.
func (i *Interpreter) evalCallExpression(exp *ast.CallExpression) object.Object {
	function := i.Interpret(exp.Function)
	if function == nil {
		return &object.Null{}
	}

	args := make([]object.Object, len(exp.Arguments))
	for index, arg := range exp.Arguments {
		args[index] = i.Interpret(arg)
	}

	return i.applyFunction(function, args)
}

// applyFunction applies a function to its arguments.
func (i *Interpreter) applyFunction(fn object.Object, args []object.Object) object.Object {
	function, ok := fn.(*object.Function)
	if !ok {
		return &object.Null{}
	}

	extendedEnv := object.NewEnclosedEnvironment(function.Env)

	for paramIdx, param := range function.Parameters {
		extendedEnv.Set(param.Name, args[paramIdx])
	}

	evaluator := NewInterpreterWithEnv(extendedEnv)
	result := evaluator.Interpret(function.Body)

	return result
}

// evalIdentifier evaluates an identifier by looking it up in the environment.
func (i *Interpreter) evalIdentifier(node *ast.Identifier) object.Object {
	if val, ok := i.env.Get(node.Value); ok {
		return val
	}
	// If the identifier is not found, return an error or null
	return &object.Null{}
}

// evalInfixExpression evaluates an infix expression.
func (i *Interpreter) evalInfixExpression(exp *ast.InfixExpression) object.Object {
	left := i.Interpret(exp.Left)
	right := i.Interpret(exp.Right)

	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return i.evalIntegerInfixExpression(exp.Operator, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return i.evalStringInfixExpression(exp.Operator, left, right)
	}

	return &object.Null{}
}

// evalIntegerInfixExpression evaluates an infix expression with integer operands.
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
		return &object.Null{}
	}
}

// evalStringInfixExpression evaluates an infix expression with string operands.
func (i *Interpreter) evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	if operator != "+" {
		return &object.Null{} // or handle as an error
	}

	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	return &object.String{Value: leftVal + rightVal}
}
