package lexer

import (
	"walli/internal/lexer/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = token.NewToken(token.ASSIGN, l.ch)
		break
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
		break
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
		break
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
		break
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
		break
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
		break
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
		break
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
		break
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		break
	default:
		if isLetter(l.ch) {
			position := l.position
			for isLetter(l.ch) {
				l.readChar()
			}
			tok.Literal = l.input[position:l.position]
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
