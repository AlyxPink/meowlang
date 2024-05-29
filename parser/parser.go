package parser

import (
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
	lit := &ast.IntegerLiteral{
		Token: p.peek(),
	}

	if !p.expectPeek(token.INT) { // consume integer token
		return nil
	}

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
