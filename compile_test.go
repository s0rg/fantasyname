package fantasyname

import (
	"errors"
	"math/rand"
	"strings"
	"testing"
)

func TestCompileErrors(t *testing.T) {
	t.Parallel()

	type testCase struct {
		Pattern string
		Err     error
	}

	cases := []testCase{
		{Pattern: "", Err: ErrEmptyStack},
		{Pattern: "a>", Err: ErrEmptyStack},
		{Pattern: "a)", Err: ErrEmptyStack},
		{Pattern: "<a", Err: ErrUnbalancedGroup},
		{Pattern: "(a", Err: ErrUnbalancedGroup},
		{Pattern: "<a)", Err: ErrUnbalancedGroup},
		{Pattern: "(a>", Err: ErrUnbalancedGroup},
		{Pattern: "a|b", Err: ErrInvalidSplit},
	}

	for idx, tc := range cases {
		_, err := Compile(tc.Pattern)
		if err == nil {
			t.Errorf("no error for case: %d", idx)
		}

		if !errors.Is(err, tc.Err) {
			t.Errorf("unexpected error for case: %d: %v", idx, err)
		}
	}
}

func TestCompileTricky(t *testing.T) {
	t.Parallel()

	const (
		pat = "(((((<(((((((((((((((a)))))))))))))))>)))))"
		val = "a"
	)

	gen, err := Compile(pat)
	if err != nil {
		t.Errorf("uexpected error: %v", err)
	}

	if rv := gen.String(); rv != val {
		t.Errorf("unexpected result: '%s'", rv)
	}
}

func TestCompileMain(t *testing.T) {
	t.Parallel()

	const (
		pat  = "~(foo)c'<s|cvc>!(a)"
		valp = "A"
		vals = "oof"
	)

	gen, err := Compile(pat, RandFn(rand.Intn))
	if err != nil {
		t.Errorf("uexpected error: %v", err)
	}

	rv := gen.String()

	if !strings.HasPrefix(rv, vals) || !strings.HasSuffix(rv, valp) {
		t.Errorf("unexpected result: '%s'", rv)
	}
}

func TestCompileCollapse(t *testing.T) {
	t.Parallel()

	const (
		pat = "(xxooo)"
		val = "xoo"
	)

	gen, err := Compile(pat, Collapse(true))
	if err != nil {
		t.Errorf("uexpected error: %v", err)
	}

	if rv := gen.String(); rv != val {
		t.Errorf("unexpected result: '%s'", rv)
	}
}
