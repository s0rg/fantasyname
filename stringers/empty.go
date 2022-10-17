package stringers

// Empty handy shortcut to create empty literals.
type Empty struct{}

func (e Empty) String() (rv string) {
	return
}
