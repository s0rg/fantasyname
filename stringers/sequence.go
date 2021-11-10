package stringers

import (
	"fmt"
	"strings"
)

// Sequence holds bunch of stringers, and construct result by calling every one in order.
type Sequence []fmt.Stringer

func (seq Sequence) String() string {
	var sb strings.Builder

	for i := 0; i < len(seq); i++ {
		sb.WriteString(seq[i].String())
	}

	return sb.String()
}
