// 常见排序算法

//参考： http://www.cnblogs.com/alsf/p/6606287.html， 有动图演示，更容易理解

package array

func QuickSort(nums []int) []int {

}

func SwitchSort(nums []int) []int {

}

func BinaryInsertSort(numt []int) []int {

}

func InsertSort(numt []int) []int {

}

func DirectSelectSort(numt []int) []int {

}

// 归并排序
// 时间复杂度：O(nlogn)
// 空间复杂度：O(n)
func MergeSort(nums []int, left, right int) []int {
	if right <= left {
		return nums[left:right]
	}

	mid := (left + right) / 2

	l := MergeSort(nums, left, mid)
	r := MergeSort(nums, mid+1, right)
	return merge(l, r)
}

func merge(l, r) []int {
	l_len := len(l)
	r_len := len(r)

	f := make([]int, l_len+r_len)

	for i := 0; i < len(l); i++ {

	}
}

// 冒泡排序：
// 时间复杂度：O(n) - O(n^2)
// 空间复杂度：O(1)
func BubbleSort(nums []int) []int {
	l := len(nums)
	var flag bool = false

	for i := 0; i < l; i++ {
		for j := 1; j < l-i; j++ {
			// incr sort
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return nums
}
