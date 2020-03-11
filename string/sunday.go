package string

import "fmt"

// Sunday to match substring in string
// http://wiki.jikexueyuan.com/project/kmp-algorithm/sunday.html

// 只不过 Sunday 算法是从前往后匹配，在匹配失败时关注的是文本串中参加匹配的最末位字符的下一位字符。
// 如果该字符没有在模式串中出现则直接跳过，即移动位数 = 匹配串长度 + 1；
// 否则，其移动位数 = 模式串中最右端的该字符到末尾的距离 +1。

// Sunday .
func Sunday(str, substr string) bool {
	strlen := len(str)
	substrlen := len(substr)
	if strlen < substrlen {
		return false
	}
	i, j := 0, 0
	moves := sundayHelper(substr)
	fmt.Printf("%v\n", moves)
	pos := 0 // 当前开始比较的位置

	for j < substrlen && i < strlen {
		println("loop: ", pos, i, j)
		if str[i] == substr[j] {
			i++
			j++
		} else {
			// true: not matched
			next := pos + substrlen
			if next >= strlen {
				break
			}

			if v, ok := moves[rune(str[next])]; ok {
				// true: 字符存在模式串中
				pos += v
				i = pos
			} else {
				// true: 字符不存在模式串中
				pos = pos + substrlen + 1
				i = pos
			}

			j = 0
		}
	}

	println("done", i, j)
	if j == substrlen {
		return true
	}

	return false
}

// sundayHepler 取得模式字符串中出现的字符，最右位置到末端的距离
// 如：acacab => map[a]int{a:2,c:3,b:1}
// moves['b'] = len("acacab") - max_idx("b") = 6 - 5 = 6 - 6 + 1
func sundayHelper(substr string) map[rune]int {
	l := len(substr)
	moves := make(map[rune]int)
	// move := make([]int, len(substr))
	for idx, v := range substr {
		moves[v] = l - idx
	}

	return moves
}
