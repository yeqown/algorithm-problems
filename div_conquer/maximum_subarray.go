// 最大子数组

/*
the maximum subarray must be:
* left part totaly > left_max
* right part totaly > right_max
* cross mid position > mid_max
* Max(left, mid, right)
*/

package div_conquer

import (
	"math"
)

func MaxOfThree(a, b, c int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

func MaxSubArray(nums []int, left, right int) int {
	if left >= right {
		return nums[left]
	}

	mid := (left + right) / 2

	left_max := MaxSubArray(nums, left, mid)
	right_max := MaxSubArray(nums, mid+1, right)
	// this is so much important
	cross_max := crossMidMaxSubArray(left, right, nums)

	// this can be thinking as merge function
	return MaxOfThree(left_max, right_max, cross_max)
}

func crossMidMaxSubArray(left, right int, nums []int) int {
	// how to find cross mid
	mid := (left + right) / 2
	left_max := int(math.MinInt64)
	right_max := int(math.MinInt64)

	sumary := 0
	for i := mid; i >= left; i-- {
		sumary += nums[i]
		if sumary > left_max {
			left_max = sumary
		}
	}

	sumary = 0
	for i := mid + 1; i <= right; i++ {
		sumary += nums[i]
		if sumary > right_max {
			right_max = sumary
		}
	}
	return left_max + right_max
}
