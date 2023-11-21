package fantasyname

import (
	"math/rand"
	"testing"

	"github.com/s0rg/fantasyname/stringers"
	"github.com/s0rg/fantasyname/wrappers"
)

func TestStateLiteralAdd(t *testing.T) {
	t.Parallel()

	s := state{IsLiteral: true}
	s.Add(stringers.Literal("a"))
	s.Add(stringers.Literal("b"))
	s.Add(stringers.Literal("c"))

	const val = "abc"

	r := s.Stringer()

	if r.String() != val {
		t.Error("fail 1")
	}

	if r.String() != val {
		t.Error("fail 2")
	}
}

func TestStateLiteralWrap(t *testing.T) {
	t.Parallel()

	s := state{IsLiteral: true}
	s.Wrap(wrappers.Capitalized)
	s.Add(stringers.Literal("a"))
	s.Add(stringers.Literal("b"))
	s.Add(stringers.Literal("c"))

	const val = "Abc"

	r := s.Stringer()

	if r.String() != val {
		t.Error("unexpected result")
	}
}

func TestStateLiteralGroupSplitNotEmpty(t *testing.T) {
	t.Parallel()

	s := state{IsGroup: true, IsLiteral: true, rand: rand.Intn}
	s.Add(stringers.Literal("a"))
	s.Add(stringers.Literal("c"))
	s.Split()
	s.Add(stringers.Literal("d"))
	s.Add(stringers.Literal("c"))
	s.Split()

	res := s.Stringer().String()

	switch res {
	case "ac", "dc":
	default:
		t.Errorf("unexpected result: '%s'", res)
	}
}

func TestStateLiteralGroupSplitEmpty(t *testing.T) {
	t.Parallel()

	s := state{IsGroup: true, IsLiteral: true, rand: rand.Intn}
	s.Split()
	s.Add(stringers.Literal("a"))
	s.Split()

	res := s.Stringer().String()

	switch res {
	case "a", "":
	default:
		t.Errorf("unexpected result: '%s'", res)
	}
}
