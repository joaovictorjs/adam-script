package lexer

import (
	"unicode"
)

type TokenKind int

const (
	EOF TokenKind = iota
	NumericLiteral
	Plus
	Minus
	Star
	Slash
	LParen
	RParen
	Unknown
)

type Token struct {
	Kind     TokenKind
	Lexeme   string
	Position int
}

type Lexer struct {
	source string
	max    int
	index  int
}

func NewLexer(source string) *Lexer {
	lexer := &Lexer{
		source: source,
		max:    len(source),
		index:  0,
	}
	return lexer
}

func (l *Lexer) GenerateTokens() []Token {
	tokens := []Token{}
	for {
		current := l.nextToken()
		tokens = append(tokens, current)
		if current.Kind == EOF {
			break
		}
	}
	return tokens
}

func (l *Lexer) nextToken() Token {
	l.skipWhitespaces()
	if l.index >= l.max {
		tk := Token{Kind: EOF, Lexeme: "", Position: l.index}
		return tk
	}

	current := rune(l.source[l.index])
	if unicode.IsDigit(current) {
		tk := l.lexNumericLiteral()
		return tk
	}

	var kind TokenKind
	switch current {
	case '+':
		kind = Plus
	case '-':
		kind = Minus
	case '*':
		kind = Star
	case '/':
		kind = Slash
	case '(':
		kind = LParen
	case ')':
		kind = RParen
	default:
		kind = Unknown
	}

	token := Token{
		Kind:     kind,
		Lexeme:   string(current),
		Position: l.index,
	}
	l.index++
	return token
}

func (l *Lexer) skipWhitespaces() {
	for l.index < l.max {
		char := l.source[l.index]
		if char == ' ' || char == '\n' || char == '\t' || char == '\r' {
			l.index++
			continue
		}
		break
	}
}

func (l *Lexer) lexNumericLiteral() Token {
	start := l.index
	for l.index < l.max {
		current := rune(l.source[l.index])
		if !unicode.IsDigit(current) {
			break
		}
		l.index++
	}

	lexeme := l.source[start:l.index]
	token := Token{
		Kind:     NumericLiteral,
		Lexeme:   lexeme,
		Position: start,
	}
	return token
}
