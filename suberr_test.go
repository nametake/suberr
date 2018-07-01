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
		wantSub  error
		wantMain error
	}{
		{
			name:     "all nil",
			arg:      Add(nil, nil),
			wantSub:  nil,
			wantMain: nil,
		},
		{
			name: "one level",
			arg: Add(
				errors.New("main error"),
				errors.New("sub error"),
			),
			wantSub:  errors.New("sub error"),
			wantMain: errors.New("main error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(SubCause(tt.arg), tt.wantSub, errCmp); diff != "" {
				t.Errorf("failed suberr.SubCause: %s\n", diff)
			}
			if diff := cmp.Diff(errors.Cause(tt.arg), tt.wantMain, errCmp); diff != "" {
				t.Error("failed errors.Cause: %s\n", diff)
			}
		})
	}
}
