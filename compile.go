package fantasyname

import (
	"fmt"
	"strings"

	"github.com/s0rg/fantasyname/wrappers"
)

// Compile creates fmt.Stringer from pattern, you can call ia as many times as you wish,
// it will produce new name on every call. Under the hood it uses 'math/rand', so you
// need to initialize it with seed on your own.
func Compile(pattern string) (rv fmt.Stringer, err error) {
	if pattern = strings.TrimSpace(pattern); pattern == "" {
		return nil, ErrEmptyStack
	}

	p := newParser()

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
