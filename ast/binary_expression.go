package ast

import "encoding/json"

type BinaryExpression struct {
	Left     Expression
	Operator uint8
	Right    Expression
}

func (BinaryExpression) node() {}

func (BinaryExpression) expression() {}

func (e BinaryExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Kind":     "BinaryExpression",
		"Left":     e.Left,
		"Operator": string(e.Operator),
		"Right":    e.Right,
	})
}
