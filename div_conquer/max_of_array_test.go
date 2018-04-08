package div_conquer

import (
	"math/rand"
	"testing"
	"time"
)

// 创建一个随机数组
func creatArray(length int) []int {
	arr := make([]int, length)

	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(length)
	}
	return arr
}

var arr10w []int
var arr100w []int

func init() {
	// 创建10w int数组
	arr10w = creatArray(100000)
	// 创建 100w int 数组
	arr100w = creatArray(1000000)
}

func Test_createArray(t *testing.T) {
	t.Log(creatArray(10))
}

func Test_NormalMaxOfArray(t *testing.T) {
	start := time.Now()
	max10w := NormalMaxOfArray(arr10w, 100000)
	end := time.Now()
	elapsed_10w := end.Sub(start)

	start = time.Now()
	max100w := NormalMaxOfArray(arr100w, 1000000)
	end = time.Now()
	elapsed_100w := end.Sub(start)

	t.Logf("time cost of NormalMaxOfArray: 10w-[%s] 100w-[%s], max result is: 10w-[%d], 100w-[%d]\n",
		elapsed_10w, elapsed_100w, max10w, max100w)
}

func Test_DivConquerMaxOfArray(t *testing.T) {
	var max10w chan int = make(chan int, 1)
	var max100w chan int = make(chan int, 1)
	// 设置10w量级最小长度为10
	start := time.Now()
	DivConquerMaxOfArray(arr10w, 100000, max10w)
	end := time.Now()
	elapsed_10w := end.Sub(start)

	start = time.Now()
	DivConquerMaxOfArray(arr100w, 1000000, max100w)
	end = time.Now()
	elapsed_100w := end.Sub(start)

	t.Logf("time cost of DivConquerMaxOfArray: 10w-[%s] 100w-[%s], max result is: 10w-[%d], 100w-[%d]\n",
		elapsed_10w, elapsed_100w, <-max10w, <-max100w)
}

// 并发改进，一定量级采用普通进行求最大值，在量级的数据时候跟一般方法一致
func Test_DivConquerMaxOfArray_Mend(t *testing.T) {
	var max10w chan int = make(chan int, 1)
	var max100w chan int = make(chan int, 1)
	// 设置10w量级最小长度为10
	SetMinLength(50000)
	start := time.Now()
	DivConquerMaxOfArray_Mend(arr10w, 100000, max10w)
	end := time.Now()
	elapsed_10w := end.Sub(start)

	start = time.Now()
	SetMinLength(2500000)
	DivConquerMaxOfArray_Mend(arr100w, 1000000, max100w)
	end = time.Now()
	elapsed_100w := end.Sub(start)

	t.Logf("time cost of DivConquerMaxOfArray_Mend: 10w-[%s] 100w-[%s], max result is: 10w-[%d], 100w-[%d]\n",
		elapsed_10w, elapsed_100w, <-max10w, <-max100w)
}

/*
 *	测试结果是第二种并发的效率远低于最常用的求最大值
 *  可能原因：创建太多的协程导致效率低下；由于channel wait。
 *  思考下改进的方案：粗分数组（3-5等分），采用并发（一般求最大值的方式），且在一定量级才采用并发的方式。
 */

// 粗分数组二等分，明显提升速度，但是不能进行细粒度的划分，否则会出现最开始分治法解决问题的结果，效率感而下降。
func Test_ConcurrenceMaxOfArray(t *testing.T) {
	start := time.Now()
	max10w := ConcurrenceMaxOfArray(arr10w, 100000)
	end := time.Now()
	elapsed_10w := end.Sub(start)

	start = time.Now()
	max100w := ConcurrenceMaxOfArray(arr100w, 1000000)
	end = time.Now()
	elapsed_100w := end.Sub(start)

	t.Logf("time cost of ConcurrenceMaxOfArray: 10w-[%s] 100w-[%s], max result is: 10w-[%d], 100w-[%d]\n",
		elapsed_10w, elapsed_100w, max10w, max100w)
}
