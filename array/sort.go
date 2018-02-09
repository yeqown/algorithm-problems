// 常见排序算法
//参考： http://www.cnblogs.com/alsf/p/6606287.html， 有动图演示，更容易理解

package array

import (
	. "alg/_utils"
	// "fmt"
)

func BinaryInsertSort(nums []int) {

}

// 插入排序：（升序排列）
// 时间复杂度：O(n) - O(n^2)
// 空间复杂度：O(1)
func InsertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		// ele = nums[i]
		for j := i; j >= 1; j-- {
			if nums[j-1] > nums[j] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				break
			}
		}
	}
	return nums
}

// 选择排序
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func DirectSelectSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		min := nums[i]
		pos := i
		for j := i; j < len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
				pos = i
			}
		}
		// switch
		nums[i], nums[pos] = nums[pos], nums[i]
	}
	return nums
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

// 归并排序, 也是分治的思想
// 时间复杂度：O(nlogn) - O(n)
// 空间复杂度：O(n)
func MergeSort(nums []int, left, right int) []int {
	if right <= left {
		return nums[left : left+1]
	}

	mid := (left + right) / 2

	l := MergeSort(nums, left, mid)
	r := MergeSort(nums, mid+1, right)
	return merge(l, r)
}

func merge(l, r []int) []int {
	l_len := len(l)
	r_len := len(r)

	// malloc new memory
	f := make([]int, l_len+r_len)
	var fptr int
	rptr, lptr := 0, 0
	for rptr < r_len && lptr < l_len {
		// 升序排列
		fptr = rptr + lptr
		if r[rptr] > l[lptr] {
			f[fptr] = l[lptr]
			lptr++
		} else {
			f[fptr] = r[rptr]
			rptr++
		}
	}
	// if one array is longger
	if rptr >= r_len {
		for lptr < l_len {
			fptr = lptr + rptr
			f[fptr] = l[lptr]
			lptr++
		}
	} else {
		for rptr < r_len {
			fptr = lptr + rptr
			f[fptr] = r[rptr]
			rptr++
		}
	}
	// fmt.Println("right: ", r, "left:", l, "final: ", f)
	return f
}

// 快速排序, 分治的思想
// 时间复杂度：
// 空间复杂度：

// !!!! not pass
func QuickSort2(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid, i := nums[0], 1
	left, right := 0, len(nums)-1
	for i = 1; i <= right; {
		if nums[i] > mid {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		} else {
			nums[i], nums[left] = nums[left], nums[i]
			left++
			i++
		}
	}
	nums[left] = mid
	l := QuickSort2(nums[:left])
	r := QuickSort2(nums[left+1:])

	return append(l, r...)
}

func QuickSort(nums []int) []int {
	left := 0
	right := len(nums) - 1
	if right <= 0 {
		return nums
	}

	// fmt.Printf("left: %d, right: %d\n", left, right)
	// fmt.Println()
	// fmt.Println()
	// fmt.Println("sort with nums: ", nums)

	mid := (left + right) / 2
	// set mid num as pivot
	pivot := MidOfThree(nums[left], nums[right], nums[mid])
	// fmt.Printf("new quick sort with left: %d, right: %d, pivot: %d\n", left, right, pivot)

	pivot_pos := mid
	if nums[left] == pivot {
		pivot_pos = left
	} else if nums[right] == pivot {
		pivot_pos = right
	}

	pivot_pos, nums = partition(nums[left:right+1], pivot_pos)

	// fmt.Println("sorted nums:", nums)
	// fmt.Println("will new quick sort: left, pivot_pos, right", left, pivot_pos, right)

	l := QuickSort(nums[:pivot_pos])
	r := QuickSort(nums[pivot_pos:])

	return append(l, r...)
}

// 分区操作
func partition(nums []int, pivot_pos int) (int, []int) {
	// fmt.Printf("partition pivot_pos: %d\n", pivot_pos)
	pivot := nums[pivot_pos]

	less := []int{}
	greater := []int{}

	for i := 0; i < len(nums); i++ {
		if i == pivot_pos {
			continue
		}
		if nums[i] < pivot {
			less = append(less, nums[i])
		} else {
			greater = append(greater, nums[i])
		}
	}
	less = append(less, pivot)

	return len(less), append(less, greater...)
}
