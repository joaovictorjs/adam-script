package parser

import (
	"fmt"
	"strconv"

	"github.com/joaovictorjs/adam-script/ast"
	"github.com/joaovictorjs/adam-script/lexer"
)

type Parser struct {
	tokens []lexer.Token
	max    int
	index  int
}

func NewParser(source string) *Parser {
	lexer := lexer.NewLexer(source)
	tokens := lexer.GenerateTokens()

	return &Parser{
		tokens: tokens,
		max:    len(tokens),
		index:  0,
	}
}

func (p *Parser) Parse() (ast.Program, error) {
	program := ast.Program{}

	for {
		token := p.peek()
		if token.Kind == lexer.EOF {
			return program, nil
		}

		statement, err := p.parseStatement()
		if err != nil {
			return program, err
		}

		program.Statements = append(program.Statements, statement)
	}
}

func (p *Parser) parseStatement() (ast.Statement, error) {
	expr, err := p.parseAdditiveExpression()
	if err != nil {
		return nil, err
	}

	statement := ast.ExpressionStatement{
		Expression: expr,
	}
	return statement, nil
}

func (p *Parser) parseAdditiveExpression() (ast.Expression, error) {
	left, err := p.parseMultiplicativeExpression()
	if err != nil {
		return nil, err
	}

	for {
		token := p.peek()
		if token.Kind != lexer.Plus && token.Kind != lexer.Minus {
			break
		}

		operator := token.Lexeme[0]
		p.index++
		right, err := p.parseMultiplicativeExpression()
		if err != nil {
			return nil, err
		}

		left = ast.BinaryExpression{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}

	return left, nil
}

func (p *Parser) parseMultiplicativeExpression() (ast.Expression, error) {
	left, err := p.parsePrimaryExpression()
	if err != nil {
		return nil, err
	}

	for {
		token := p.peek()
		if token.Kind != lexer.Star && token.Kind != lexer.Slash {
			break
		}

		operator := token.Lexeme[0]
		p.index++
		right, err := p.parsePrimaryExpression()
		if err != nil {
			return nil, err
		}

		left = ast.BinaryExpression{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}

	return left, nil
}

func (p *Parser) parsePrimaryExpression() (ast.Expression, error) {
	token := p.peek()

	switch token.Kind {
	case lexer.NumericLiteral:
		{
			valueAsFloat, err := strconv.ParseFloat(token.Lexeme, 64)
			if err != nil {
				return nil, err
			}

			p.index++
			expr := ast.NumericLiteralExpression{Value: valueAsFloat}
			return expr, nil
		}
	case lexer.LParen:
		{
			p.index++
			expr, err := p.parseAdditiveExpression()
			p.index++
			if err != nil {
				return nil, err
			}

			return expr, nil
		}
	default:
		{
			p.index++
			err := fmt.Errorf("Unexpected token '%s' at position %d.", token.Lexeme, token.Position)
			return nil, err
		}
	}
}

func (p *Parser) peek() lexer.Token {
	if p.index < p.max {
		return p.tokens[p.index]
	}
	return p.tokens[p.max-1]
}
