package wrappers

import (
	"fmt"
)

type collapser struct {
	s fmt.Stringer
}

func Collapsed(s fmt.Stringer) fmt.Stringer {
	return &collapser{s: s}
}

func (clp *collapser) String() string {
	var (
		oidx      int
		cur, prev rune
		rres      = []rune(clp.s.String())
		rout      = make([]rune, len(rres))
	)

	for i := 0; i < len(rres); i++ {
		if cur = rres[i]; cur == prev && needCollapse(cur) {
			continue
		}

		rout[oidx] = cur
		prev = cur
		oidx++
	}

	return string(rout)
}

func needCollapse(r rune) (yes bool) {
	switch r {
	case 'a', 'i', 'u', 'y':
	case 'h', 'j', 'q', 'v', 'w', 'x':
	default:
		return true
	}

	return false
}
