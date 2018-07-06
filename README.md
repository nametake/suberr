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

func first() error {
	err := errors.New("cause")
	return errors.Wrap(err, "first")
}

func second() error {
	err := first()
	err = suberr.Add(err, errors.New("sub"))
	return errors.Wrap(err, "second")
}

func third() error {
	err := second()
	return errors.Wrap(err, "third")
}

func main() {
	err := third()
	fmt.Println(err)                  // third: second: sub: first: cause
	fmt.Println(suberr.SubCause(err)) // sub
	fmt.Println(errors.Cause(err))    // main
}
```
