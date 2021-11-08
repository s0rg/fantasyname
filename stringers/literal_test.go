package stringers

import "testing"

func TestLiteral(t *testing.T) {
	const val = "foo"

	t.Parallel()

	l := Literal(val)
	if l.String() != val {
		t.Error("unexpected result")
	}
}
