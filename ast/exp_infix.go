package ast

import "github.com/AlyxPink/meowlang/token"

type InfixExpression struct {
	Token    token.Token // The operator token, e.g., '+'
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}
