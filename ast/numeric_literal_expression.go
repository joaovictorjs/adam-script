package ast

import "encoding/json"

type NumericLiteralExpression struct {
	Value float64
}

func (NumericLiteralExpression) node() {}

func (NumericLiteralExpression) expression() {}

func (e NumericLiteralExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Kind":  "NumericLiteralExpression",
		"Value": e.Value,
	})
}
