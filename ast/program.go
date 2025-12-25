package ast

import "encoding/json"

type Program struct {
	Statements []Statement
}

func (Program) node() {}

func (n Program) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Kind":       "Program",
		"Statements": n.Statements,
	})
}
