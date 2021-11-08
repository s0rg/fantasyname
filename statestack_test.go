package fantasyname

import "testing"

func TestStateStackEmpty(t *testing.T) {
	t.Parallel()

	stack := statestack{}

	if len(stack) != 0 {
		t.Error("not empty")
	}

	if _, ok := stack.Top(); ok {
		t.Error("empty has top")
	}

	if _, ok := stack.Pop(); ok {
		t.Error("empty can pop")
	}
}

func TestStateStackPushPop(t *testing.T) {
	t.Parallel()

	stack := statestack{}

	stack.Push(state{IsLiteral: true})

	if len(stack) != 1 {
		t.Error("unexpected len")
	}

	top, ok := stack.Top()
	if !ok {
		t.Error("no top")
	}

	if !top.IsLiteral {
		t.Error("bad top")
	}

	val, ok := stack.Pop()
	if !ok {
		t.Error("no pop")
	}

	if !val.IsLiteral {
		t.Error("bad val")
	}

	if len(stack) != 0 {
		t.Error("not empty")
	}
}
