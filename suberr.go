// Package suberr provides ...
package suberr

import (
	"fmt"

	"github.com/pkg/errors"
)

func Add(main, sub error) error {
	return &subError{
		main: main,
		sub:  sub,
	}
}

func WithMessage(main, sub error, msg string) error {
	return &subError{
		main: errors.Wrap(main, msg),
		sub:  sub,
	}
}

func SubCause(err error) error {
	for err != nil {
		subCause, ok := err.(subCauser)
		if ok {
			return subCause.SubCause()
		}
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return nil
}

type causer interface {
	Cause() error
}

type subCauser interface {
	SubCause() error
}

var (
	_ error     = (*subError)(nil)
	_ causer    = (*subError)(nil)
	_ subCauser = (*subError)(nil)
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

func (s *subError) SubCause() error {
	return s.sub
}
