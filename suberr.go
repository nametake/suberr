// Package suberr provides to add sub error to error.
//
// The suberr.Add function returns new erorr that added sub error to the original error.
// The sub error can be retrieved with suberr.SubCause function.
//
// This is supported to be used with github.com/pkg/errors.
// So, errors that added sub error implement the Cause method and also supported to errors.Cause.
// And, it dosen't lose stacktrace that given by github.com/pkg/errors
package suberr

import (
	"fmt"

	"github.com/pkg/errors"
)

// Add returns an error added sub error to the main error.
func Add(main, sub error) error {
	return &subError{
		main: main,
		sub:  sub,
	}
}

// WithMessage returns an error added sub error and message to the main error.
func WithMessage(main, sub error, msg string) error {
	err := &subError{
		main: main,
		sub:  sub,
	}
	if msg != "" {
		return errors.Wrap(err, msg)
	}
	return err
}

// SubCause returns the last added sub error.
// Even if errors are wrapped using errors.Wrap.
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
	if s.main == nil && s.sub == nil {
		return ""
	} else if s.main != nil && s.sub == nil {
		return s.main.Error()
	} else if s.main == nil && s.sub != nil {
		return s.sub.Error()
	}
	return fmt.Sprintf("%v: %v", s.sub, s.main)
}

func (s *subError) Cause() error {
	return s.main
}

func (s *subError) SubCause() error {
	return s.sub
}

func (s *subError) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		if st.Flag('+') {
			fmt.Fprintf(st, "%+v", s.main)
			return
		}
		fallthrough
	case 's':
		fmt.Fprintf(st, "%s", s.main)
	case 'q':
		fmt.Fprintf(st, "%q", s.main)
	}
}
