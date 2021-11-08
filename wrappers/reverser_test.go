package wrappers

import (
	"testing"

	"github.com/s0rg/fantasyname/stringers"
)

func TestReverser(t *testing.T) {
	t.Parallel()

	val := stringers.Literal("foo")
	rev := Reversed(val)

	if rev.String() != "oof" {
		t.Error("unexpected result")
	}
}
