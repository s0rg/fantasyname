package stringers

// Literal wraps string/rune to make it fmt.Stringer.
type Literal string

func (l Literal) String() (rv string) {
	return string(l)
}
