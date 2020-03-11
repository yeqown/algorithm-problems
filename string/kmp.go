// Package string .
// KMP(Knuth-Morris-Pratt) to match substring in string
//
// http://www.ruanyifeng.com/blog/2013/05/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm.html
// http://wiki.jikexueyuan.com/project/kmp-algorithm/define.html
package string

// KMP . O(m+n)
func KMP(str, substr string) (matched bool) {
	next := next(substr)
	// _ = next

	strlen := len(str)
	substrlen := len(substr)
	i := 0 // pointer of str
	j := 0 // pointer of substr

	for i <= strlen && j < substrlen {
		if j == -1 || str[i] == substr[j] {
			i++
			j++
		} else {
			// 最大长度表：失配时，模式串向右移动的位数为：已匹配字符数 - 失配字符的上一位字符所对应的最大长度值
			// j = j - next[j]
			// next数组：失配时，模式串向右移动的位数为：失配字符所在位置 - 失配字符对应的next值
			j = next[j]
		}
	}

	if j == substrlen {
		// 循环终止，且subtr被遍历完全
		matched = true
		return
	}

	return
}

// @substr "ABAB" => [0,0,1,2]
//
// 对于 P = p[0] p[1] ...p[j-1] p[j] 寻找模式串 P 中长度最大且相等的前缀和后缀。
// 如果存在 p[0] p[1] ...p[k-1] p[k] = p[j- k] p[j-k+1]...p[j-1] p[j]
// 那么在包含 p[j] 的模式串中有最大长度为 k+1 的相同前缀后缀
func next(substr string) []int {
	maxl := make([]int, len(substr))
	maxl = append(maxl, 0)

	if len(substr) == 1 {
		return maxl
	}

	j := 0
	k := 0

	for i := 1; i < len(substr); i++ {
		if substr[k] == substr[j] {

		}
	}

	return maxl
}
