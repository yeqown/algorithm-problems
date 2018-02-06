// 最长连续增长子序列

/*
* Q: Given an unsorted array of integers, find the number of longest increasing subsequence.
* Example:
*  Input: [1,3,5,4,7]
*  Output: 2
*  Explanation: The two longest increasing subsequence are [1, 3, 4, 7] and [1, 3, 5, 7].
 */

package dp

import (
	"fmt"
	// "math"
)

type NodeState struct {
	Longest int
	Lis     []int
}

func LIS(nums []int) []int {

	// longest 用来标记最终最长子序列的终结位置和上一个state位置
	longest := &NodeState{Longest: 1, Lis: []int{nums[0]}}
	states := map[int]*NodeState{}

	for i := 0; i < len(nums); i++ {

		// 找到states中值小于nums[i]的最长子序列，如果没有states[i]为自身
		maxlen := 1
		prev := -1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && states[j].Longest+1 > maxlen {
				maxlen = states[j].Longest
				prev = j
			}
		}
		if prev != -1 {
			maxlen += 1
		}
		states[i] = &NodeState{Longest: maxlen}

		if states[i].Longest > longest.Longest {
			longest.Longest = states[i].Longest
			longest.Lis = append(longest.Lis, nums[i])
		}
	}

	fmt.Println(longest)
	return longest.Lis
}
