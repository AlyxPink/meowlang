package ast

import "github.com/AlyxPink/meowlang/token"

type AssignStatement struct {
	Token token.Token // the token.LICK token
	Name  *Identifier
	Value Expression
}

func (ls *AssignStatement) statementNode() {}

func (ls *AssignStatement) TokenLiteral() string {
	return ls.Token.Literal
}
