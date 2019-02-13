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

// "fmt"

// TrapRainWater approch 1
// 总结：时间耗费和空间耗费都比较高，意味着还需要继续训练算法思想，这个算比较笨的蛮力法
// Time Limit Exceeded
func TrapRainWater(nums []int) int {
	var trappedRw int

	for true {
		if len(nums) < 3 {
			break
		}

		nums = trimZeroBesideOfArray(nums)
		// fmt.Println("trim array:", nums)

		for i := 0; i < len(nums); i++ {
			// nums[i] is 0
			if nums[i] == 0 {
				trappedRw++
			}
		}
		nums = lessOneArray(nums)
		// fmt.Println("less one array:", nums, "now trap_water: ", trappedRw)
	}
	return trappedRw
}

// 去除数组两端的0元素
func trimZeroBesideOfArray(nums []int) []int {
	l := len(nums)
	i := 0
	for nums[i] == 0 {
		i++
		if i > l-1 {
			return []int{}
		}
	}
	j := len(nums) - 1
	for nums[j] == 0 {
		j--
		if j < 0 {
			return []int{}
		}
	}
	// fmt.Println(nums[i:j+1], i, j)
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

// TrapRainWaterApproch2 approch 2
//
func TrapRainWaterApproch2(height []int) int {
	l := len(height)
	var (
		left, leftMax   int
		right, rightMax int = l - 1, 0
		traped          int
	)

	for left < right {
		// 为什么这样设置切换条件呢？
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				traped += (leftMax - height[left])
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				traped += (rightMax - height[right])
			}
			right--
		}
	}
	return traped
}
