package suberr_test

import (
	"fmt"

	"github.com/nametake/suberr"
	"github.com/pkg/errors"
)

func ExampleSubCause() {
	err := errors.New("cause")

	err = errors.Wrap(err, "first")
	err = suberr.Add(err, errors.New("sub"))
	err = errors.Wrap(err, "second")
	err = errors.Wrap(err, "third")

	fmt.Println(err)
	fmt.Println(suberr.SubCause(err))
	fmt.Println(errors.Cause(err))

	// Output:
	// third: second: sub: first: cause
	// sub
	// cause
}
