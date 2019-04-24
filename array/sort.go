// 常见排序算法
//参考： http://www.cnblogs.com/alsf/p/6606287.html， 有动图演示，更容易理解

package array

import (
	"fmt"
	"github.com/yeqown/algorithm-problems/utils"
)

// "fmt"

// BinaryInsertSort ...
func BinaryInsertSort(nums []int) {

}

// InsertSort ...
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

// DirectSelectSort ...
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

// BubbleSort 冒泡排序：
// 时间复杂度：O(n) - O(n^2)
// 空间复杂度：O(1)
func BubbleSort(nums []int) []int {
	l := len(nums)
	var (
		flag bool
	)

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

// MergeSort 归并排序, 也是分治的思想
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
	leftLen := len(l)
	rightLen := len(r)

	// malloc new memory
	f := make([]int, leftLen+rightLen)
	var fptr int
	rptr, lptr := 0, 0
	for rptr < rightLen && lptr < leftLen {
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
	if rptr >= rightLen {
		for lptr < leftLen {
			fptr = lptr + rptr
			f[fptr] = l[lptr]
			lptr++
		}
	} else {
		for rptr < rightLen {
			fptr = lptr + rptr
			f[fptr] = r[rptr]
			rptr++
		}
	}
	// fmt.Println("right: ", r, "left:", l, "final: ", f)
	return f
}

// QuickSort2 快速排序, 分治的思想
// 时间复杂度：
// 空间复杂度：

// QuickSort2 !!!! not pass
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

// QuickSort ...
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
	pivot := utils.MidOfThree(nums[left], nums[right], nums[mid])
	// fmt.Printf("new quick sort with left: %d, right: %d, pivot: %d\n", left, right, pivot)

	pivotPos := mid
	if nums[left] == pivot {
		pivotPos = left
	} else if nums[right] == pivot {
		pivotPos = right
	}

	pivotPos, nums = partition(nums[left:right+1], pivotPos)

	// fmt.Println("sorted nums:", nums)
	// fmt.Println("will new quick sort: left, pivotPos, right", left, pivotPos, right)

	l := QuickSort(nums[:pivotPos])
	r := QuickSort(nums[pivotPos:])

	return append(l, r...)
}

// 分区操作
func partition(nums []int, pivotPos int) (int, []int) {
	// fmt.Printf("partition pivotPos: %d\n", pivotPos)
	pivot := nums[pivotPos]

	less := []int{}
	greater := []int{}

	for i := 0; i < len(nums); i++ {
		if i == pivotPos {
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

// HeapSort 堆排序
// O(n) = nlog(n),非稳定排序
// reffered to https://www.cnblogs.com/chengxiao/p/6129630.html
// 将给定无序序列构造成一个大顶堆（一般升序采用大顶堆，降序采用小顶堆)。
//
// 用数组表示的（完全二叉树），可以这样表达：
// i的子节点下标为 2*i + 1 和 2 * i + 2. i的父节点下标为 (i-1)/2。
//
//        i
//      /   \
//  2*i+1   2*i+2
//
func HeapSort(arr []int) {
	// 构建大顶堆
	n := len(arr)
	m := n / 2
	for i := m - 1; i >= 0; i-- {
		heap(arr, i, n)
	}

	// 交换堆顶元素到最后一位，并调整堆
	for j := n - 1; j > 0; j-- {
		arr[0], arr[j] = arr[j], arr[0] // 将堆顶元素与末尾元素进行交换
		heap(arr, 0, j)                 // 重新对堆进行调整
	}

	fmt.Println(arr)
}

// 构建大顶堆
func heap(arr []int, i, end int) {
	cur := arr[i]

	// 从i结点的左子结点开始，也就是2i+1处开始
	for k := 2*i + 1; k < end; k = k*2 + 1 {
		if k+1 < end && arr[k] < arr[k+1] {
			// 如果左子结点小于右子结点，k指向右子结点
			k++
		}

		if arr[k] > cur {
			// 如果子节点大于父节点，将元素放到父节点的位置
			arr[i] = arr[k]
			i = k
		} else {
			// 如果子节点小于父节点，则不用处理
			break
		}

		// 将父节点的值放到最终位置
		arr[i] = cur
	}
}
