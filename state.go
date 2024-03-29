package fantasyname

import (
	"fmt"

	"github.com/s0rg/fantasyname/stringers"
)

type wrapper func(fmt.Stringer) fmt.Stringer

type state struct {
	rand      func(int) int
	items     []fmt.Stringer
	wrappers  []wrapper
	splitpos  int
	IsLiteral bool
	IsGroup   bool
}

func (s *state) Add(v fmt.Stringer) {
	if !s.IsGroup {
		v = s.applyWrappers(v)
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
		rv = stringers.MakeRandom(s.items, s.rand)
	default:
		rv = stringers.Sequence(s.items)
	}

	return s.applyWrappers(rv)
}

func (s *state) applyWrappers(v fmt.Stringer) (rv fmt.Stringer) {
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
