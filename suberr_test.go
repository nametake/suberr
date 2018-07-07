package suberr

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
)

var (
	errCmp = cmp.Comparer(func(x, y error) bool {
		return x.Error() == y.Error()
	})
)

func TestSubCause(t *testing.T) {
	tests := []struct {
		name     string
		arg      error
		wantStr  string
		wantSub  error
		wantMain error
	}{
		{
			name:     "all nil",
			arg:      Add(nil, nil),
			wantStr:  "",
			wantSub:  nil,
			wantMain: nil,
		},
		{
			name:     "one level",
			arg:      Add(errors.New("main error"), errors.New("sub error")),
			wantStr:  "sub error: main error",
			wantSub:  errors.New("sub error"),
			wantMain: errors.New("main error"),
		},
		{
			name:     "nil main error",
			arg:      Add(nil, errors.New("sub error")),
			wantStr:  "sub error",
			wantSub:  errors.New("sub error"),
			wantMain: nil,
		},
		{
			name:     "nil sub error",
			arg:      Add(errors.New("main error"), nil),
			wantStr:  "main error",
			wantSub:  nil,
			wantMain: errors.New("main error"),
		},
		{
			name:     "use errors.Wrap outermost",
			arg:      errors.Wrap(Add(errors.New("main error"), errors.New("sub error")), "outer"),
			wantStr:  "outer: sub error: main error",
			wantSub:  errors.New("sub error"),
			wantMain: errors.New("main error"),
		},
		{
			name:     "wraped main error",
			arg:      Add(errors.Wrap(errors.New("main error"), "wrap main"), errors.New("sub error")),
			wantStr:  "sub error: wrap main: main error",
			wantSub:  errors.New("sub error"),
			wantMain: errors.New("main error"),
		},
		{
			name:     "use errors.Wrap outermost and wraped main error",
			arg:      errors.Wrap(Add(errors.Wrap(errors.New("main error"), "wrap main"), errors.New("sub error")), "outer"),
			wantStr:  "outer: sub error: wrap main: main error",
			wantSub:  errors.New("sub error"),
			wantMain: errors.New("main error"),
		},
		{
			name:     "use multiple suberr.Add",
			arg:      Add(Add(errors.New("main error"), errors.New("inner sub error")), errors.New("outer sub error")),
			wantStr:  "outer sub error: inner sub error: main error",
			wantSub:  errors.New("outer sub error"),
			wantMain: errors.New("main error"),
		},
		{
			name:     "use multiple suberr.Add and nil outer sub error",
			arg:      Add(Add(errors.New("main error"), errors.New("inner sub error")), nil),
			wantStr:  "inner sub error: main error",
			wantSub:  nil,
			wantMain: errors.New("main error"),
		},
		{
			name:     "use WithMessage",
			arg:      WithMessage(errors.New("main error"), errors.New("sub error"), "msg"),
			wantStr:  "msg: sub error: main error",
			wantSub:  errors.New("sub error"),
			wantMain: errors.New("main error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.arg.Error(), tt.wantStr); diff != "" {
				t.Errorf("not equals Error() : %s\n", diff)
			}
			if diff := cmp.Diff(SubCause(tt.arg), tt.wantSub, errCmp); diff != "" {
				t.Errorf("failed suberr.SubCause: %s\n", diff)
			}
			if diff := cmp.Diff(errors.Cause(tt.arg), tt.wantMain, errCmp); diff != "" {
				t.Errorf("failed errors.Cause: %s\n", diff)
			}
		})
	}
}
