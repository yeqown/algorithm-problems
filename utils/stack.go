package utils

import (
	"container/list"
	"fmt"
)

// Stack ....
type Stack struct {
	list *list.List
}

// NewStack ...
func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

// Push method of stack
func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

// Pop method of stack
func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

// Peak method of stack
func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

// Len method of Stack
func (stack *Stack) Len() int {
	return stack.list.Len()
}

// Empty method of Stack
func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}

// String method of Stack
func (stack *Stack) String() (s string) {
	elem := stack.list.Front()
	if elem == nil {
		return s
	}
	for elem.Next() != nil {
		s += fmt.Sprintf("%v,", elem.Value)
		elem = elem.Next()
	}
	return s
}
