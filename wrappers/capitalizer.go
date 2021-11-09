package wrappers

import (
	"fmt"
	"strings"
)

type capitalizer struct {
	s fmt.Stringer
}

// Capitalized creates wrapper that capitalize (foo -> Foo) its input.
func Capitalized(s fmt.Stringer) fmt.Stringer {
	return &capitalizer{s: s}
}

func (cpt *capitalizer) String() string {
	return strings.Title(cpt.s.String())
}
