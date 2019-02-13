package array

/*
Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
Find all the elements of [1, n] inclusive that do not appear in this array.
Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:
	Input:
	[4,3,2,7,8,2,3,1]
	Output:
	[5,6]
*/

func findDisappearedNumbers(nums []int) []int {
	length := len(nums)
	// 将对应数字交换到对应的位置：1位于0, 2位于1
	// [4, 3, 2, 7, 8, 2, 3] => [3, 2, 3, 4, 8, 2, 7]
	for i := 0; i < length; i++ {
		// 内部为什么是for, 而不是if
		for nums[i] > 0 && nums[i] <= length && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	out := make([]int, 0)
	for i := 0; i < length; i++ {
		if nums[i] != i+1 {
			out = append(out, i+1)
		}
	}
	return out
}
