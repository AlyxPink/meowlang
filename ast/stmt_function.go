package ast

import (
	"github.com/AlyxPink/meowlang/token"
)

type FunctionStatement struct {
	Token      token.Token // the token.MEOW token
	Name       *Identifier
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fs *FunctionStatement) statementNode() {}

func (fs *FunctionStatement) TokenLiteral() string {
	return fs.Token.Literal
}
