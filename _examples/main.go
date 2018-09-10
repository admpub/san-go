package main

import (
	"fmt"

	"github.com/phasersec/san-go"
	"github.com/phasersec/san-go/parser"
)

type S struct {
	A string  `san:"a"`
	B []int64 `san:"b"`
}

func main() {
	str1 := `#hello
a = 3.3
b = { a = 5 }
c = [1,2]`
	tree, err := parser.Parse([]byte(str1))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", tree)

	str2 := `
a = "a"
b = [1, 2]
`
	var s S

	err = san.Unmarshal([]byte(str2), &s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", s)
}
