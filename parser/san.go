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

func (t *Tree) SetPath(keys []string, value interface{}) {
	subtree := t
	for _, intermediateKey := range keys[:len(keys)-1] {
		nextTree, exists := subtree.Values[intermediateKey]
		if !exists {
			nextTree = NewTree()
			subtree.Values[intermediateKey] = nextTree // add new element here
		}
		switch node := nextTree.(type) {
		case *Tree:
			subtree = node
		case []*Tree:
			// go to most recent element
			if len(node) == 0 {
				// create element if it does not exist
				subtree.Values[intermediateKey] = append(node, NewTree())
			}
			subtree = node[len(node)-1]
		}
	}

	var toInsert interface{}

	switch value.(type) {
	case *Tree:
		//tt := value.(*Tree)
		toInsert = value
	case []*Tree:
		toInsert = value
	case *Value:
		tt := value.(*Value)
		toInsert = tt
	default:
		toInsert = &Value{Val: value}
	}

	subtree.Values[keys[len(keys)-1]] = toInsert
}
