package wrappers

import (
	"fmt"
)

type reverser struct {
	s fmt.Stringer
}

// Reversed creates wrapper, that reverses (foo -> oof) its input.
func Reversed(s fmt.Stringer) fmt.Stringer {
	return &reverser{s: s}
}

func (rv *reverser) String() string {
	return reverse(rv.s.String())
}

func reverse(s string) (rv string) {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}
