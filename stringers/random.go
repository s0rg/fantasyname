package stringers

import (
	"fmt"
	"math/rand"
)

// Random holds bunch of stringers, and choose among them on every call.
type Random []fmt.Stringer

func (rnd Random) String() string {
	return rnd[rand.Intn(len(rnd))].String()
}
