/*
 Implement a basic calculator to evaluate a simple expression string.
 The expression string may contain open ( and closing parentheses ), the plus + or minus sign -, non-negative integers and empty spaces .
 You may assume that the given expression is always valid.
*/

/*
 中缀表达式                              ( 1 + 2 ) - 4 + 5
 前缀表达式(or Polish Notation)          + - + 1 2 4 5
 后缀表达式(or Reverse Polish Notation)  1 2 + 4 - 5 +
*/

package array

import (
	. "alg/_utils"

	"fmt"
)

var PriorityMap = map[rune]int{
	43: 1, // +
	45: 1, // -
	42: 2, // *
	47: 2, // /
	40: 3, // (
	41: 3, // )
}

func ConvMid2Polish(expr string) Stack {
	s1 := NewStack()
	s2 := NewStack()

	var (
		num = 0
	)

	for i := 0; i < len(expr); i++ {
		c := expr[i]

		if IsNumber(c) {
			num := ConvRuneNum2Int(c) + num*10
			continue
		} else if IsOpChar(c) {
			s2.Push(num)
			num = 0

		Compare:
			if s1.IsEmpty() || s1.Peek() == ')' {
				s1.Push(c)
			} else if PriorityMap[c] <= PriorityMap[s1.Peek()] {
				s1.Push(c)
			} else {
				s2.Push(s1.Pop())
				goto Compare
			}
		} else if IsParenthesis(c) {
			s2.Push(num)
			num = 0

			if IsRightParenthesis(c) {
				s1.Push(c)
			} else {
				for s1.Peek() != ')' {
					s2.Push(s1.Pop())
				}
			}
		}
	}

	for !s2.IsEmpty() {
		s1.Push(s2.Pop())
	}

	fmt.Println(s1)
	return s1
}
