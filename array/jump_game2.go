/*
Given an array of non-negative integers,
you are initially positioned at the first index of the array.
Each element in the array represents your maximum jump length at that position.
Your goal is to reach the last index in the minimum number of jumps.

For example:
Given array A = [2,3,1,1,4], The minimum number of jumps to reach the last index is 2.
(Jump 1 step from index 0 to 1, then 3 steps to the last index.)

Note:
	You can assume that you can always reach the last index.
*/

package array

import (
	"fmt"
)

func findMaxValue(nums []int, start, limit int) (step int) {
	max_weight := nums[start] + start
	step = 1
	end := start + limit

	for i := start + 1; i <= end; i++ {
		weight := i + nums[i]
		if weight > max_weight {
			max_weight = weight
			step = i - start
		}
	}
	return
}

// 使用贪心算法求解
func Jump2(nums []int) int {
	var (
		jump_time int = 0
	)

	pos := 0
	for pos < len(nums)-1 {
		if pos+nums[pos] >= len(nums)-1 {
			fmt.Println("out of range")
			jump_time++
			break
		}

		step := findMaxValue(nums, pos, nums[pos])
		fmt.Println("pos: ", pos, "nums[pos]:", nums[pos], "step:", step)

		pos += step
		jump_time++
	}
	return jump_time
}
