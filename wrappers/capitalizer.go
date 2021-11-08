package wrappers

import (
	"fmt"
	"strings"
)

type capitalizer struct {
	s fmt.Stringer
}

func Capitalized(s fmt.Stringer) fmt.Stringer {
	return &capitalizer{s: s}
}

func (cpt *capitalizer) String() string {
	return strings.Title(cpt.s.String())
}
