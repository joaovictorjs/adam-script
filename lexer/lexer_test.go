package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GivenSource_WhenGenerateTokens_ThenShouldReturnCorrectTokenSequence(t *testing.T) {
	type TestCase struct {
		name           string
		source         string
		expectedTokens []Token
	}

	testcases := []TestCase{
		{
			name:   "single digit",
			source: "5",
			expectedTokens: []Token{
				{Kind: NumericLiteral, Lexeme: "5", Position: 0},
				{Kind: EOF, Lexeme: "", Position: 1},
			},
		},
		{
			name:   "multiple digits",
			source: "123",
			expectedTokens: []Token{
				{Kind: NumericLiteral, Lexeme: "123", Position: 0},
				{Kind: EOF, Lexeme: "", Position: 3},
			},
		},
		{
			name:   "plus operator",
			source: "+",
			expectedTokens: []Token{
				{Kind: Plus, Lexeme: "+", Position: 0},
				{Kind: EOF, Lexeme: "", Position: 1},
			},
		},
		{
			name:   "minus operator",
			source: "-",
			expectedTokens: []Token{
				{Kind: Minus, Lexeme: "-", Position: 0},
				{Kind: EOF, Lexeme: "", Position: 1},
			},
		},
		{
			name:   "star operator",
			source: "*",
			expectedTokens: []Token{
				{Kind: Star, Lexeme: "*", Position: 0},
				{Kind: EOF, Lexeme: "", Position: 1},
			},
		},
		{
			name:   "slash operator",
			source: "/",
			expectedTokens: []Token{
				{Kind: Slash, Lexeme: "/", Position: 0},
				{Kind: EOF, Lexeme: "", Position: 1},
			},
		},
		{
			name:   "left paren",
			source: "(",
			expectedTokens: []Token{
				{Kind: LParen, Lexeme: "(", Position: 0},
				{Kind: EOF, Lexeme: "", Position: 1},
			},
		},
		{
			name:   "right paren",
			source: ")",
			expectedTokens: []Token{
				{Kind: RParen, Lexeme: ")", Position: 0},
				{Kind: EOF, Lexeme: "", Position: 1},
			},
		},
		{
			name:   "simple addition",
			source: "1 + 2",
			expectedTokens: []Token{
				{Kind: NumericLiteral, Lexeme: "1", Position: 0},
				{Kind: Plus, Lexeme: "+", Position: 2},
				{Kind: NumericLiteral, Lexeme: "2", Position: 4},
				{Kind: EOF, Lexeme: "", Position: 5},
			},
		},
		{
			name:   "simple subtraction",
			source: "10 - 5",
			expectedTokens: []Token{
				{Kind: NumericLiteral, Lexeme: "10", Position: 0},
				{Kind: Minus, Lexeme: "-", Position: 3},
				{Kind: NumericLiteral, Lexeme: "5", Position: 5},
				{Kind: EOF, Lexeme: "", Position: 6},
			},
		},
		{
			name:   "simple multiplication",
			source: "3 * 4",
			expectedTokens: []Token{
				{Kind: NumericLiteral, Lexeme: "3", Position: 0},
				{Kind: Star, Lexeme: "*", Position: 2},
				{Kind: NumericLiteral, Lexeme: "4", Position: 4},
				{Kind: EOF, Lexeme: "", Position: 5},
			},
		},
		{
			name:   "simple division",
			source: "8 / 2",
			expectedTokens: []Token{
				{Kind: NumericLiteral, Lexeme: "8", Position: 0},
				{Kind: Slash, Lexeme: "/", Position: 2},
				{Kind: NumericLiteral, Lexeme: "2", Position: 4},
				{Kind: EOF, Lexeme: "", Position: 5},
			},
		},
		{
			name:   "addition without spaces",
			source: "1+2",
			expectedTokens: []Token{
				{Kind: NumericLiteral, Lexeme: "1", Position: 0},
				{Kind: Plus, Lexeme: "+", Position: 1},
				{Kind: NumericLiteral, Lexeme: "2", Position: 2},
				{Kind: EOF, Lexeme: "", Position: 3},
			},
		},
		{
			name:   "complex expression without spaces",
			source: "12+34*56",
			expectedTokens: []Token{
				{Kind: NumericLiteral, Lexeme: "12", Position: 0},
				{Kind: Plus, Lexeme: "+", Position: 2},
				{Kind: NumericLiteral, Lexeme: "34", Position: 3},
				{Kind: Star, Lexeme: "*", Position: 5},
				{Kind: NumericLiteral, Lexeme: "56", Position: 6},
				{Kind: EOF, Lexeme: "", Position: 8},
			},
		},
		{
			name:   "parenthesized expression",
			source: "(1 + 2)",
			expectedTokens: []Token{
				{Kind: LParen, Lexeme: "(", Position: 0},
				{Kind: NumericLiteral, Lexeme: "1", Position: 1},
				{Kind: Plus, Lexeme: "+", Position: 3},
				{Kind: NumericLiteral, Lexeme: "2", Position: 5},
				{Kind: RParen, Lexeme: ")", Position: 6},
				{Kind: EOF, Lexeme: "", Position: 7},
			},
		},
		{
			name:   "nested parentheses",
			source: "((1 + 2) * 3)",
			expectedTokens: []Token{
				{Kind: LParen, Lexeme: "(", Position: 0},
				{Kind: LParen, Lexeme: "(", Position: 1},
				{Kind: NumericLiteral, Lexeme: "1", Position: 2},
				{Kind: Plus, Lexeme: "+", Position: 4},
				{Kind: NumericLiteral, Lexeme: "2", Position: 6},
				{Kind: RParen, Lexeme: ")", Position: 7},
				{Kind: Star, Lexeme: "*", Position: 9},
				{Kind: NumericLiteral, Lexeme: "3", Position: 11},
				{Kind: RParen, Lexeme: ")", Position: 12},
				{Kind: EOF, Lexeme: "", Position: 13},
			},
		},
		{
			name:   "complex arithmetic",
			source: "10 + 20 * 30 - 40 / 5",
			expectedTokens: []Token{
				{Kind: NumericLiteral, Lexeme: "10", Position: 0},
				{Kind: Plus, Lexeme: "+", Position: 3},
				{Kind: NumericLiteral, Lexeme: "20", Position: 5},
				{Kind: Star, Lexeme: "*", Position: 8},
				{Kind: NumericLiteral, Lexeme: "30", Position: 10},
				{Kind: Minus, Lexeme: "-", Position: 13},
				{Kind: NumericLiteral, Lexeme: "40", Position: 15},
				{Kind: Slash, Lexeme: "/", Position: 18},
				{Kind: NumericLiteral, Lexeme: "5", Position: 20},
				{Kind: EOF, Lexeme: "", Position: 21},
			},
		},
		{
			name:   "complex with parentheses",
			source: "(10 + 20) * (30 - 5)",
			expectedTokens: []Token{
				{Kind: LParen, Lexeme: "(", Position: 0},
				{Kind: NumericLiteral, Lexeme: "10", Position: 1},
				{Kind: Plus, Lexeme: "+", Position: 4},
				{Kind: NumericLiteral, Lexeme: "20", Position: 6},
				{Kind: RParen, Lexeme: ")", Position: 8},
				{Kind: Star, Lexeme: "*", Position: 10},
				{Kind: LParen, Lexeme: "(", Position: 12},
				{Kind: NumericLiteral, Lexeme: "30", Position: 13},
				{Kind: Minus, Lexeme: "-", Position: 16},
				{Kind: NumericLiteral, Lexeme: "5", Position: 18},
				{Kind: RParen, Lexeme: ")", Position: 19},
				{Kind: EOF, Lexeme: "", Position: 20},
			},
		},
		{
			name:   "multiple spaces",
			source: "1    +    2",
			expectedTokens: []Token{
				{Kind: NumericLiteral, Lexeme: "1", Position: 0},
				{Kind: Plus, Lexeme: "+", Position: 5},
				{Kind: NumericLiteral, Lexeme: "2", Position: 10},
				{Kind: EOF, Lexeme: "", Position: 11},
			},
		},
		{
			name:   "empty string",
			source: "",
			expectedTokens: []Token{
				{Kind: EOF, Lexeme: "", Position: 0},
			},
		},
		{
			name:   "only spaces",
			source: "   ",
			expectedTokens: []Token{
				{Kind: EOF, Lexeme: "", Position: 3},
			},
		},
		{
			name:   "unknown character",
			source: "1 + @ 2",
			expectedTokens: []Token{
				{Kind: NumericLiteral, Lexeme: "1", Position: 0},
				{Kind: Plus, Lexeme: "+", Position: 2},
				{Kind: Unknown, Lexeme: "@", Position: 4},
				{Kind: NumericLiteral, Lexeme: "2", Position: 6},
				{Kind: EOF, Lexeme: "", Position: 7},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()
			lexer := NewLexer(testcase.source)
			tokens := lexer.GenerateTokens()
			assert.Equal(t, testcase.expectedTokens, tokens)
		})
	}
}
