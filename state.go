package fantasyname

import (
	"fmt"

	"github.com/s0rg/fantasyname/stringers"
)

type wrapper func(fmt.Stringer) fmt.Stringer

type state struct {
	items     []fmt.Stringer
	wrappers  []wrapper
	rndfn     func(int) int
	splitpos  int
	IsLiteral bool
	IsGroup   bool
}

func (s *state) Add(v fmt.Stringer) {
	if !s.IsGroup {
		v = s.withWrappers(v)
	}

	s.items = append(s.items, v)
}

func (s *state) Wrap(w wrapper) {
	s.wrappers = append(s.wrappers, w)
}

func (s *state) Split() (rv *state) {
	cutoff := len(s.items) - s.splitpos

	if len(s.items) == 0 || cutoff == 0 {
		s.items = append(s.items, stringers.Empty{})
		s.splitpos++

		return s
	}

	switch cutoff {
	case 1:
		// optimization - do not create sequence for single stringer.
	default:
		tmp := make([]fmt.Stringer, cutoff)
		copy(tmp, s.items[s.splitpos:])

		s.items[s.splitpos] = stringers.Sequence(tmp)
	}

	s.splitpos++
	s.items = s.items[:s.splitpos]

	return s
}

func (s *state) Stringer() (rv fmt.Stringer) {
	switch {
	case len(s.items) == 1:
		rv = s.items[0]
	case s.IsGroup:
		rv = stringers.MakeRandom(s.items, s.rndfn)
	default:
		rv = stringers.Sequence(s.items)
	}

	return s.withWrappers(rv)
}

func (s *state) withWrappers(v fmt.Stringer) (rv fmt.Stringer) {
	rv = v

	if len(s.wrappers) > 0 {
		for i, w := range s.wrappers {
			rv = w(rv)
			s.wrappers[i] = nil
		}

		s.wrappers = s.wrappers[:0]
	}

	return rv
}
