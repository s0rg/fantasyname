package fantasyname

import (
	"errors"
	"fmt"

	"github.com/s0rg/fantasyname/stringers"
	"github.com/s0rg/fantasyname/wrappers"
)

var (
	ErrEmptyStack      = errors.New("stack is empty")
	ErrInvalidSplit    = errors.New("out-of-group split")
	ErrInvalidGroupEnd = errors.New("unbalanced group")
)

type parser struct {
	stack statestack
}

func newParser() (p *parser) {
	p = &parser{}

	p.stack.Push(state{}) // "root" state

	return p
}

func (p *parser) OnGroupStart(isLiteral bool) {
	p.stack.Push(state{IsGroup: true, IsLiteral: isLiteral})
}

func (p *parser) OnGroupEnd() (err error) {
	cur, ok := p.stack.Pop()
	if !ok {
		return ErrEmptyStack
	}

	top, ok := p.stack.Top()
	if !ok {
		return ErrEmptyStack
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
		return ErrInvalidGroupEnd
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
		if sym, ok := symbolMap[r]; ok {
			v = stringers.Random(sym)
		}
	}

	if v == nil {
		v = stringers.Literal(string(r))
	}

	top.Add(v)

	return nil
}

func (p *parser) Build() (s fmt.Stringer, err error) {
	top, ok := p.stack.Top()
	if !ok {
		return nil, ErrEmptyStack
	}

	return wrappers.Collapsed(top.Stringer()), nil
}

func Parse(pattern string) (rv fmt.Stringer, err error) {
	p := newParser()

	for pos, r := range pattern {
		switch r {
		case '<':
			p.OnGroupStart(false)
		case '(':
			p.OnGroupStart(true)
		case '>', ')':
			err = p.OnGroupEnd()
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
