package parser

import (
	"strconv"

	"github.com/AlyxPink/meowlang/ast"
	"github.com/AlyxPink/meowlang/token"
)

// Precedence levels
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

// precedences maps token types to their precedence levels.
var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
}

// Parser represents a parser for the MeowLang programming language.
type Parser struct {
	tokens  []token.Token
	current int
	errors  []string
}

// NewParser creates a new instance of Parser.
func NewParser(tokens []token.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

// ParseProgram parses the entire input and returns the root of the AST.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.isAtEnd() {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		} else {
			p.advance() // Advance only if no valid statement was parsed to avoid infinite loop
		}
	}

	return program
}

// parseStatement parses a single statement.
func (p *Parser) parseStatement() ast.Statement {
	switch p.peek().Type {
	case token.LICK:
		return p.parseAssignStatement()
	case token.MEOW:
		return p.parseFunctionStatement()
	case token.CLAW:
		return p.parseReturnStatement()
	case token.PURR:
		return p.parsePrintStatement()
	default:
		return nil
	}
}

// parseAssignStatement parses an assignment statement.
func (p *Parser) parseAssignStatement() *ast.AssignStatement {
	stmt := &ast.AssignStatement{
		Token: p.advance(), // consume 'lick' token
	}

	stmt.Name = &ast.Identifier{
		Token: p.peek(),
		Value: p.peek().Literal,
	}
	if !p.expectPeek(token.IDENT) { // consume identifier token
		return nil
	}

	if !p.expectPeek(token.ASSIGN) { // consume assign token
		return nil
	}

	stmt.Value = p.parseExpression(LOWEST)

	if p.peek().Type == token.SEMICOLON {
		p.advance() // consume optional semicolon token
	}

	return stmt
}

// parseFunctionStatement parses a function definition statement.
func (p *Parser) parseFunctionStatement() *ast.FunctionStatement {
	stmt := &ast.FunctionStatement{
		Token: p.advance(), // consume 'meow' token
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{
		Token: p.peek(),
		Value: p.previous().Literal,
	}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	stmt.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

// parseFunctionParameters parses the parameters of a function.
func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	var parameters []*ast.Identifier

	if p.peek().Type == token.RPAREN {
		p.advance()
		return parameters
	}

	p.advance() // consume first parameter
	param := &ast.Identifier{
		Token: p.previous(),
		Value: p.previous().Literal,
	}
	parameters = append(parameters, param)

	for p.peek().Type == token.COMMA {
		p.advance() // consume ',' token

		p.advance() // consume next parameter
		param := &ast.Identifier{
			Token: p.previous(),
			Value: p.previous().Literal,
		}
		parameters = append(parameters, param)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return parameters
}

// parseBlockStatement parses a block of statements enclosed in curly braces.
func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{
		Token: p.peek(),
	}
	block.Statements = []ast.Statement{}

	for !p.isAtEnd() && p.peek().Type != token.RBRACE {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		} else {
			p.advance() // Advance to avoid infinite loop if no valid statement was parsed
		}
	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return block
}

// parseReturnStatement parses a return statement.
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{
		Token: p.advance(), // consume 'claw' token
	}

	stmt.ReturnValue = p.parseExpression(LOWEST)

	if p.peek().Type == token.SEMICOLON {
		p.advance()
	}

	return stmt
}

// parsePrintStatement parses a print statement.
func (p *Parser) parsePrintStatement() *ast.PrintStatement {
	stmt := &ast.PrintStatement{
		Token: p.advance(), // consume 'purr' token
	}

	stmt.Value = p.parseExpression(LOWEST)

	if p.peek().Type == token.SEMICOLON {
		p.advance() // consume optional semicolon token
	}

	return stmt
}

// parseCallExpression parses a function call expression.
func (p *Parser) parseCallExpression(function ast.Expression) *ast.CallExpression {
	exp := &ast.CallExpression{
		Token:    p.peek(),
		Function: function,
	}

	exp.Arguments = p.parseExpressionList(token.RPAREN)
	return exp
}

// parseExpressionList parses a list of expressions, separated by commas, and ending with a specified token.
func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	var list []ast.Expression

	if p.peek().Type == end {
		p.advance()
		return list
	}

	p.advance()
	list = append(list, p.parseExpression(LOWEST))

	for p.peek().Type == token.COMMA {
		p.advance() // consume ','
		p.advance()
		list = append(list, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}

// parseExpression is the main entry point for parsing expressions. It handles
// all kinds of expressions, taking into account operator precedence and associativity.
func (p *Parser) parseExpression(precedence int) ast.Expression {
	// Parse the left-hand side of the expression
	leftExp := p.parsePrimary()

	// Handle infix operators and function calls
	for !p.isAtEnd() && precedence < p.peekPrecedence() {
		infix := p.peek()

		if infix.Type != token.PLUS && infix.Type != token.MINUS &&
			infix.Type != token.SLASH && infix.Type != token.ASTERISK &&
			infix.Type != token.LPAREN {
			return leftExp
		}

		// Handle function calls
		if infix.Type == token.LPAREN {
			leftExp = p.parseCallExpression(leftExp)
		} else {
			// Handle infix expressions
			p.advance()
			leftExp = p.parseInfixExpression(leftExp)
		}
	}

	return leftExp
}

// parsePrimary is responsible for parsing the simplest (atomic) units of expressions,
// which do not involve any operators. These are the building blocks of more complex expressions.
func (p *Parser) parsePrimary() ast.Expression {
	switch p.peek().Type {
	case token.INT:
		return p.parseIntegerLiteral()
	case token.STRING:
		return p.parseStringLiteral()
	case token.IDENT:
		return p.parseIdentifier()
	case token.LPAREN:
		p.advance()
		expr := p.parseExpression(LOWEST)
		if !p.expectPeek(token.RPAREN) {
			return nil
		}
		return expr
	default:
		return nil
	}
}

// parseIntegerLiteral parses an integer literal.
func (p *Parser) parseIntegerLiteral() *ast.IntegerLiteral {
	lit := &ast.IntegerLiteral{
		Token: p.advance(),
	}

	value, err := strconv.ParseInt(lit.Token.Literal, 0, 64)
	if err != nil {
		p.errors = append(p.errors, "could not parse "+lit.Token.Literal+" as integer")
		return nil
	}

	lit.Value = value
	return lit
}

// parseStringLiteral parses a string literal.
func (p *Parser) parseStringLiteral() *ast.StringLiteral {
	lit := &ast.StringLiteral{
		Token: p.advance(),
		Value: p.previous().Literal,
	}
	return lit
}

// parseIdentifier parses an identifier.
func (p *Parser) parseIdentifier() *ast.Identifier {
	ident := &ast.Identifier{
		Token: p.advance(),
		Value: p.previous().Literal,
	}
	return ident
}

// parseInfixExpression parses an infix expression.
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token:    p.previous(), // The operator token
		Left:     left,
		Operator: p.previous().Literal,
	}

	precedence := p.currentPrecedence()
	exp.Right = p.parseExpression(precedence)

	return exp
}

// Helper methods

// advance advances the parser to the next token.
func (p *Parser) advance() token.Token {
	tok := p.tokens[p.current]
	p.current++
	return tok
}

// previous returns the previous token.
func (p *Parser) previous() token.Token {
	return p.tokens[p.current-1]
}

// peek returns the current token without advancing the parser.
func (p *Parser) peek() token.Token {
	return p.tokens[p.current]
}

// expectPeek checks if the next token is of the expected type and advances the parser if it is.
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peek().Type == t {
		p.advance()
		return true
	} else {
		p.errors = append(p.errors, "expected next token to be "+string(t)+", got "+string(p.peek().Type)+" instead")
		return false
	}
}

// isAtEnd checks if the parser has reached the end of the token stream.
func (p *Parser) isAtEnd() bool {
	return p.current >= len(p.tokens) || p.peek().Type == token.EOF
}

// currentPrecedence returns the precedence of the current token.
func (p *Parser) currentPrecedence() int {
	if p.isAtEnd() {
		return LOWEST
	}
	if prec, ok := precedences[p.previous().Type]; ok {
		return prec
	}
	return LOWEST
}

// peekPrecedence returns the precedence of the next token.
func (p *Parser) peekPrecedence() int {
	if p.isAtEnd() {
		return LOWEST
	}
	if prec, ok := precedences[p.peek().Type]; ok {
		return prec
	}
	return LOWEST
}
