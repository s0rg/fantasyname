package stringers

// Empty handy shortcut to create empty literals.
type Empty struct{}

func (emp Empty) String() (rv string) {
	return
}
