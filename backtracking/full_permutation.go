package backtracking

var result [][]int

// FullPermutation . 全排列
// 123 => 123 132 213 231 312 321
func FullPermutation(arr []int) [][]int {
	choice := make([]int, 0, len(arr))
	fullPermutationBacktrack(arr, choice)
	return result
}

func fullPermutationBacktrack(arr []int, choice []int) {
	if len(arr) == len(choice) {
		dst := make([]int, len(arr))
		copy(dst, choice)
		result = append(result, dst)
		choice = make([]int, 0, len(arr))
		return
	}

	for _, v := range arr {
		if fullPermutationContains(choice, v) {
			continue
		}
		choice = append(choice, v)
		fullPermutationBacktrack(arr, choice)
		choice = choice[:len(choice)-1]
	}
}

func fullPermutationContains(arr []int, tatget int) bool {
	for _, v := range arr {
		if v == tatget {
			return true
		}
	}

	return false
}
