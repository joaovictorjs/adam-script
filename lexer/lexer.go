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
	Let
	Const
	Identifier
	Equals
	Semicolon
	StringLiteral
)

var keywords = map[string]TokenKind{"let": Let, "const": Const}

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

	if current == '"' {
		tk := l.lexString()
		return tk
	}

	if unicode.IsLetter(current) {
		tk := l.lexMultichar()
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
	case '=':
		kind = Equals
	case ';':
		kind = Semicolon
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

func (l *Lexer) lexString() Token {
	start := l.index
	l.index++
	for l.index < l.max {
		char := l.source[l.index]
		if char == '\\' && l.index+1 < l.max {
			l.index += 2
			continue
		}

		if char == '"' {
			l.index++
			lexeme := l.source[start:l.index]
			token := Token{
				Kind:     StringLiteral,
				Lexeme:   lexeme,
				Position: start,
			}
			return token
		}
		l.index++
	}

	lexeme := l.source[start:l.index]
	token := Token{
		Kind:     Unknown,
		Lexeme:   lexeme,
		Position: start,
	}
	return token
}

func (l *Lexer) lexMultichar() Token {
	start := l.index
	for l.index < l.max {
		char := rune(l.source[l.index])
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) && !unicode.IsDigit(char) && char != '_' {
			break
		}
		l.index++
	}

	lexeme := l.source[start:l.index]
	var tokenKind TokenKind
	if kind, ok := keywords[lexeme]; ok {
		tokenKind = kind
	} else {
		tokenKind = Identifier
	}

	token := Token{
		Kind:     tokenKind,
		Lexeme:   lexeme,
		Position: start,
	}
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
