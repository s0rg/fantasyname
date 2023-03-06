package fantasyname

import (
	"fmt"
	"math/rand"
)

type config struct {
	RandIntN   func(int) int
	Dictionary map[rune][]fmt.Stringer
	Collapse   bool
}

func (c *config) validate() {
	if c.RandIntN == nil {
		c.RandIntN = rand.Intn
	}

	if c.Dictionary == nil {
		c.Dictionary = symbolMap
	}
}

// Option is a configuration func.
type Option func(*config)

// Collapse creates option for collapsing doubles in result.
func Collapse(v bool) Option {
	return func(c *config) {
		c.Collapse = v
	}
}

// RandFn creates option with custom random func.
func RandFn(v func(int) int) Option {
	return func(c *config) {
		c.RandIntN = v
	}
}

// Dictionary sets custom dictionary (that maps runes to possible strings) for generation.
func Dictionary(d map[rune][]string) Option {
	return func(c *config) {
		t := make(map[rune][]fmt.Stringer)

		for k, v := range d {
			t[k] = convert(v)
		}

		c.Dictionary = t
	}
}
