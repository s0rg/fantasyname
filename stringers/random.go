package stringers

import (
	"fmt"
	"math/rand"
)

type Random []fmt.Stringer

func (rnd Random) String() string {
	return rnd[rand.Intn(len(rnd))].String()
}
