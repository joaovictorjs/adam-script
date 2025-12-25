package ast

type Statement interface {
	Node
	statement()
}
