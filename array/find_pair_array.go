package array

import (
	"fmt"
)

// Pair include num A and B
// A + B = target
type Pair struct {
	NumA int
	NumB int
}

// FindPairInArray find a pair in arr which sum(NumA, NumB) = target
// eg.
//
// target = 6, wanted {-3, 9}, {-3, 9}, {0, 6}, {0, 6}
//  i -->                 <-- j
// -3, -3, -1, 0, 1, 3, 6, 6, 9
//
// 前提：有序数组
func FindPairInArray(arr []int, target int) []Pair {

	i, j := 0, len(arr)-1
	pairs := make([]Pair, 0)
	wanted := target - arr[i]

	marki := i // 上一个匹配到期望值的i的位置
	markj := j // arr[i]匹配到期望值的的第一个位置

	for j > i {
		fmt.Printf("arr[%d]=%d, arr[%d]=%d, wanted=%d\n", i, arr[i], j, arr[j], wanted)
		if arr[j] == wanted {
			pairs = append(pairs, Pair{NumA: arr[i], NumB: arr[j]})
			// 如果i的值没有变化则不更新markj
			// markj 只用于表示arr[i]匹配到的第一个元素的位置
			if i != marki {
				markj = j
			}

			// 更新marki的值
			// marki 用于标记上一个匹配到期望值的位置元素的的位置i
			marki = i

			// arr[i] == wanted, 并不表示arr[j-1]不等于期望值
			// 因此继续判断是否arr[j] 有重复值
			if j-1 > i && arr[j-1] == arr[j] {
				j--
				continue
			}

			// 否则移动i指针，但是因为有 arr[i] == wanted, 并不表示arr[j-1]不等于期望值的逻辑
			// 因此 arr[i+1] == arr[i]的情况存在时候，需要将j指针重置到arr[i]匹配到的第一markj的位置
			if i+1 < j && arr[i+1] == arr[i] {
				j = markj
			}

			i++
			wanted = target - arr[i]
		} else if arr[j] < wanted {
			// 小于期望值，则不需要继续移动j指针
			i++
			wanted = target - arr[i]
		} else {
			// 超过期望值，继续移动j指针
			j--
		}
	}

	return pairs
}
