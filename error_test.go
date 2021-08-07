package httperror_test

import (
	"fmt"
	"testing"

	"github.com/shurcooL/httperror"
)

func TestIsRedirect(t *testing.T) {
	for _, tt := range [...]struct {
		name string
		in   error
		want bool
	}{
		{name: "bare", in: httperror.Redirect{URL: "u"}, want: true},
		{name: "wrapped in %w", in: fmt.Errorf("desc: %w", httperror.Redirect{URL: "u"}), want: true},
		{name: "wrapped in %v", in: fmt.Errorf("desc: %v", httperror.Redirect{URL: "u"}), want: false},
		{name: "not redirect", in: httperror.Method{Allowed: []string{"GET"}}, want: false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, got := httperror.IsRedirect(tt.in)
			if got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
