package lexer

import (
	"testing"

	"github.com/AlyxPink/meowlang/token"
)

func TestNextToken(t *testing.T) {
	input := `lick x = 5;
              lick y = 10;
              lick add = meow(a, b) {
                  claw a + b;
              };
              purr add(x, y);`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LICK, "lick"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LICK, "lick"},
		{token.IDENT, "y"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LICK, "lick"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.MEOW, "meow"},
		{token.LPAREN, "("},
		{token.IDENT, "a"},
		{token.COMMA, ","},
		{token.IDENT, "b"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.CLAW, "claw"},
		{token.IDENT, "a"},
		{token.PLUS, "+"},
		{token.IDENT, "b"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.PURR, "purr"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := NewLexer(input)
	tokens := l.Tokenize()

	for i, tt := range tests {
		tok := tokens[i]

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
