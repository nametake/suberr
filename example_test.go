package suberr_test

import (
	"fmt"

	"github.com/nametake/suberr"
	"github.com/pkg/errors"
)

var errSub = errors.New("sub error")

func cause() error {
	return errors.New("cause")
}

func first() error {
	err := cause()
	err = suberr.Add(err, errSub)
	return errors.Wrap(err, "first")

	// one line:
	// return suberr.WithMessage(err, errSub, "first")
}

func second() error {
	err := first()
	return errors.Wrap(err, "second")
}

func Example() {
	err := second()

	fmt.Println(err)
	fmt.Println(suberr.SubCause(err))
	fmt.Println(errors.Cause(err))

	// Output:
	// second: first: sub error: cause
	// sub error
	// cause
}
