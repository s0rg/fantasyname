package fantasyname

import "math/rand"

type config struct {
	RandIntN func(int) int
	Collapse bool
}

func (c *config) validate() {
	if c.RandIntN == nil {
		c.RandIntN = rand.Intn
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
