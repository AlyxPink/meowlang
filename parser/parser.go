package parser

import (
	"strconv"

	"github.com/AlyxPink/meowlang/ast"
	"github.com/AlyxPink/meowlang/token"
)

type Parser struct {
	tokens  []token.Token
	current int
	errors  []string
}

func NewParser(tokens []token.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

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

func (p *Parser) parseStatement() ast.Statement {
	switch p.peek().Type {
	case token.LICK:
		return p.parseAssignStatement()
	case token.MEOW:
		return p.parseFunctionStatement()
	case token.CLAW:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseAssignStatement() *ast.AssignStatement {
	stmt := &ast.AssignStatement{Token: p.advance()} // consume 'lick' token

	stmt.Name = p.parseIdentifier() // parse identifier

	if !p.expectPeek(token.ASSIGN) { // consume assign token
		return nil
	}

	stmt.Value = p.parseExpression()

	return stmt
}

func (p *Parser) parseFunctionStatement() *ast.FunctionStatement {
	stmt := &ast.FunctionStatement{Token: p.advance()} // consume 'meow' token

	stmt.Name = p.parseIdentifier() // parse identifier

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	// Parse function parameters
	stmt.Parameters = p.parseFunctionParameters()

	// Parse function body
	stmt.Body = p.parseBlockStatement()

	return stmt
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	var parameters []*ast.Identifier

	// Check if function has no parameters
	if p.peek().Type == token.RPAREN {
		p.advance() // consume ')' token
		return parameters
	}

	param := p.parseIdentifier()
	parameters = append(parameters, param)

	for p.peek().Type == token.COMMA {
		p.advance() // consume ',' token

		param := p.parseIdentifier()
		parameters = append(parameters, param)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return parameters
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{
		Token: p.peek(),
	}
	block.Statements = []ast.Statement{}

	p.advance() // consume '{' token

	for p.peek().Type != token.RBRACE && !p.isAtEnd() {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.advance()
	}

	return block
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.advance()} // consume 'claw' token

	stmt.ReturnValue = p.parseExpression()

	return stmt
}

func (p *Parser) parseExpression() ast.Expression {
	return p.parseInfixExpression()
}

func (p *Parser) parseInfixExpression() ast.Expression {
	left := p.parsePrimary()

	for p.peek().Type == token.PLUS || p.peek().Type == token.MINUS ||
		p.peek().Type == token.ASTERISK || p.peek().Type == token.SLASH {
		operator := p.advance()
		right := p.parsePrimary()
		left = &ast.InfixExpression{
			Token:    operator,
			Left:     left,
			Operator: operator.Literal,
			Right:    right,
		}
	}

	return left
}

func (p *Parser) parsePrimary() ast.Expression {
	switch p.peek().Type {
	case token.INT:
		return p.parseIntegerLiteral()
	case token.IDENT:
		return p.parseIdentifier()
	default:
		return nil
	}
}

func (p *Parser) parseIntegerLiteral() *ast.IntegerLiteral {
	lit := &ast.IntegerLiteral{Token: p.advance()} // consume integer token

	value, err := strconv.ParseInt(lit.Token.Literal, 0, 64)
	if err != nil {
		p.errors = append(p.errors, "could not parse "+lit.Token.Literal+" as integer")
		return nil
	}

	lit.Value = value
	return lit
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	ident := &ast.Identifier{
		Token: p.peek(),
		Value: p.peek().Literal,
	}

	if !p.expectPeek(token.IDENT) { // consume identifier token
		return nil
	}

	return ident
}

func (p *Parser) advance() token.Token {
	tok := p.tokens[p.current]
	p.current++
	return tok
}

func (p *Parser) peek() token.Token {
	return p.tokens[p.current]
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peek().Type == t {
		p.advance()
		return true
	} else {
		p.errors = append(p.errors, "expected next token to be "+string(t)+", got "+string(p.peek().Type)+" instead")
		return false
	}
}

func (p *Parser) isAtEnd() bool {
	return p.current >= len(p.tokens) || p.peek().Type == token.EOF
}
