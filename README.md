# suberr

[![CircleCI](https://circleci.com/gh/nametake/suberr.svg?style=svg)](https://circleci.com/gh/nametake/suberr)
[![GoDoc](https://godoc.org/github.com/nametake/suberr?status.svg)](https://godoc.org/github.com/nametake/suberr)

Package suberr provides to add sub error to error.

## Install

`go get github.com/nametake/suberr`

## Usage

```go
package main

import (
	"fmt"

	"github.com/nametake/suberr"
	"github.com/pkg/errors"
)

func main() {
	err := errors.New("cause")

	err = errors.Wrap(err, "first")
	err = suberr.Add(err, errors.New("sub"))
	err = errors.Wrap(err, "second")
	err = errors.Wrap(err, "third")

	fmt.Println(err)                  // third: second: sub: first: cause
	fmt.Println(suberr.SubCause(err)) // sub
	fmt.Println(errors.Cause(err))    // main
}
```
