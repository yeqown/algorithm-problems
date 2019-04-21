package array

import (
	"math"
)

/*
* @param nums []int 有序序列
* @param target int 目标数
* @return pos int 目标数的位置
 */

// MISS ... if has no element in array return this CODE to mark
const MISS = -1

// LinearSearch ... form index 0 to index n
func LinearSearch(nums []int, target int) int {
	for idx, v := range nums {
		if v == target {
			return idx
		}
	}
	return MISS
}

// BinarySearch ...
// from index `left` to `right`
// if equal to middle return middle
// else if bigger than middle then find the right block
// else find the left block
func BinarySearch(nums []int, left, right, target int) int {
	mid := (left + right) / 2
	pos := MISS

	// fmt.Printf("left %d, mid %d, right %d\n", left, mid, right)

	if left > right {
		return MISS
	}

	if nums[mid] == target {
		// fmt.Println("target is target")
		pos = mid
	} else if nums[mid] > target {
		pos = BinarySearch(nums, left, mid-1, target)
	} else {
		pos = BinarySearch(nums, mid+1, right, target)
	}

	return pos
}

// JumpSearch ... search a element in an ordered array
// n = len(arr)
// m = block size to jump
//
// search from index 0 to m:
func JumpSearch(nums []int, target int) int {
	// best step size: m = sqrt(n)
	step := int(math.Sqrt(float64(len(nums))))
	n := len(nums)

	var prev, cur = 0, 0

	// find the bigger element index by jump m size
	for nums[cur] < target {
		prev, cur = cur, cur+step
		if cur > n-1 {
			cur = n - 1
			break
		}
	}

	if nums[cur] == target {
		return cur
	}

	// start linear search
	linearIdx := prev + 1
	for linearIdx <= cur {
		if nums[linearIdx] == target {
			return linearIdx
		}
		linearIdx++
	}

	return MISS
}
