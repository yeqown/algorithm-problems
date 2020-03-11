package string

// BM(Boyer–Moore) to match substring in string
// http://wiki.jikexueyuan.com/project/kmp-algorithm/bm.html
// 该算法从模式串的尾部开始匹配，且拥有在最坏情况下 O(N) 的时间复杂度。在实践中，比 KMP 算法的实际效能高。
//
// BM 算法定义了两个规则：
// 坏字符规则：当文本串中的某个字符跟模式串的某个字符不匹配时，我们称文本串中的这个失配字符为坏字符，此时模式串需要向右移动：
// 	移动的位数 = 坏字符在模式串中的位置 - 坏字符在模式串中最右出现的位置。此外，如果"坏字符"不包含在模式串之中，则最右出现位置为 -1。
//
// 好后缀规则：当字符失配时，
// 	移动的位数 = 好后缀在模式串中的位置 - 好后缀在模式串上一次出现的位置，且如果好后缀在模式串中没有再次出现，则为 -1。

// BoyerMoore . O(m/n)
func BoyerMoore(str, substr string) bool {
	strlen := len(str)
	substrlen := len(substr)

	if substrlen > strlen {
		return false
	}

	i, j := substrlen, substrlen
	pos := 0
	matchedCnt := 0
	badchar := badcharHelper(substr)
	// goodsuffix := goodsuffixHelper(substr)

	for i < strlen && j > 0 {
		if str[i] == substr[j] {
			i--
			j--
			matchedCnt++
		} else {
			// TODO:
			// 不匹配:
			// 应用坏字符或者好后缀
			pos += maxBadcharOrGoodsuffix(rune(str[i]), j, matchedCnt, badchar)
			i = pos
			j = substrlen
		}
	}

	return false
}

// c 坏字符
// j 坏字符在模式串的位置
// matchedCnt 已经匹配的字符数 用于好后缀判定
// moves 字符在模式串中最右的位置
func maxBadcharOrGoodsuffix(c rune, j, matchedCnt int, moves map[rune]int) int {
	var (
		badcharMov    int
		goodsuffixMov int
	)

	if v, ok := moves[c]; ok {
		badcharMov = j - v
	} else {
		v = -1
		badcharMov = j - v
	}

	if matchedCnt == 0 {
		return badcharMov
	}

	// TODO:
	goodsuffixMov = 1

	return max(badcharMov, goodsuffixMov)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func goodsuffixHelper(substr string) []int {
	substrlen := len(substr)
	gs := make([]int, substrlen)

	return gs
}

// Like to sunday.go sundayHelper
func badcharHelper(substr string) map[rune]int {
	moves := make(map[rune]int)
	substrlen := len(substr)

	for idx, v := range substr {
		moves[v] = substrlen - idx //
	}

	return moves
}
