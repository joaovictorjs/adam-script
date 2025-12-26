package parser

import (
	"fmt"
	"testing"

	"github.com/joaovictorjs/adam-script/ast"
	"github.com/stretchr/testify/assert"
)

func Test_GivenSource_WhenParse_ThenShouldReturnCorrectProgram(t *testing.T) {
	type TestCase struct {
		name            string
		source          string
		expectedProgram ast.Program
	}

	testcases := []TestCase{
		{
			name:   "single number",
			source: "5",
			expectedProgram: ast.Program{
				Statements: []ast.Statement{
					ast.ExpressionStatement{
						Expression: ast.NumericLiteralExpression{
							Value: 5,
						},
					},
				},
			},
		},
		{
			name:   "simple addition",
			source: "1 + 2",
			expectedProgram: ast.Program{
				Statements: []ast.Statement{
					ast.ExpressionStatement{
						Expression: ast.BinaryExpression{
							Left: ast.NumericLiteralExpression{
								Value: 1,
							},
							Operator: '+',
							Right: ast.NumericLiteralExpression{
								Value: 2,
							},
						},
					},
				},
			},
		},
		{
			name:   "simple subtraction",
			source: "10 - 5",
			expectedProgram: ast.Program{
				Statements: []ast.Statement{
					ast.ExpressionStatement{
						Expression: ast.BinaryExpression{
							Left: ast.NumericLiteralExpression{
								Value: 10,
							},
							Operator: '-',
							Right: ast.NumericLiteralExpression{
								Value: 5,
							},
						},
					},
				},
			},
		},
		{
			name:   "simple multiplication",
			source: "3 * 4",
			expectedProgram: ast.Program{
				Statements: []ast.Statement{
					ast.ExpressionStatement{
						Expression: ast.BinaryExpression{
							Left: ast.NumericLiteralExpression{
								Value: 3,
							},
							Operator: '*',
							Right: ast.NumericLiteralExpression{
								Value: 4,
							},
						},
					},
				},
			},
		},
		{
			name:   "simple division",
			source: "8 / 2",
			expectedProgram: ast.Program{
				Statements: []ast.Statement{
					ast.ExpressionStatement{
						Expression: ast.BinaryExpression{
							Left: ast.NumericLiteralExpression{
								Value: 8,
							},
							Operator: '/',
							Right: ast.NumericLiteralExpression{
								Value: 2,
							},
						},
					},
				},
			},
		},
		{
			name:   "operator precedence multiplication before addition",
			source: "1 + 2 * 3",
			expectedProgram: ast.Program{
				Statements: []ast.Statement{
					ast.ExpressionStatement{
						Expression: ast.BinaryExpression{
							Left: ast.NumericLiteralExpression{
								Value: 1,
							},
							Operator: '+',
							Right: ast.BinaryExpression{
								Left: ast.NumericLiteralExpression{
									Value: 2,
								},
								Operator: '*',
								Right: ast.NumericLiteralExpression{
									Value: 3,
								},
							},
						},
					},
				},
			},
		},
		{
			name:   "operator precedence division before subtraction",
			source: "10 - 6 / 2",
			expectedProgram: ast.Program{
				Statements: []ast.Statement{
					ast.ExpressionStatement{
						Expression: ast.BinaryExpression{
							Left: ast.NumericLiteralExpression{
								Value: 10,
							},
							Operator: '-',
							Right: ast.BinaryExpression{
								Left: ast.NumericLiteralExpression{
									Value: 6,
								},
								Operator: '/',
								Right: ast.NumericLiteralExpression{
									Value: 2,
								},
							},
						},
					},
				},
			},
		},
		{
			name:   "left associativity addition",
			source: "1 + 2 + 3",
			expectedProgram: ast.Program{
				Statements: []ast.Statement{
					ast.ExpressionStatement{
						Expression: ast.BinaryExpression{
							Left: ast.BinaryExpression{
								Left: ast.NumericLiteralExpression{
									Value: 1,
								},
								Operator: '+',
								Right: ast.NumericLiteralExpression{
									Value: 2,
								},
							},
							Operator: '+',
							Right: ast.NumericLiteralExpression{
								Value: 3,
							},
						},
					},
				},
			},
		},
		{
			name:   "left associativity multiplication",
			source: "2 * 3 * 4",
			expectedProgram: ast.Program{
				Statements: []ast.Statement{
					ast.ExpressionStatement{
						Expression: ast.BinaryExpression{
							Left: ast.BinaryExpression{
								Left: ast.NumericLiteralExpression{
									Value: 2,
								},
								Operator: '*',
								Right: ast.NumericLiteralExpression{
									Value: 3,
								},
							},
							Operator: '*',
							Right: ast.NumericLiteralExpression{
								Value: 4,
							},
						},
					},
				},
			},
		},
		{
			name:   "parenthesized expression",
			source: "(1 + 2)",
			expectedProgram: ast.Program{
				Statements: []ast.Statement{
					ast.ExpressionStatement{
						Expression: ast.BinaryExpression{
							Left: ast.NumericLiteralExpression{
								Value: 1,
							},
							Operator: '+',
							Right: ast.NumericLiteralExpression{
								Value: 2,
							},
						},
					},
				},
			},
		},
		{
			name:   "parentheses override precedence",
			source: "(1 + 2) * 3",
			expectedProgram: ast.Program{
				Statements: []ast.Statement{
					ast.ExpressionStatement{
						Expression: ast.BinaryExpression{
							Left: ast.BinaryExpression{
								Left: ast.NumericLiteralExpression{
									Value: 1,
								},
								Operator: '+',
								Right: ast.NumericLiteralExpression{
									Value: 2,
								},
							},
							Operator: '*',
							Right: ast.NumericLiteralExpression{
								Value: 3,
							},
						},
					},
				},
			},
		},
		{
			name:   "nested parentheses",
			source: "((1 + 2) * 3)",
			expectedProgram: ast.Program{
				Statements: []ast.Statement{
					ast.ExpressionStatement{
						Expression: ast.BinaryExpression{
							Left: ast.BinaryExpression{
								Left: ast.NumericLiteralExpression{
									Value: 1,
								},
								Operator: '+',
								Right: ast.NumericLiteralExpression{
									Value: 2,
								},
							},
							Operator: '*',
							Right: ast.NumericLiteralExpression{
								Value: 3,
							},
						},
					},
				},
			},
		},
		{
			name:   "complex expression with all operators",
			source: "1 + 2 * 3 - 4 / 2",
			expectedProgram: ast.Program{
				Statements: []ast.Statement{
					ast.ExpressionStatement{
						Expression: ast.BinaryExpression{
							Left: ast.BinaryExpression{
								Left: ast.NumericLiteralExpression{
									Value: 1,
								},
								Operator: '+',
								Right: ast.BinaryExpression{
									Left: ast.NumericLiteralExpression{
										Value: 2,
									},
									Operator: '*',
									Right: ast.NumericLiteralExpression{
										Value: 3,
									},
								},
							},
							Operator: '-',
							Right: ast.BinaryExpression{
								Left: ast.NumericLiteralExpression{
									Value: 4,
								},
								Operator: '/',
								Right: ast.NumericLiteralExpression{
									Value: 2,
								},
							},
						},
					},
				},
			},
		},
		{
			name:   "complex with multiple parentheses",
			source: "(10 + 20) * (30 - 5)",
			expectedProgram: ast.Program{
				Statements: []ast.Statement{
					ast.ExpressionStatement{
						Expression: ast.BinaryExpression{
							Left: ast.BinaryExpression{
								Left: ast.NumericLiteralExpression{
									Value: 10,
								},
								Operator: '+',
								Right: ast.NumericLiteralExpression{
									Value: 20,
								},
							},
							Operator: '*',
							Right: ast.BinaryExpression{
								Left: ast.NumericLiteralExpression{
									Value: 30,
								},
								Operator: '-',
								Right: ast.NumericLiteralExpression{
									Value: 5,
								},
							},
						},
					},
				},
			},
		},
		{
			name:   "deeply nested expression",
			source: "((10 + 20) * (30 - 5)) / ((8 + 2) * (15 - 3))",
			expectedProgram: ast.Program{
				Statements: []ast.Statement{
					ast.ExpressionStatement{
						Expression: ast.BinaryExpression{
							Left: ast.BinaryExpression{
								Left: ast.BinaryExpression{
									Left: ast.NumericLiteralExpression{
										Value: 10,
									},
									Operator: '+',
									Right: ast.NumericLiteralExpression{
										Value: 20,
									},
								},
								Operator: '*',
								Right: ast.BinaryExpression{
									Left: ast.NumericLiteralExpression{
										Value: 30,
									},
									Operator: '-',
									Right: ast.NumericLiteralExpression{
										Value: 5,
									},
								},
							},
							Operator: '/',
							Right: ast.BinaryExpression{
								Left: ast.BinaryExpression{
									Left: ast.NumericLiteralExpression{
										Value: 8,
									},
									Operator: '+',
									Right: ast.NumericLiteralExpression{
										Value: 2,
									},
								},
								Operator: '*',
								Right: ast.BinaryExpression{
									Left: ast.NumericLiteralExpression{
										Value: 15,
									},
									Operator: '-',
									Right: ast.NumericLiteralExpression{
										Value: 3,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:   "expression with spaces",
			source: "  1   +   2  ",
			expectedProgram: ast.Program{
				Statements: []ast.Statement{
					ast.ExpressionStatement{
						Expression: ast.BinaryExpression{
							Left: ast.NumericLiteralExpression{
								Value: 1,
							},
							Operator: '+',
							Right: ast.NumericLiteralExpression{
								Value: 2,
							},
						},
					},
				},
			},
		},
		{
			name:   "expression without spaces",
			source: "1+2*3",
			expectedProgram: ast.Program{
				Statements: []ast.Statement{
					ast.ExpressionStatement{
						Expression: ast.BinaryExpression{
							Left: ast.NumericLiteralExpression{
								Value: 1,
							},
							Operator: '+',
							Right: ast.BinaryExpression{
								Left: ast.NumericLiteralExpression{
									Value: 2,
								},
								Operator: '*',
								Right: ast.NumericLiteralExpression{
									Value: 3,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			parser := NewParser(test.source)
			program, err := parser.Parse()
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, test.expectedProgram, program)
		})
	}
}

func Test_GivenSourceWithUnexpectedTokens_WhenParse_ThenShouldReturnCorrectError(t *testing.T) {
	type TestCase struct {
		name          string
		source        string
		expectedError error
	}

	testcases := []TestCase{
		{
			name:          "operator at start - plus",
			source:        "+",
			expectedError: fmt.Errorf("Unexpected token '+' at position 0."),
		},
		{
			name:          "operator at start - minus",
			source:        "-",
			expectedError: fmt.Errorf("Unexpected token '-' at position 0."),
		},
		{
			name:          "operator at start - star",
			source:        "*",
			expectedError: fmt.Errorf("Unexpected token '*' at position 0."),
		},
		{
			name:          "operator at start - slash",
			source:        "/",
			expectedError: fmt.Errorf("Unexpected token '/' at position 0."),
		},
		{
			name:          "missing right operand",
			source:        "1 +",
			expectedError: fmt.Errorf("Unexpected token '' at position 3."),
		},
		{
			name:          "consecutive operators",
			source:        "1 + * 2",
			expectedError: fmt.Errorf("Unexpected token '*' at position 4."),
		},
		{
			name:          "missing operator between numbers",
			source:        "1 2",
			expectedError: fmt.Errorf("Unexpected token '2' at position 2."),
		},
		{
			name:          "trailing operator",
			source:        "1 + 2 +",
			expectedError: fmt.Errorf("Unexpected token '' at position 7."),
		},
		{
			name:          "multiplication at start",
			source:        "* 5",
			expectedError: fmt.Errorf("Unexpected token '*' at position 0."),
		},
		{
			name:          "double multiplication",
			source:        "5 * * 2",
			expectedError: fmt.Errorf("Unexpected token '*' at position 4."),
		},
		{
			name:          "double plus operator",
			source:        "1 + + 2",
			expectedError: fmt.Errorf("Unexpected token '+' at position 4."),
		},
		{
			name:          "double division",
			source:        "1 / / 2",
			expectedError: fmt.Errorf("Unexpected token '/' at position 4."),
		},
		{
			name:          "operator after operator - division",
			source:        "1 + / 2",
			expectedError: fmt.Errorf("Unexpected token '/' at position 4."),
		},
		{
			name:          "multiple operators at start",
			source:        "+ + 1",
			expectedError: fmt.Errorf("Unexpected token '+' at position 0."),
		},
		{
			name:          "empty parentheses",
			source:        "()",
			expectedError: fmt.Errorf("Unexpected token ')' at position 1."),
		},
		{
			name:          "operator in parentheses",
			source:        "(+)",
			expectedError: fmt.Errorf("Unexpected token '+' at position 1."),
		},
		{
			name:          "empty parentheses in expression",
			source:        "1 + ()",
			expectedError: fmt.Errorf("Unexpected token ')' at position 5."),
		},
		{
			name:          "operator before closing paren",
			source:        "(1 +)",
			expectedError: fmt.Errorf("Unexpected token ')' at position 4."),
		},
		{
			name:          "operator before closing paren with number after",
			source:        "(1 +) 2",
			expectedError: fmt.Errorf("Unexpected token ')' at position 4."),
		},
		{
			name:          "unclosed left paren",
			source:        "(1 + 2",
			expectedError: fmt.Errorf("Unexpected token '' at position 6."),
		},
		{
			name:          "unmatched right paren",
			source:        "1 + 2)",
			expectedError: fmt.Errorf("Unexpected token ')' at position 5."),
		},
		{
			name:          "nested unclosed parens",
			source:        "((1 + 2)",
			expectedError: fmt.Errorf("Unexpected token '' at position 8."),
		},
		{
			name:          "extra closing paren",
			source:        "(1 + 2))",
			expectedError: fmt.Errorf("Unexpected token ')' at position 7."),
		},
		{
			name:          "multiple nested unclosed parens",
			source:        "1 + (2 + (3 + 4)",
			expectedError: fmt.Errorf("Unexpected token '' at position 16."),
		},
		{
			name:          "lone closing paren",
			source:        ")",
			expectedError: fmt.Errorf("Unexpected token ')' at position 0."),
		},
		{
			name:          "number before opening paren",
			source:        "1 (+ 2)",
			expectedError: fmt.Errorf("Unexpected token '(' at position 2."),
		},
		{
			name:          "number followed by number in parens",
			source:        "5 (3)",
			expectedError: fmt.Errorf("Unexpected token '(' at position 2."),
		},
		{
			name:          "closing paren followed by number",
			source:        "(1 + 2) 3",
			expectedError: fmt.Errorf("Unexpected token '3' at position 8."),
		},
		{
			name:          "closing paren followed by opening paren",
			source:        "(1) (2)",
			expectedError: fmt.Errorf("Unexpected token '(' at position 4."),
		},
		{
			name:          "unknown character alone",
			source:        "@",
			expectedError: fmt.Errorf("Unexpected token '@' at position 0."),
		},
		{
			name:          "unknown character in expression",
			source:        "1 + @ + 2",
			expectedError: fmt.Errorf("Unexpected token '@' at position 4."),
		},
		{
			name:          "hash character in expression",
			source:        "1 # 2",
			expectedError: fmt.Errorf("Unexpected token '#' at position 2."),
		},
		{
			name:          "dollar sign in expression",
			source:        "1 + $ 2",
			expectedError: fmt.Errorf("Unexpected token '$' at position 4."),
		},
		{
			name:          "ampersand in expression",
			source:        "5 & 3",
			expectedError: fmt.Errorf("Unexpected token '&' at position 2."),
		},
		{
			name:          "trailing unknown character",
			source:        "1 + 2 @",
			expectedError: fmt.Errorf("Unexpected token '@' at position 6."),
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			parser := NewParser(test.source)
			_, err := parser.Parse()
			assert.Equal(t, test.expectedError, err)
		})
	}
}
