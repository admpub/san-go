package lexer

import (
	"reflect"
	"testing"
)

func testLex(t *testing.T, input string, expectedTokens []Token) {
	Tokens := Lex(input)

	if len(expectedTokens) != len(Tokens) {
		t.Fatalf(
			"Not enought Token lexed. Expected: %d\nGot: %d (%#v)\n",
			len(expectedTokens),
			len(Tokens),
			Tokens,
		)
	}

	for i := range expectedTokens {

		if reflect.DeepEqual(Tokens[i], expectedTokens[i]) != true {
			t.Errorf("Different Token.\nExpected: %#v\nGot     : %#v\n", expectedTokens[i], Tokens[i])
		}
	}
}

func TestSimpleComment(t *testing.T) {
	testLex(t, "# blahblah", []Token{
		{Position{1, 1}, TokenComment, "blahblah"},
		{Position{1, 11}, TokenEOF, ""},
	})
}

func TestSimpleString(t *testing.T) {
	testLex(t, `str = "myString"`, []Token{
		{Position{1, 1}, TokenKey, "str"},
		{Position{1, 5}, TokenEquals, "="},
		{Position{1, 7}, TokenString, "myString"},
		{Position{1, 17}, TokenEOF, ""},
	})
}

func TestFloatWithExponent1(t *testing.T) {
	testLex(t, "a = 5e+22", []Token{
		{Position{1, 1}, TokenKey, "a"},
		{Position{1, 3}, TokenEquals, "="},
		{Position{1, 5}, TokenFloat, "5e+22"},
		{Position{1, 10}, TokenEOF, ""},
	})
}

func TestFloatWithExponent2(t *testing.T) {
	testLex(t, "a = 6.69e-22", []Token{
		{Position{1, 1}, TokenKey, "a"},
		{Position{1, 3}, TokenEquals, "="},
		{Position{1, 5}, TokenFloat, "6.69e-22"},
		{Position{1, 13}, TokenEOF, ""},
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
		{Position{1, 8}, TokenString, "hello"},
		{Position{1, 15}, TokenComma, ","},
		{Position{1, 17}, TokenString, "world"},
		{Position{1, 24}, TokenRightBracket, "]"},
		{Position{1, 26}, TokenRightBracket, "]"},
		{Position{1, 27}, TokenEOF, ""},
	})
}

func TestNestedIntegerList(t *testing.T) {
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

func TestIntegerList(t *testing.T) {
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

func TestBooleanList(t *testing.T) {
	testLex(t, "foo = [true, false, true]", []Token{
		{Position{1, 1}, TokenKey, "foo"},
		{Position{1, 5}, TokenEquals, "="},
		{Position{1, 7}, TokenLeftBracket, "["},
		{Position{1, 8}, TokenBoolean, "true"},
		{Position{1, 12}, TokenComma, ","},
		{Position{1, 14}, TokenBoolean, "false"},
		{Position{1, 19}, TokenComma, ","},
		{Position{1, 21}, TokenBoolean, "true"},
		{Position{1, 25}, TokenRightBracket, "]"},
		{Position{1, 26}, TokenEOF, ""},
	})
}

func TestMultiString(t *testing.T) {
	str := `a = "tesla"
b = "spacex"
`
	testLex(t, str, []Token{
		{Position{1, 1}, TokenKey, "a"},
		{Position{1, 3}, TokenEquals, "="},
		{Position{1, 5}, TokenString, "tesla"},
		{Position{2, 1}, TokenKey, "b"},
		{Position{2, 3}, TokenEquals, "="},
		{Position{2, 5}, TokenString, "spacex"},
		{Position{3, 1}, TokenEOF, ""},
	})
}

func TestMultiInteger(t *testing.T) {
	testLex(t, "foo = 42\nbar=21", []Token{
		{Position{1, 1}, TokenKey, "foo"},
		{Position{1, 5}, TokenEquals, "="},
		{Position{1, 7}, TokenInteger, "42"},
		{Position{2, 1}, TokenKey, "bar"},
		{Position{2, 4}, TokenEquals, "="},
		{Position{2, 5}, TokenInteger, "21"},
		{Position{2, 7}, TokenEOF, ""},
	})
}

func TestNestedLists(t *testing.T) {
	testLex(t, "foo = [[[]]]", []Token{
		{Position{1, 1}, TokenKey, "foo"},
		{Position{1, 5}, TokenEquals, "="},
		{Position{1, 7}, TokenLeftBracket, "["},
		{Position{1, 8}, TokenLeftBracket, "["},
		{Position{1, 9}, TokenLeftBracket, "["},
		{Position{1, 10}, TokenRightBracket, "]"},
		{Position{1, 11}, TokenRightBracket, "]"},
		{Position{1, 12}, TokenRightBracket, "]"},
		{Position{1, 13}, TokenEOF, ""},
	})
}

func TestSimpleMultilineString(t *testing.T) {
	testLex(t, `foo = """hello "literal" world"""`, []Token{
		{Position{1, 1}, TokenKey, "foo"},
		{Position{1, 5}, TokenEquals, "="},
		{Position{1, 7}, TokenString, `hello "literal" world`},
		{Position{1, 34}, TokenEOF, ""},
	})
}

func TestLiteralString(t *testing.T) {
	testLex(t, `foo = 'C:\Users\nodejs\templates'`, []Token{
		{Position{1, 1}, TokenKey, "foo"},
		{Position{1, 5}, TokenEquals, "="},
		{Position{1, 7}, TokenString, `C:\Users\nodejs\templates`},
		{Position{1, 34}, TokenEOF, ""},
	})
	testLex(t, `foo = '\\ServerX\admin$\system32\'`, []Token{
		{Position{1, 1}, TokenKey, "foo"},
		{Position{1, 5}, TokenEquals, "="},
		{Position{1, 7}, TokenString, `\\ServerX\admin$\system32\`},
		{Position{1, 35}, TokenEOF, ""},
	})
	testLex(t, `foo = 'Tom "Dubs" Preston-Werner'`, []Token{
		{Position{1, 1}, TokenKey, "foo"},
		{Position{1, 5}, TokenEquals, "="},
		{Position{1, 7}, TokenString, `Tom "Dubs" Preston-Werner`},
		{Position{1, 34}, TokenEOF, ""},
	})
	testLex(t, `foo = '<\i\c*\s*>'`, []Token{
		{Position{1, 1}, TokenKey, "foo"},
		{Position{1, 5}, TokenEquals, "="},
		{Position{1, 7}, TokenString, `<\i\c*\s*>`},
		{Position{1, 19}, TokenEOF, ""},
	})
	testLex(t, `foo = 'C:\Users\nodejs\unfinis`, []Token{
		{Position{1, 1}, TokenKey, "foo"},
		{Position{1, 5}, TokenEquals, "="},
		{Position{1, 7}, TokenError, "lexer: unclosed string"},
	})
}
