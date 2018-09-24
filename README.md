[![Build Status](https://travis-ci.org/mamal72/stringish.svg?branch=master)](https://travis-ci.org/mamal72/stringish)
[![Go Report Card](https://goreportcard.com/badge/github.com/mamal72/stringish)](https://goreportcard.com/report/github.com/mamal72/stringish)
[![GoDoc](https://godoc.org/github.com/mamal72/stringish?status.svg)](https://godoc.org/github.com/mamal72/stringish)
[![license](https://img.shields.io/github/license/mamal72/stringish.svg)](https://github.com/mamal72/stringish/blob/master/LICENSE)


# stringish

A set of chain-able string helpers to reduce the pain of manipulating strings.


## Installation

```bash
go get github.com/mamal72/stringish
# or use dep, vgo, glide or anything else
```


## Usage


### Implemented Methods


- `ReplaceN`
- `ReplaceAll`
- `ToLower`
- `ToUpper`
- `TrimPrefix`
- `TrimSuffix`
- `TrimSpaces`
- `TrimPrefixSpaces`
- `TrimSuffixSpaces`
- `HasPrefix`
- `HasSuffix`
- `Equals`
- `IsEmpty`
- `IsBlank`
- `Contains`
- `Len`
- `Index`
- `LastIndex`
- `Map`
- `Filter`
- `GetString`


### Example


```go
package main

import (
	"log"

	"github.com/mamal72/stringish"
)

func main() {
	str := stringish.New("  hello there ")
	str.TrimSpaces().Map(func(char string) string {
		return strings.ToUpper(char)
	})
	log.Println(str.GetString()) // "HELLO THERE"
	log.Println(str.Replace("HELLO", "bye").Contains("bye")) // true
	log.Println(str.GetString()) // "bye THERE"
}
```


## Ideas || Issues

Just create an issue and describe it. I'll check it ASAP!


## Contribution

You can fork the repository, improve or fix some part of it and then send the pull requests back (with some tests ofc) if you want to see them here. I really appreciate that. ❤️
