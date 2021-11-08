package fantasyname

import (
	"fmt"

	"github.com/s0rg/fantasyname/stringers"
)

type wrapper func(fmt.Stringer) fmt.Stringer

type state struct {
	items     []fmt.Stringer
	wrappers  []wrapper
	splitpos  int
	IsLiteral bool
	IsGroup   bool
}

func (s *state) Wrap(w wrapper) {
	s.wrappers = append(s.wrappers, w)
}

func (s *state) wrapStringer(v fmt.Stringer) (rv fmt.Stringer) {
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

func (s *state) Add(v fmt.Stringer) {
	if !s.IsGroup {
		v = s.wrapStringer(v)
	}

	s.items = append(s.items, v)
}

func (s *state) Split() (rv *state) {
	cutoff := len(s.items) - s.splitpos

	if len(s.items) == 0 || cutoff == 0 {
		s.items = append(s.items, stringers.Empty{})
		s.splitpos++

		return s
	}

	tmp := make([]fmt.Stringer, cutoff)
	copy(tmp, s.items[s.splitpos:])

	s.items[s.splitpos] = stringers.Sequence(tmp)

	s.splitpos++
	s.items = s.items[:s.splitpos]

	return s
}

func (s *state) Stringer() (rv fmt.Stringer) {
	if s.IsGroup {
		rv = stringers.Random(s.items)
	} else {
		rv = stringers.Sequence(s.items)
	}

	return s.wrapStringer(rv)
}
