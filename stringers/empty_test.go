package stringers

import "testing"

func TestEmpty(t *testing.T) {
	t.Parallel()

	e := &Empty{}
	if e.String() != "" {
		t.Error("unexpected result")
	}
}
