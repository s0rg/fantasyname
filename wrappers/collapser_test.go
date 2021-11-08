package wrappers

import (
	"testing"

	"github.com/s0rg/fantasyname/stringers"
)

func TestCollapser(t *testing.T) {
	t.Parallel()

	val := stringers.Literal("foo")
	col := Collapsed(val)

	if col.String() != "foo" {
		t.Error("1: unexpected result", col.String())
	}

	val = stringers.Literal("fiiqqoo")
	col = Collapsed(val)

	if col.String() != "fiqoo" {
		t.Error("2: unexpected result", col.String())
	}
}
