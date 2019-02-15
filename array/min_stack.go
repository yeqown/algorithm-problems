package array

import (
	"fmt"
	"math"
)

// MinStack ...
type MinStack struct {
	length int
	min    int
	data   []int
}

// Constructor ...
/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{length: 0, min: math.MaxInt32, data: make([]int, 0)}
}

// Top ...
func (s *MinStack) Top() int {
	if s.length == 0 {
		return math.MinInt32
	}
	// fmt.Println(s)
	top := s.data[s.length-1]
	if top < 0 {
		top = s.min
	} else {
		top = s.min + top
	}
	return top
}

// GetMin ...
func (s *MinStack) GetMin() int {
	// fmt.Println(s)
	return s.min
}

// Pop Pop the top item of the stack and return it, if not ex return MIN value
func (s *MinStack) Pop() int {
	if s.length == 0 {
		return math.MinInt32
	}
	// get top element
	top := s.data[s.length-1]
	s.length--
	s.data = s.data[:s.length]

	if top < 0 {
		s.min, top = s.min-top, s.min
	} else {
		top = s.min + top
	}

	// fmt.Println(s)
	return top
}

// Push Push a value onto the top of the stack
func (s *MinStack) Push(value int) {
	s.data = append(s.data, value-s.min)
	s.length++
	if value < s.min {
		s.min = value
	}
	// fmt.Println(s)
}

func (s *MinStack) String() string {
	return fmt.Sprintf("Stack(data: %v, length: %d, min: %d)", s.data, s.length, s.min)
}
