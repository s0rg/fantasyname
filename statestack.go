package fantasyname

type statestack []*state

func (ss *statestack) Push(s *state) {
	*ss = append(*ss, s)
}

func (ss *statestack) Pop() (rv *state, ok bool) {
	sp := *ss

	if n := len(sp); n > 0 {
		n--
		rv, ok, *ss = sp[n], true, sp[:n]
	}

	return
}

func (ss *statestack) Top() (rv *state, ok bool) {
	sp := *ss

	if n := len(sp); n > 0 {
		rv, ok = sp[n-1], true
	}

	return
}
