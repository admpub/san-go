package main

import (
	"fmt"

	"github.com/phasersec/san-go"
	"github.com/phasersec/san-go/parser"
)

type D struct {
	A int64
}

type C struct {
	A int64 `san:"a"`
	D D
}

type S struct {
	A string  `san:"a"`
	B []int64 `san:"b"`
	C C       `san:"c"`
}

func main() {
	str1 := `
a = "a"
b = [1, 2]
c = { a = 1, d = { a = 3 } }
`
	tree, err := parser.Parse([]byte(str1))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n\n", tree.Values["c"].(*parser.Value).Val)

	var s S

	err = san.Unmarshal([]byte(str1), &s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", s)
}
