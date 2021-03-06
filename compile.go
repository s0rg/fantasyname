package fantasyname

import (
	"fmt"
	"strings"

	"github.com/s0rg/fantasyname/wrappers"
)

// Compile creates fmt.Stringer from given pattern, you can call it as many times as
// you wish, it will produce new name on every call.
func Compile(pattern string, opts ...Option) (rv fmt.Stringer, err error) {
	if pattern = strings.TrimSpace(pattern); pattern == "" {
		return nil, ErrEmptyStack
	}

	p := newParser(opts...)

	for pos, r := range pattern {
		switch r {
		case '<':
			p.OnGroupStart(false)
		case '>':
			err = p.OnGroupEnd(false)
		case '(':
			p.OnGroupStart(true)
		case ')':
			err = p.OnGroupEnd(true)
		case '|':
			err = p.OnGroupSplit()
		case '!':
			err = p.Wrap(wrappers.Capitalized)
		case '~':
			err = p.Wrap(wrappers.Reversed)
		default:
			err = p.OnSymbol(r)
		}

		if err != nil {
			return nil, fmt.Errorf("pos: %d error: %w", pos, err)
		}
	}

	return p.Build()
}
