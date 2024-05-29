package ast

import (
	"github.com/AlyxPink/meowlang/token"
)

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
