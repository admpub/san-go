package san

import (
	"reflect"
	"testing"
)

func testLex(t *testing.T, input string, expectedTokens []Token) {
	Tokens := Lex(input)
	if !reflect.DeepEqual(Tokens, expectedTokens) {
		t.Fatalf("Different Token. Expected: %#v\nGot: %#v\n", expectedTokens, Tokens)
	}
}

func TestSimpleString(t *testing.T) {
	testLex(t, `str = "myString"`, []Token{
		{Position{1, 1}, TokenKey, "str"},
		{Position{1, 5}, TokenEquals, "="},
		{Position{1, 8}, TokenString, "myString"},
		{Position{1, 17}, TokenEOF, ""},
	})
}

func TestSimpleBoolean(t *testing.T) {
	testLex(t, "foo = false", []Token{
		{Position{1, 1}, TokenKey, "foo"},
		{Position{1, 5}, TokenEquals, "="},
		{Position{1, 7}, TokenBoolean, "false"},
		{Position{1, 12}, TokenEOF, ""},
	})
}

func TestNestedStringList(t *testing.T) {
	testLex(t, `a = [ ["hello", "world"] ]`, []Token{
		{Position{1, 1}, TokenKey, "a"},
		{Position{1, 3}, TokenEquals, "="},
		{Position{1, 5}, TokenLeftBracket, "["},
		{Position{1, 7}, TokenLeftBracket, "["},
		{Position{1, 9}, TokenString, "hello"},
		{Position{1, 15}, TokenComma, ","},
		{Position{1, 18}, TokenString, "world"},
		{Position{1, 24}, TokenRightBracket, "]"},
		{Position{1, 26}, TokenRightBracket, "]"},
		{Position{1, 27}, TokenEOF, ""},
	})
}

func TestNestedIntList(t *testing.T) {
	testLex(t, "a = [ [42, 21], [10] ]", []Token{
		{Position{1, 1}, TokenKey, "a"},
		{Position{1, 3}, TokenEquals, "="},
		{Position{1, 5}, TokenLeftBracket, "["},
		{Position{1, 7}, TokenLeftBracket, "["},
		{Position{1, 8}, TokenInteger, "42"},
		{Position{1, 10}, TokenComma, ","},
		{Position{1, 12}, TokenInteger, "21"},
		{Position{1, 14}, TokenRightBracket, "]"},
		{Position{1, 15}, TokenComma, ","},
		{Position{1, 17}, TokenLeftBracket, "["},
		{Position{1, 18}, TokenInteger, "10"},
		{Position{1, 20}, TokenRightBracket, "]"},
		{Position{1, 22}, TokenRightBracket, "]"},
		{Position{1, 23}, TokenEOF, ""},
	})
}

func TestIntList(t *testing.T) {
	testLex(t, "a = [ 42, 21, 10, ]", []Token{
		{Position{1, 1}, TokenKey, "a"},
		{Position{1, 3}, TokenEquals, "="},
		{Position{1, 5}, TokenLeftBracket, "["},
		{Position{1, 7}, TokenInteger, "42"},
		{Position{1, 9}, TokenComma, ","},
		{Position{1, 11}, TokenInteger, "21"},
		{Position{1, 13}, TokenComma, ","},
		{Position{1, 15}, TokenInteger, "10"},
		{Position{1, 17}, TokenComma, ","},
		{Position{1, 19}, TokenRightBracket, "]"},
		{Position{1, 20}, TokenEOF, ""},
	})
}
