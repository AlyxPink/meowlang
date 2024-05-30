package ast

import "github.com/AlyxPink/meowlang/token"

type PrintStatement struct {
	Token token.Token // the token.PURR token
	Value Expression
}

func (ps *PrintStatement) statementNode() {}

func (ps *PrintStatement) TokenLiteral() string {
	return ps.Token.Literal
}
