package suberr

import (
	"testing"

	"github.com/pkg/errors"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		main error
		sub  error
		want string
	}{
		{
			name: "simple",
			main: errors.New("main error"),
			sub:  errors.New("sub error"),
			want: "sub error: main error",
		},
		{
			name: "nil",
			main: nil,
			sub:  nil,
			want: "",
		},
		{
			name: "nil main error",
			main: nil,
			sub:  errors.New("sub error"),
			want: "sub error",
		},
		{
			name: "nil sub error",
			main: errors.New("main error"),
			sub:  nil,
			want: "main error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.main, tt.sub).Error(); got != tt.want {
				t.Errorf("Add.Error(): got: %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithMessage(t *testing.T) {
	tests := []struct {
		name string
		main error
		sub  error
		msg  string
		want string
	}{
		{
			name: "simple",
			main: errors.New("main error"),
			sub:  errors.New("sub error"),
			msg:  "message",
			want: "message: sub error: main error",
		},
		{
			name: "simple and no message",
			main: errors.New("main error"),
			sub:  errors.New("sub error"),
			msg:  "",
			want: "sub error: main error",
		},
		{
			name: "nil and no msg",
			main: nil,
			sub:  nil,
			msg:  "",
			want: "",
		},
		{
			name: "nil main error and message",
			main: nil,
			sub:  errors.New("sub error"),
			msg:  "message",
			want: "message: sub error",
		},
		{
			name: "nil main error and no message",
			main: nil,
			sub:  errors.New("sub error"),
			msg:  "",
			want: "sub error",
		},
		{
			name: "nil sub error and message",
			main: errors.New("main error"),
			sub:  nil,
			msg:  "message",
			want: "message: main error",
		},
		{
			name: "nil sub error and no message",
			main: errors.New("main error"),
			sub:  nil,
			msg:  "",
			want: "main error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithMessage(tt.main, tt.sub, tt.msg).Error(); got != tt.want {
				t.Errorf("WithMessage.Error(): got: %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubCause(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "simple",
			err:  Add(errors.New("main error"), errors.New("sub error")),
			want: "sub error",
		},
		{
			name: "nil main error",
			err:  Add(nil, errors.New("sub error")),
			want: "sub error",
		},
		{
			name: "use errors.Wrap outermost",
			err:  errors.Wrap(Add(errors.New("main error"), errors.New("sub error")), "outer"),
			want: "sub error",
		},
		{
			name: "wraped main error",
			err:  Add(errors.Wrap(errors.New("main error"), "wrap main"), errors.New("sub error")),
			want: "sub error",
		},
		{
			name: "use errors.Wrap outermost and wraped main error",
			err:  errors.Wrap(Add(errors.Wrap(errors.New("main error"), "wrap main"), errors.New("sub error")), "outer"),
			want: "sub error",
		},
		{
			name: "use multiple suberr.Add",
			err:  Add(Add(errors.New("main error"), errors.New("inner sub error")), errors.New("outer sub error")),
			want: "outer sub error",
		},
		{
			name: "use WithMessage",
			err:  WithMessage(errors.New("main error"), errors.New("sub error"), "msg"),
			want: "sub error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubCause(tt.err); got.Error() != tt.want {
				t.Errorf("SubCause(): got: %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubCauseNil(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{
			name: "nil",
			err:  nil,
		},
		{
			name: "all nil",
			err:  Add(nil, nil),
		},
		{
			name: "nil sub error",
			err:  Add(errors.New("main error"), nil),
		},
		{
			name: "use multiple suberr.Add and nil outer sub error",
			err:  Add(Add(errors.New("main error"), errors.New("inner sub error")), nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubCause(tt.err); got != nil {
				t.Errorf("SubCuase(): got: %s, want nil", got)
			}
		})
	}
}

func TestCause(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "simple",
			err:  Add(errors.New("main error"), errors.New("sub error")),
			want: "main error",
		},
		{
			name: "nil sub error",
			err:  Add(errors.New("main error"), nil),
			want: "main error",
		},
		{
			name: "use errors.Wrap outermost",
			err:  errors.Wrap(Add(errors.New("main error"), errors.New("sub error")), "outer"),
			want: "main error",
		},
		{
			name: "wraped main error",
			err:  Add(errors.Wrap(errors.New("main error"), "wrap main"), errors.New("sub error")),
			want: "main error",
		},
		{
			name: "use errors.Wrap outermost and wraped main error",
			err:  errors.Wrap(Add(errors.Wrap(errors.New("main error"), "wrap main"), errors.New("sub error")), "outer"),
			want: "main error",
		},
		{
			name: "use multiple suberr.Add",
			err:  Add(Add(errors.New("main error"), errors.New("inner sub error")), errors.New("outer sub error")),
			want: "main error",
		},
		{
			name: "use multiple suberr.Add and nil outer sub error",
			err:  Add(Add(errors.New("main error"), errors.New("inner sub error")), nil),
			want: "main error",
		},
		{
			name: "use WithMessage",
			err:  WithMessage(errors.New("main error"), errors.New("sub error"), "msg"),
			want: "main error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errors.Cause(tt.err); got.Error() != tt.want {
				t.Errorf("errors.Cause(): got: %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCauseNil(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{
			name: "all nil",
			err:  Add(nil, nil),
		},
		{
			name: "nil main error",
			err:  Add(nil, errors.New("sub error")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errors.Cause(tt.err); got != nil {
				t.Errorf("errors.Cause(): got: %v, want nil ", got)
			}
		})
	}
}
