package wrappers

import (
	"fmt"
)

type collapser struct {
	s fmt.Stringer
}

// Collapsed creates wrapper that colapses similar runes in its input.
func Collapsed(s fmt.Stringer) fmt.Stringer {
	return &collapser{s: s}
}

func (clp *collapser) String() string {
	var (
		cur, prev rune
		rres      = []rune(clp.s.String())
		rout      = make([]rune, 0, len(rres))
	)

	for i := 0; i < len(rres); i++ {
		if cur = rres[i]; cur == prev && needCollapse(cur) {
			continue
		}

		rout = append(rout, cur)
		prev = cur
	}

	return string(rout)
}

func needCollapse(r rune) (yes bool) {
	switch r {
	case 'a', 'i', 'u', 'y':
		return true
	case 'h', 'j', 'q', 'v', 'w', 'x':
		return true
	}

	return false
}
