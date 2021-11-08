package wrappers

import (
	"testing"

	"github.com/s0rg/fantasyname/stringers"
)

func TestCapitalizer(t *testing.T) {
	t.Parallel()

	val := stringers.Literal("foo")
	cpt := Capitalized(val)

	if cpt.String() != "Foo" {
		t.Error("unexpected result")
	}
}
