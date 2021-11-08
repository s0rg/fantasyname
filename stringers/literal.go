package stringers

type Literal string

func (lit Literal) String() (rv string) {
	return string(lit)
}
