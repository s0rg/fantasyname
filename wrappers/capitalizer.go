package wrappers

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var caser = cases.Title(language.English)

type capitalizer struct {
	s fmt.Stringer
}

// Capitalized creates wrapper that capitalize (foo -> Foo) its input.
func Capitalized(s fmt.Stringer) fmt.Stringer {
	return &capitalizer{s: s}
}

func (cpt *capitalizer) String() string {
	return caser.String(cpt.s.String())
}
