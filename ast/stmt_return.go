package ast

import "github.com/AlyxPink/meowlang/token"

type ReturnStatement struct {
	Token       token.Token // the token.CLAW token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
