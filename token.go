package san

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
	tokenDubleQuote
	tokenBool
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
