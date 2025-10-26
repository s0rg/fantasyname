package fantasyname

import (
	"errors"
	"testing"

	"github.com/s0rg/fantasyname/wrappers"
)

func TestParserErrors(t *testing.T) {
	t.Parallel()

	var e error

	p := &parser{}

	if e = p.OnGroupEnd(true); e == nil {
		t.Error("group end: no error")
	}

	if !errors.Is(e, ErrEmptyStack) {
		t.Error("group end: unexpected error")
	}

	if e = p.OnGroupSplit(); e == nil {
		t.Error("group split: no error")
	}

	if !errors.Is(e, ErrEmptyStack) {
		t.Error("group split: unexpected error")
	}

	if e = p.OnSymbol(' '); e == nil {
		t.Error("symbol: no error")
	}

	if !errors.Is(e, ErrEmptyStack) {
		t.Error("symbol: unexpected error")
	}

	if e = p.Wrap(wrappers.Collapsed); e == nil {
		t.Error("wrap: no error")
	}

	if !errors.Is(e, ErrEmptyStack) {
		t.Error("wrap: unexpected error")
	}
}
