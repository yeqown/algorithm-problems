package array

import (
	"fmt"
)

var (
	// MIN varible
	MIN = -1
)

// Stack ...
type Stack struct {
	top    *node
	length int
}

type node struct {
	value int
	prev  *node
}

// NewStack Create a new stack
func NewStack() *Stack {
	return &Stack{nil, 0}
}

// SetMIN Set MIN value by self default is -1
func SetMIN(min int) {
	MIN = min
}

// PeekMIN ...
func PeekMIN() int {
	return MIN
}

// Len Return the number of items in the stack
func (s *Stack) Len() int {
	return s.length
}

// IsEmpty if stack is empty return true
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

// Peek View the top item on the stack, if not ex return MIN
func (s *Stack) Peek() int {
	if s.length == 0 {
		return MIN
	}
	return s.top.value
}

// Pop Pop the top item of the stack and return it, if not ex return MIN value
func (s *Stack) Pop() int {
	if s.length == 0 {
		return MIN
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

// Push Push a value onto the top of the stack
func (s *Stack) Push(value int) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}

// Clear Clear stack, let it be empty
func (s *Stack) Clear() {
	s.top = nil
	s.length = 0
}

func (s *Stack) String() (str string) {
	for !s.IsEmpty() {
		str += fmt.Sprintf("%d,", s.Pop())
	}
	return
}
