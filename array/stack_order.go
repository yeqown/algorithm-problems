package array

// ABCD入栈顺序，一共有几种出栈情况
// https://blog.csdn.net/shikelang_pp/article/details/77170438
// f(4) = f(0)*f(3) + f(1)*f(2) + f(2) * f(1) + f(3)*f(0)
// f(n) = f(0)*f(n-1) + f(1)*f(n-2) + ... + f(n-1)*f(0)

var cache = make(map[int]int)

func init() {
	cache[0] = 1
	cache[1] = 1
	cache[2] = 2
}

// StackOrder ...
func StackOrder(order string) int {
	length := len(order)
	return helper(length)
}
func helper(n int) int {
	if cnt, ok := cache[n]; ok {
		return cnt
	}

	var (
		cnt = 0
	)

	for i := 0; i < n; i++ {
		cnt += (helper(i) * helper(n-i-1))
	}

	cache[n] = cnt
	return cnt
}
