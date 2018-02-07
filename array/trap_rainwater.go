/*
Given n non-negative integers representing an elevation map where the width of each bar is 1,
compute how much water it is able to trap after raining.

For example,
Given [0,1,0,2,1,0,1,3,2,1,2,1], return 6.
|
|                     3
|          2          __  2    2
|   1     __  1    1 |__|__ 1  __  1
|0  __ 0 |__|__  0 __|__|__|__|__|__
|__|__|__|__|__|__|__|__|__|__|__|__|
*/

// leetcode solution

package array

import (
// "fmt"
)

// approch 1
// Time Limit Exceeded
func TrapRainWater(nums []int) int {
	var trapped_rw int = 0

	for true {
		if len(nums) < 3 {
			break
		}

		nums = trimZeroBesideOfArray(nums)
		// fmt.Println("trim array:", nums)

		for i := 0; i < len(nums); i++ {
			// nums[i] is 0
			if nums[i] == 0 {
				trapped_rw += 1
			}
		}
		nums = lessOneArray(nums)
		// fmt.Println("less one array:", nums, "now trap_water: ", trapped_rw)
	}
	return trapped_rw
}

// 去除数组两端的0元素
func trimZeroBesideOfArray(nums []int) []int {
	l := len(nums)
	i := 0
	for nums[i] == 0 {
		if i >= l-1 {
			return []int{}
		}
		i++
	}
	j := len(nums) - 1
	for nums[j] == 0 {
		j--
		if j <= 0 {
			return []int{}
		}
	}
	// fmt.Println(nums[i:j+1], i, j)
	// new_array = append(new_array, nums[i:j]...)
	// fmt.Println(new_array)
	return nums[i : j+1]
}

func lessOneArray(nums []int) []int {
	for pos, val := range nums {
		if val-1 < 0 {
			nums[pos] = 0
			continue
		}
		nums[pos] = val - 1
	}
	return nums
}

// approch 2
func TrapRainWaterApproch2(height []int) int {
	l := len(height)
	var (
		left      int = 0
		right     int = l - 1
		left_max  int = 0
		right_max int = 0
		traped    int = 0
	)

	for left < right {
		// 为什么这样设置切换条件呢？
		if height[left] < height[right] {
			if height[left] >= left_max {
				left_max = height[left]
			} else {
				traped += (left_max - height[left])
			}
			left++
		} else {
			if height[right] >= right_max {
				right_max = height[right]
			} else {
				traped += (right_max - height[right])
			}
			right--
		}
	}
	return traped
}
