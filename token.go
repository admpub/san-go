package san

import (
	"fmt"
)

// Define tokens
type tokenType int

const (
	tokenError tokenType = iota
	tokenEOF
	tokenComma
	tokenHash
	tokenEquals
	tokenLeftBracket
	tokenRightBracket
	tokenLeftBrace
	tokenRightBrace
	tokenSignleQuote
	tokenDoubleQuote
	tokenKey
	tokenBoolean
	tokenInteger
	tokenFloat
	tokenString
)

const (
	eof          = 0
	comma        = ','
	hash         = '#'
	equals       = '='
	leftBracket  = '{'
	rightBracket = '}'
	leftBrace    = '['
	rightBrace   = ']'
	singleQuote  = '\''
	doubleQuote  = '"'
)

type token struct {
	Position
	Type  tokenType
	Value string
}

func (t tokenType) String() string {
	switch t {
	case tokenError:
		return "Error"
	case tokenEOF:
		return "EOF"
	case tokenComma:
		return "Comma"
	case tokenHash:
		return "Hash"
	case tokenEquals:
		return "Equals"
	case tokenLeftBracket:
		return "LeftBracket"
	case tokenRightBracket:
		return "RightBracket"
	case tokenLeftBrace:
		return "LeftBrace"
	case tokenRightBrace:
		return "RightBrace"
	case tokenSignleQuote:
		return "SingleQuote"
	case tokenDoubleQuote:
		return "DoubleQuote"
	case tokenKey:
		return "Key"
	case tokenBoolean:
		return "Boolean"
	case tokenInteger:
		return "Integer"
	case tokenFloat:
		return "Float"
	case tokenString:
		return "String"
	}
	return fmt.Sprintf("BUG: Unknown token type '%d'.", int(t))
}
