package stringers

import (
	"fmt"
	"strings"
)

type Sequence []fmt.Stringer

func (seq Sequence) String() string {
	var sb strings.Builder

	for i := 0; i < len(seq); i++ {
		sb.WriteString(seq[i].String())
	}

	return sb.String()
}
