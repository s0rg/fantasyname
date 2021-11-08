package stringers

import "testing"

func TestSequence(t *testing.T) {
	t.Parallel()

	val := Sequence{
		Literal("a"),
		Literal("b"),
		Literal("c"),
	}

	if val.String() != "abc" {
		t.Error("unexpected result")
	}
}
