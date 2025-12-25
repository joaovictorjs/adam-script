package ast

import "encoding/json"

type ExpressionStatement struct {
	Expression Expression
}

func (ExpressionStatement) node() {}

func (ExpressionStatement) statement() {}

func (e ExpressionStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Kind":       "ExpressionStatement",
		"Expression": e.Expression,
	})
}
