package fantasyname

import (
	"errors"
	"fmt"

	"github.com/s0rg/fantasyname/stringers"
	"github.com/s0rg/fantasyname/wrappers"
)

var (
	// ErrEmptyStack is returned, if parser counts stack as non-empty, but it was.
	ErrEmptyStack = errors.New("empty stack")
	// ErrInvalidSplit is returned when split - "|" is found in pattern out of any group.
	ErrInvalidSplit = errors.New("invalid split")
	// ErrUnbalancedGroup is returned when stuck upon unexpected group close/open brackets.
	ErrUnbalancedGroup = errors.New("unbalanced group")
)

type parser struct {
	conf  *config
	stack statestack
}

func newParser(opts ...Option) (p *parser) {
	c := &config{}

	for _, o := range opts {
		o(c)
	}

	c.validate()

	p = &parser{conf: c}

	// "root" state
	p.stack.Push(&state{
		rand: p.conf.RandIntN,
	})

	return p
}

func (p *parser) OnGroupStart(isLiteral bool) {
	p.stack.Push(&state{
		IsGroup:   true,
		IsLiteral: isLiteral,
		rand:      p.conf.RandIntN,
	})
}

func (p *parser) OnGroupEnd(isLiteral bool) (err error) {
	cur, ok := p.stack.Pop()
	if !ok {
		return ErrEmptyStack
	}

	top, ok := p.stack.Top()
	if !ok {
		return ErrEmptyStack
	}

	if !cur.IsGroup || cur.IsLiteral != isLiteral {
		return ErrUnbalancedGroup
	}

	top.Add(cur.Split().Stringer())

	return nil
}

func (p *parser) OnGroupSplit() (err error) {
	top, ok := p.stack.Top()
	if !ok {
		return ErrEmptyStack
	}

	if !top.IsGroup {
		return ErrInvalidSplit
	}

	top.Split()

	return nil
}

func (p *parser) Wrap(w wrapper) (err error) {
	top, ok := p.stack.Top()
	if !ok {
		return ErrEmptyStack
	}

	top.Wrap(w)

	return nil
}

func (p *parser) OnSymbol(r rune) (err error) {
	top, ok := p.stack.Top()
	if !ok {
		return ErrEmptyStack
	}

	var v fmt.Stringer

	if !top.IsLiteral {
		if sym, ok := p.conf.Dictionary[r]; ok {
			v = stringers.MakeRandom(sym, p.conf.RandIntN)
		}
	}

	if v == nil {
		v = stringers.Literal(string(r))
	}

	top.Add(v)

	return nil
}

func (p *parser) Build() (s fmt.Stringer, err error) {
	if len(p.stack) != 1 {
		return nil, ErrUnbalancedGroup
	}

	// no need to check `ok` - its already checked above.
	top, _ := p.stack.Top()

	s = top.Stringer()

	if p.conf.Collapse {
		s = wrappers.Collapsed(s)
	}

	return s, nil
}
