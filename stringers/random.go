package stringers

import (
	"fmt"
)

// Random holds bunch of stringers, and choose among them on every call.
type Random struct {
	rndfn func(int) int
	items []fmt.Stringer
}

// MakeRandom create new Random stringer from list of stringers and random func.
func MakeRandom(items []fmt.Stringer, rndfn func(int) int) fmt.Stringer {
	return &Random{items: items, rndfn: rndfn}
}

// String implements stringer for Random.
func (r *Random) String() string {
	return r.items[r.rndfn(len(r.items))].String()
}
