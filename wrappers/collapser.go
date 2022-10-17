package wrappers

import (
	"fmt"
)

type collapser struct {
	s fmt.Stringer
}

// Collapsed creates wrapper that colapses similar runes triples (and some doubles) in its input.
func Collapsed(s fmt.Stringer) fmt.Stringer {
	return &collapser{s: s}
}

func (cl *collapser) String() string {
	var (
		max, cnt  int
		cur, prev rune
		rres      = []rune(cl.s.String())
		rout      = make([]rune, 0, len(rres))
	)

	for i := 0; i < len(rres); i++ {
		if cur = rres[i]; cur == prev {
			cnt++
		} else {
			cnt = 0
		}

		max = 2

		if needCollapse(cur) {
			max = 1
		}

		if cnt < max {
			rout = append(rout, cur)
		}

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
