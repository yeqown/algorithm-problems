// 分执法解决数组最大值问题
package div_conquer

import (
	"math"
	_ "net/http/pprof"
)

// 一般方法求数组最大值，complexty: O-time(n) O-space(1)
func NormalMaxOfArray(arr []int, length int) int {
	max := math.MinInt64
	for i := 0; i < length; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// 分治法求数组最大值，并使用并发来提升时间效率
// channel 来同步结果
// (原始版)
func DivConquerMaxOfArray(arr []int, length int, max chan int) {
	// 只有一个元素那么久返回自己
	if length == 1 {
		max <- arr[0]
		return
	}
	var chanLeftMax chan int = make(chan int, 1)
	var chanRightMax chan int = make(chan int, 1)

	mid := int(length / 2)
	// 左边
	go DivConquerMaxOfArray(arr[:mid], mid, chanLeftMax)
	// 右边
	go DivConquerMaxOfArray(arr[mid:], length-mid, chanRightMax)

	clm := <-chanLeftMax
	crm := <-chanRightMax

	if clm > crm {
		max <- clm
		return
	}
	max <- crm
}

// 分治法求数组最大值，并使用并发来提升时间效率
// channel 来同步结果
// (改进版本1)
var minl int

func SetMinLength(v_minl int) {
	minl = v_minl
}

func DivConquerMaxOfArray_Mend(arr []int, length int, max chan int) {
	// 只有一个元素那么久返回自己
	if length <= minl {
		max <- NormalMaxOfArray(arr, length)
		return
	}
	var chanLeftMax chan int = make(chan int, 1)
	var chanRightMax chan int = make(chan int, 1)
	defer close(chanLeftMax)
	defer close(chanRightMax)

	mid := int(length / 2)
	// 左边
	go DivConquerMaxOfArray(arr[:mid], mid, chanLeftMax)
	// 右边
	go DivConquerMaxOfArray(arr[mid:], length-mid, chanRightMax)

	clm := <-chanLeftMax
	crm := <-chanRightMax

	if clm > crm {
		max <- clm
		return
	}
	max <- crm
	return
}

// 并发求最大值
func NormalMaxOfArrayWithChan(arr []int, length int, max chan int) {
	tmpMax := math.MinInt64
	for i := 0; i < length; i++ {
		if arr[i] > tmpMax {
			tmpMax = arr[i]
		}
	}
	max <- tmpMax
	return
}

func ConcurrenceMaxOfArray(arr []int, length int) int {
	mid := int(length / 2)
	var left chan int = make(chan int, 1)
	var right chan int = make(chan int, 1)
	defer close(left)
	defer close(right)

	go NormalMaxOfArrayWithChan(arr[:mid], mid, left)
	go NormalMaxOfArrayWithChan(arr[mid:], length-mid, right)

	clm := <-left
	crm := <-right

	if clm > crm {
		return clm
	}
	return crm
}
