package san

import (
	"testing"
)

type _testTokenString struct {
	token tokenType
	str   string
}

func TestTokenString(t *testing.T) {
	toTest := []_testTokenString{
		{tokenError, "Error"},
		{tokenEOF, "EOF"},
		{tokenComma, "Comma"},
		{tokenHash, "Hash"},
		{tokenEquals, "Equals"},
		{tokenLeftBracket, "LeftBracket"},
		{tokenRightBracket, "RightBracket"},
		{tokenLeftBrace, "LeftBrace"},
		{tokenRightBrace, "RightBrace"},
		{tokenSignleQuote, "SingleQuote"},
		{tokenDoubleQuote, "DoubleQuote"},
		{tokenKey, "Key"},
		{tokenBoolean, "Boolean"},
		{tokenInteger, "Integer"},
		{tokenFloat, "Float"},
		{tokenString, "String"},
	}

	for _, pair := range toTest {
		str := pair.token.String()
		if str != pair.str {
			t.Errorf("Expected %s, got %s instead", pair.str, str)
		}
	}
}
