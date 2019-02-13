package array

/*
* @param nums []int 有序序列
* @param target int 目标数
* @return pos int 目标数的位置
 */

// MISS ...
const MISS = -1

// BinarySearch ...
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
