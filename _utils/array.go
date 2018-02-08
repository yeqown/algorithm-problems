// 数组常用操作

package _utils

func EqualArray(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for pos, val := range arr1 {
		if arr2[pos] != val {
			return false
		}
	}
	return true
}

func MaxOfArray(arr []int) (max int) {
	max = arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func MidOfArrary(arr []int) (mid int) {
	return arr[0]
}
