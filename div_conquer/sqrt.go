package div_conquer

import (
	"math"
)

func Sqrt(n int) int {
	if n == 0 {
		return 0
	}

	var (
		left  = 1
		right = math.MaxInt32
	)

	for left < right {
		mid := left + (right-left)/2

		mid_pow := mid * mid

		// if mid^2 > n
		if mid_pow > n {
			right = mid - 1
		} else if mid_pow < n {
			left = mid + 1
		} else {
			return mid
		}
	}
	if left*left > n {
		return left - 1
	}
	return left
}

func NewtonIterSqrt(n int) float64 {
	var x float64 = 1
	for i := 0; i < 100; i++ {
		x = (x + float64(n)/x) / 2
		// if condition {
		//		break
		// }
	}
	return x
}
