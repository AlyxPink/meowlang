package ast

import (
	"github.com/AlyxPink/meowlang/token"
)

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
