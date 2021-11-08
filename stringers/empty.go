package stringers

type Empty struct{}

func (emp Empty) String() (rv string) {
	return
}
