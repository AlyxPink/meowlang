package token

type TokenType string

const (
	// Special tokens
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + literals
	IDENT  TokenType = "IDENT"
	INT    TokenType = "INT"
	STRING TokenType = "STRING"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	CLAW    = "CLAW"
	GROWL   = "GROWL"
	HISS    = "HISS"
	LICK    = "LICK"
	MEOW    = "MEOW"
	NAP     = "NAP"
	PURR    = "PURR"
	SCRATCH = "SCRATCH"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"claw":    CLAW,
	"growl":   GROWL,
	"hiss":    HISS,
	"lick":    LICK,
	"meow":    MEOW,
	"nap":     NAP,
	"purr":    PURR,
	"scratch": SCRATCH,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
