package lexer

import (
	"fmt"
)

const (
	TokenError TokenType = iota
	TokenEOF
	TokenComma
	TokenHash
	TokenEquals
	TokenLeftBracket
	TokenRightBracket
	TokenLeftBrace
	TokenRightBrace
	TokenKey
	TokenBoolean
	TokenInteger
	TokenFloat
	TokenString
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

// TokenType represents all the possible values of a Token
type TokenType int

// Token is a token emited by the Lexer
type Token struct {
	Position
	Type  TokenType
	Value string
}

func (t TokenType) String() string {
	switch t {
	case TokenError:
		return "Error"
	case TokenEOF:
		return "EOF"
	case TokenComma:
		return "Comma"
	case TokenHash:
		return "Hash"
	case TokenEquals:
		return "Equals"
	case TokenLeftBracket:
		return "LeftBracket"
	case TokenRightBracket:
		return "RightBracket"
	case TokenLeftBrace:
		return "LeftBrace"
	case TokenRightBrace:
		return "RightBrace"
	case TokenKey:
		return "Key"
	case TokenBoolean:
		return "Boolean"
	case TokenInteger:
		return "Integer"
	case TokenFloat:
		return "Float"
	case TokenString:
		return "String"
	}
	return fmt.Sprintf("BUG: Unknown token type '%d'.", int(t))
}

func (t Token) String() string {
	switch t.Type {
	case TokenEOF:
		return "EOF"
	case TokenError:
		return t.Value
	}

	return fmt.Sprintf("%q", t.Value)
}
