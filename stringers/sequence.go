package stringers

import (
	"fmt"
	"strings"
)

// Sequence holds bunch of stringers, and construct result by calling every one in order.
type Sequence []fmt.Stringer

func (s Sequence) String() string {
	var b strings.Builder

	for i := 0; i < len(s); i++ {
		b.WriteString(s[i].String())
	}

	return b.String()
}
