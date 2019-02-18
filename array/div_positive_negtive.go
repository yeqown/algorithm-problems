package array

/* 将一个含有N各元素的数组，进行正负分开；要求正数在左边，负数在右边（没有0值,正负数无排序要求）【考虑如果有0值怎么办】
 */

func divPositiveNegative(nums []int) []int {
	s, e := 0, len(nums)-1
	for s < e {
		if nums[s] < 0 {
			if nums[e] > 0 {
				nums[s], nums[e] = nums[e], nums[s]
				s++
				e--
			} else {
				// nop
				e--
			}
		} else {
			if nums[e] > 0 {
				// nop
				s++
			} else {
				// nop
				s++
				e--
			}
		}
	}

	return nums
}
