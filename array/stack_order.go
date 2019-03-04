package array

// ABCD入栈顺序，一共有几种出栈情况

func stackOrder(order string) int {
	if len(order) == 1 {
		return 2
	}
	return stackOrder(order[:1]) + stackOrder(order[1:])
}
