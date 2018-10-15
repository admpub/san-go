# SAN-go

[![GoDoc](https://godoc.org/github.com/bloom42/san-go?status.svg)](https://godoc.org/github.com/bloom42/san-go)
[![Build Status](https://travis-ci.org/bloom42/san-go.svg?branch=master)](https://travis-ci.org/bloom42/san-go)
[![GitHub release](https://img.shields.io/github/release/bloom42/san-go.svg)](https://github.com/bloom42/san-go/releases/latest)

SAN (pronounce `/seɪn/`, like sane) CLI and parser for Go.

Spec: [https://astrocorp.net/san](https://astrocorp.net/san/)

Compatible with SAN version: [v1.0.0](https://astrocorp.net/san/versions/v1.0.0/)

1. [Installation](#installation)
2. [Library](#library)
3. [CLI](#cli)

-------------------


## Installation

```bash
go get -u github.com/bloom42/san-go/...
```



## Library

```go
package main

import (
	"fmt"

	"github.com/bloom42/san-go"
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
	var s S

	err := san.Unmarshal([]byte(str1), &s)
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
```

```bash
go run main.go
main.S{A:"a", B:[]int64{1, 2}, C:main.C{A:1, D:[]main.D{main.D{A:"3.3"}, main.D{A:"xxx"}}}}

a = "a"
b = [
  1,
  2,
]
c = {
  a = 1

  d = [
    {
      A = "3.3"
    },
    {
      A = "xxx"
    },
  ]
}
```



## CLI


This repo also contains a CLI helper for the SAN format. It can be installed with the following command:
```bash
$ go get -u github.com/bloom42/san-go/...
```

### Examples

Convert a [.toml, .json, .yml, .yaml] file to a .san
```bash
$ san convert ../config.yml # wil create ../config.san
```

Automatically formats a SAN file
```bash
$ san fmt config.san
```

Check a `.san` file validity
```bash
$ san validate config.san
```
