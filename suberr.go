// Package suberr provides ...
package suberr

import (
	"fmt"
)

func Add(main, sub error) error {
	return &subError{
		main: main,
		sub:  sub,
	}
}

type causer interface {
	Cause() error
}

type getter interface {
	Get() error
}

var (
	_ error = (*subError)(nil)
	// _ fmt.Formatter = (*subError)(nil)
	_ causer = (*subError)(nil)
	// _ getter = (*subError)(nil)
)

type subError struct {
	main, sub error
}

func (s *subError) Error() string {
	return fmt.Sprintf("%v: %v", s.main, s.sub)
}

func (s *subError) Cause() error {
	return s.main
}
