package lexer

import (
	"testing"

	"github.com/AlyxPink/meowlang/token"
)

func TestNextToken(t *testing.T) {
	input := `
lick x = 5;
lick y = 10;
lick add = meow(a, b) {
	claw a + b;
};
purr add(x, y);
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LICK, "lick"}, {token.IDENT, "x"}, {token.ASSIGN, "="}, {token.INT, "5"}, {token.SEMICOLON, ";"},
		{token.LICK, "lick"}, {token.IDENT, "y"}, {token.ASSIGN, "="}, {token.INT, "10"}, {token.SEMICOLON, ";"},
		{token.LICK, "lick"}, {token.IDENT, "add"}, {token.ASSIGN, "="}, {token.MEOW, "meow"}, {token.LPAREN, "("}, {token.IDENT, "a"}, {token.COMMA, ","}, {token.IDENT, "b"}, {token.RPAREN, ")"}, {token.LBRACE, "{"},
		{token.CLAW, "claw"}, {token.IDENT, "a"}, {token.PLUS, "+"}, {token.IDENT, "b"}, {token.SEMICOLON, ";"},
		{token.RBRACE, "}"}, {token.SEMICOLON, ";"},
		{token.PURR, "purr"}, {token.IDENT, "add"}, {token.LPAREN, "("}, {token.IDENT, "x"}, {token.COMMA, ","}, {token.IDENT, "y"}, {token.RPAREN, ")"}, {token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := NewLexer(input)
	tokens := l.Tokenize()

	compareTokens(t, tokens, tests)
}

func TestVariableDeclaration(t *testing.T) {
	input := `
lick x = 5
lick y = 10;
lick z = 25
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LICK, "lick"}, {token.IDENT, "x"}, {token.ASSIGN, "="}, {token.INT, "5"},
		{token.LICK, "lick"}, {token.IDENT, "y"}, {token.ASSIGN, "="}, {token.INT, "10"}, {token.SEMICOLON, ";"},
		{token.LICK, "lick"}, {token.IDENT, "z"}, {token.ASSIGN, "="}, {token.INT, "25"},
		{token.EOF, ""},
	}

	l := NewLexer(input)
	tokens := l.Tokenize()

	compareTokens(t, tokens, tests)
}

func TestFunctionDefinition(t *testing.T) {
	input := `
meow add(p, q) {
	claw p + q
}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.MEOW, "meow"}, {token.IDENT, "add"}, {token.LPAREN, "("}, {token.IDENT, "p"}, {token.COMMA, ","}, {token.IDENT, "q"}, {token.RPAREN, ")"}, {token.LBRACE, "{"},
		{token.CLAW, "claw"}, {token.IDENT, "p"}, {token.PLUS, "+"}, {token.IDENT, "q"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := NewLexer(input)
	tokens := l.Tokenize()

	compareTokens(t, tokens, tests)
}

func TestConditional(t *testing.T) {
	input := `
hiss (a < b) {
	purr "a is less than b"
} growl {
	purr "a is not less than b"
}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.HISS, "hiss"}, {token.LPAREN, "("}, {token.IDENT, "a"}, {token.LT, "<"}, {token.IDENT, "b"}, {token.RPAREN, ")"}, {token.LBRACE, "{"},
		{token.PURR, "purr"}, {token.STRING, "a is less than b"},
		{token.RBRACE, "}"}, {token.GROWL, "growl"}, {token.LBRACE, "{"},
		{token.PURR, "purr"}, {token.STRING, "a is not less than b"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := NewLexer(input)
	tokens := l.Tokenize()

	compareTokens(t, tokens, tests)
}

func TestLoop(t *testing.T) {
	input := `
scratch (a < b) {
    purr a
    a = a + 1
    nap(1)
}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.SCRATCH, "scratch"}, {token.LPAREN, "("}, {token.IDENT, "a"}, {token.LT, "<"}, {token.IDENT, "b"}, {token.RPAREN, ")"}, {token.LBRACE, "{"},
		{token.PURR, "purr"}, {token.IDENT, "a"},
		{token.IDENT, "a"}, {token.ASSIGN, "="}, {token.IDENT, "a"}, {token.PLUS, "+"}, {token.INT, "1"},
		{token.NAP, "nap"}, {token.LPAREN, "("}, {token.INT, "1"}, {token.RPAREN, ")"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := NewLexer(input)
	tokens := l.Tokenize()

	compareTokens(t, tokens, tests)
}

func TestFunctionCall(t *testing.T) {
	input := `
lick result = add(a, b)
purr "Result of addition: " + result`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LICK, "lick"}, {token.IDENT, "result"}, {token.ASSIGN, "="},
		{token.IDENT, "add"}, {token.LPAREN, "("}, {token.IDENT, "a"}, {token.COMMA, ","}, {token.IDENT, "b"}, {token.RPAREN, ")"},
		{token.PURR, "purr"}, {token.STRING, "Result of addition: "}, {token.PLUS, "+"}, {token.IDENT, "result"},
		{token.EOF, ""},
	}

	l := NewLexer(input)
	tokens := l.Tokenize()

	compareTokens(t, tokens, tests)
}

func TestComments(t *testing.T) {
	input := `
lick a = 5
purr a // This is a comment
/* This is a
multiline comment */
purr a`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LICK, "lick"}, {token.IDENT, "a"}, {token.ASSIGN, "="}, {token.INT, "5"},
		{token.PURR, "purr"}, {token.IDENT, "a"},
		{token.PURR, "purr"}, {token.IDENT, "a"},
		{token.EOF, ""},
	}

	l := NewLexer(input)
	tokens := l.Tokenize()

	compareTokens(t, tokens, tests)
}

func compareTokens(t *testing.T, tokens []token.Token, tests []struct {
	expectedType    token.TokenType
	expectedLiteral string
}) {
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
