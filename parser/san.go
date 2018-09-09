package parser

import (
	"github.com/phasersec/san-go/lexer"
)

// Tree is the result of the parsing of a SAN file.
type Tree struct {
	Values   map[string]interface{} // string -> *tomlValue, *Tree, []*Tree, comments
	Position lexer.Position
}

// Value represents a SAN value
type Value struct {
	Val      interface{} // string, int64, uint64, float64, bool or [] of any
	Position lexer.Position
}

// NewTree initialize a new Tree
func NewTree() *Tree {
	return &Tree{
		Values:   make(map[string]interface{}),
		Position: lexer.Position{Line: 1, Col: 1},
	}
}
