package array

import (
	"testing"
)

func Test_MinStack(t *testing.T) {
	ms := Constructor()

	ms.Push(5)
	ms.Push(3)
	ms.Push(1)
	ms.Push(2)
	ms.Push(0)

	if top := ms.Top(); top != 0 {
		t.Errorf("wrong top value, got: %d, want: %d", top, 0)
	}

	if min := ms.GetMin(); min != 0 {
		t.Errorf("wrong min value, got: %d, want: %d", min, 0)
	}

	if pop := ms.Pop(); pop != 0 {
		t.Errorf("wrong pop value, got: %d, want: %d", pop, 0)
	}

	if min := ms.GetMin(); min != 1 {
		t.Errorf("wrong min value, got: %d, want: %d", min, 1)
	}

	if top := ms.Top(); top != 2 {
		t.Errorf("wrong top value, got: %d, want: %d", top, 2)
	}

}
