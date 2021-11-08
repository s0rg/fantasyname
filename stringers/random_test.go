package stringers

import "testing"

func TestRandom(t *testing.T) {
	t.Parallel()

	val := Random{
		Literal("a"),
		Literal("b"),
		Literal("c"),
	}

	var hit, miss int

	for i := 0; i < 5; i++ {
		if val.String() == "a" {
			hit++
		} else {
			miss++
		}
	}

	if hit == 0 || miss == 0 {
		t.Error("unexpected result")
	}
}
