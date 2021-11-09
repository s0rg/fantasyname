package stringers

// Literal wraps string/rune to make it fmt.Stringer.
type Literal string

func (lit Literal) String() (rv string) {
	return string(lit)
}
