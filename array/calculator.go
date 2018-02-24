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

// 由于堆栈的实现原因导致，无法存储float类型数据，如果包含除法运算会出现结果不准

package array

import (
	. "alg/_utils"

	"math"
	"strconv"
	"strings"
	// "fmt"
)

var PriorityMap = map[int]int{
	43: 1, // +
	45: 1, // -
	42: 2, // *
	47: 2, // /
	40: 3, // (
	41: 3, // )
}

func PolishTonationCalcu(polish string) int {
	splited := strings.Split(polish, ",")
	s := NewStack()
	for i := len(splited) - 1; i >= 0; i-- {
		if splited[i] == "" {
			continue
		}
		num, _ := strconv.ParseFloat(splited[i], 64)
		if num < 0 {
			// 操作符
			switch rune(math.Abs(num)) {
			case '+':
				s.Push(s.Pop() + s.Pop())
			case '-':
				s.Push(s.Pop() - s.Pop())
			case '*':
				s.Push(s.Pop() * s.Pop())
			case '/':
				s.Push(int(s.Pop() / s.Pop()))
			default:
				panic("wrong op char")
			}
		} else {
			s.Push(int(num))
		}
	}
	return s.Pop()
}

// 都是整型，如何区分操作符和数字
func ConvMid2Polish(expr string) string {
	s1 := NewStack()
	s2 := NewStack()

	// 用来处理多位整数
	var (
		num    int     = 0
		bitNum float64 = 0.0
	)

	for i := len(expr) - 1; i >= 0; i-- {
		c := rune(expr[i])
		int_c := int(c)
		neg_c := -int_c

		if IsNumber(c) {
			// 数字字符
			int_c = int_c - 48
			if num != 0 {
				bitNum++
			}
			num = int_c*int(math.Pow(10, bitNum)) + num

			if i <= 0 {
				s2.Push(num)
			}
		} else {
			// 非数字字符
			if num != 0 {
				s2.Push(num)
				num = 0
				bitNum = 0
			}

			if IsOpChar(c) {
			Compare:
				s1_peek := int(math.Abs(float64(s1.Peek())))

				if s1.IsEmpty() || s1_peek == ')' {
					s1.Push(neg_c)
				} else if PriorityMap[int_c] <= PriorityMap[s1_peek] {
					s1.Push(neg_c)
				} else {
					s2.Push(s1.Pop())
					goto Compare
				}
			} else if IsParenthesis(c) {
				if IsRightParenthesis(c) {
					s1.Push(int_c)
				} else {
					// 消除一对括号
					for s1.Peek() != ')' && !s1.IsEmpty() {
						s2.Push(s1.Pop())
					}
					s1.Pop()
				}
			}
		}
	} // for loop

	// 把所有S1中的剩余元素全部压入S2中
	for !s1.IsEmpty() {
		s2.Push(s1.Pop())
	}
	return s2.String()
}
