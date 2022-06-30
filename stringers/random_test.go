package stringers

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRandom(t *testing.T) {
	t.Parallel()

	val := MakeRandom(
		[]fmt.Stringer{Literal("a"), Literal("b"), Literal("c")},
		rand.Intn,
	)

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
