package main

import (
	"fmt"

	"github.com/bloom42/san-go"
	"github.com/bloom42/san-go/parser"
)

type D struct {
	A string
}

type C struct {
	A int64 `san:"a"`
	D []D   `san:"d"`
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
c = { a = 1, d = [ { a = "3.3" }, { a = "xxx" } ] }
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
	fmt.Printf("%#v\n\n", s)

	b, err := san.Marshal(s)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(b))
}
