package fantasyname

import "math/rand"

type config struct {
	Collapse bool
	RandIntN func(int) int
}

type Option func(*config)

func Collapse(v bool) Option {
	return func(c *config) {
		c.Collapse = v
	}
}

func RandFn(v func(int) int) Option {
	return func(c *config) {
		c.RandIntN = v
	}
}

func (c *config) validate() {
	if c.RandIntN == nil {
		c.RandIntN = rand.Intn
	}
}
