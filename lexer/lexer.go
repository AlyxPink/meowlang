package lexer

import (
	"unicode"

	"github.com/AlyxPink/meowlang/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) Tokenize() []token.Token {
	var tokens []token.Token
	for l.ch != 0 {
		switch l.ch {
		case '=':
			tokens = append(tokens, token.Token{Type: token.ASSIGN, Literal: string(l.ch)})
		case '+':
			tokens = append(tokens, token.Token{Type: token.PLUS, Literal: string(l.ch)})
		case '-':
			tokens = append(tokens, token.Token{Type: token.MINUS, Literal: string(l.ch)})
		case '*':
			tokens = append(tokens, token.Token{Type: token.ASTERISK, Literal: string(l.ch)})
		case ';':
			tokens = append(tokens, token.Token{Type: token.SEMICOLON, Literal: string(l.ch)})
		case '(':
			tokens = append(tokens, token.Token{Type: token.LPAREN, Literal: string(l.ch)})
		case ')':
			tokens = append(tokens, token.Token{Type: token.RPAREN, Literal: string(l.ch)})
		case '{':
			tokens = append(tokens, token.Token{Type: token.LBRACE, Literal: string(l.ch)})
		case '}':
			tokens = append(tokens, token.Token{Type: token.RBRACE, Literal: string(l.ch)})
		case '>':
			tokens = append(tokens, token.Token{Type: token.GT, Literal: string(l.ch)})
		case '<':
			tokens = append(tokens, token.Token{Type: token.LT, Literal: string(l.ch)})
		case ',':
			tokens = append(tokens, token.Token{Type: token.COMMA, Literal: string(l.ch)})
		case '/': // Comment or division operator
			if l.peekChar() == '/' {
				l.skipSingleLineComment()
			} else if l.peekChar() == '*' {
				l.skipBlockComment()
			} else {
				tokens = append(tokens, token.Token{Type: token.SLASH, Literal: string(l.ch)})
			}
		default:
			if isSpace(l.ch) {
				l.readChar()
				continue
			} else if isLetter(l.ch) {
				literal := l.readIdentifier()
				tokens = append(tokens, token.Token{Type: token.LookupIdent(literal), Literal: literal})
				continue
			} else if isDigit(l.ch) {
				tokens = append(tokens, token.Token{Type: token.INT, Literal: l.readNumber()})
				continue
			} else {
				tokens = append(tokens, token.Token{Type: token.ILLEGAL, Literal: string(l.ch)})
			}
		}
		l.readChar()
	}
	tokens = append(tokens, token.Token{Type: token.EOF, Literal: ""})
	return tokens
}

func (l *Lexer) readChar() {
	l.peekChar()
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	return l.ch
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipSingleLineComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}

func (l *Lexer) skipBlockComment() {
	l.readChar() // consume '*'
	for {
		if l.ch == '*' && l.peekChar() == '/' {
			l.readChar() // consume '*'
			l.readChar() // consume '/'
			break
		}
		l.readChar()
		if l.ch == 0 {
			break
		}
	}
}

func isSpace(ch byte) bool {
	return unicode.IsSpace(rune(ch))
}

func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch))
}

func isDigit(ch byte) bool {
	return unicode.IsDigit(rune(ch))
}
