package wrappers

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type capitalizer struct {
	s fmt.Stringer
}

// Capitalized creates wrapper that capitalize (foo -> Foo) its input.
func Capitalized(s fmt.Stringer) fmt.Stringer {
	return &capitalizer{s: s}
}

func (cp *capitalizer) String() string {
	return cases.Title(language.English).String(cp.s.String())
}
