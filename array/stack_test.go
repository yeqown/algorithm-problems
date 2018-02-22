package array

import (
	"testing"
)

func Test_Stack(t *testing.T) {
	s := NewStack()
	if l := s.Len(); l != 0 {
		t.Fatal("error initial stack")
	}
	if PeekMIN() != -1 {
		t.Fatal("error PeekMIN")
	}

	if s.Pop() != PeekMIN() {
		t.Fatal("error got Stack.Push")
	}

	s.Push(99)

	if s.Pop() != 99 {
		t.Fatal("error got Stack.Pop")
	}

}

func Test_StackLen(t *testing.T) {
	s := NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	if s.Len() != 4 {
		t.Fatal("error got Stack.Len")
	}
}

func Test_SetMIN(t *testing.T) {
	SetMIN(99)

	if PeekMIN() != 99 {
		t.Fatal("error SetMIN")
	}
}
