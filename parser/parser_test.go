package parser

import (
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
