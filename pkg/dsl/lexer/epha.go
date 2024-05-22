package lexer

import (
	"epha/pkg/dsl/token"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           rune // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch, _ = utf8.DecodeRuneInString(l.input[l.readPosition:])
	}
	l.position = l.readPosition
	l.readPosition += utf8.RuneLen(l.ch)
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.Assign, string(l.ch))
	case '(':
		tok = newToken(token.LParen, string(l.ch))
	case ')':
		tok = newToken(token.RParen, string(l.ch))
	case ',':
		tok = newToken(token.Comma, string(l.ch))
	case '"':
		tok.Type = token.String
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tok.Type = l.identifyType(literal)
			tok.Literal = literal
			return tok
		} else {
			tok = newToken(token.Illegal, string(l.ch))
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || l.ch == '.' {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) identifyType(identifier string) token.TokenType {
	switch identifier {
	case "import":
		return token.Import
	case "as":
		return token.As
	// Expand to other specific keywords as needed
	default:
		return token.Identifier
	}
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}
