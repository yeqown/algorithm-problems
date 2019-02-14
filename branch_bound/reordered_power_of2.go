package branch_bound

/*
Starting with a positive integer N, we reorder the digits in any order (including the original order) such that the leading digit is not zero.
Return true if and only if we can do this in a way such that the resulting number is a power of 2.

Example 1:
	Input: 1
	Output: true

Example 2:
	Input: 10
	Output: false

Example 3:
	Input: 16
	Output: true

Example 4:
	Input: 24
	Output: false

Example 5:
	Input: 46
	Output: true
*/

// 存放描述一个数的特征，数的位数
type numNode struct {
	nums map[int]int // map[2]2 表示数字2有2位
	len  int         // 10 表示node标志的数一共有10位数
}

func (n numNode) equal(node numNode) bool {
	if n.len != node.len {
		return false
	}

	for k, v := range n.nums {
		if _, ok := node.nums[k]; !ok {
			return false
		}

		if node.nums[k] != v {
			return false
		}
	}

	return true
}

var (
	// 2^29 = 0,536,870,912
	// 2^30 = 1,073,741,824
	// 10^9 = 1,000,000,000

	powOf2 map[int][]numNode // map[1]4 表示一位数的2^n特征列表
)

func convert(n int) numNode {
	var (
		mod  int
		node = numNode{
			nums: make(map[int]int),
		}
	)

	for {
		mod = n % 10
		n = n / 10
		node.nums[mod]++
		node.len++
		if n == 0 {
			break
		}
	}
	return node
}

func init() {
	powOf2 = make(map[int][]numNode)
	var n = 1
	powOf2[1] = append(powOf2[1], numNode{len: 1, nums: map[int]int{1: 1}})
	for i := 1; i <= 29; i++ {
		n = n << 1
		node := convert(n)
		powOf2[node.len] = append(powOf2[node.len], node)
	}
}

func reorderedPowerOf2(N int) bool {
	target := convert(N)
	if nodes, ok := powOf2[target.len]; ok {
		for _, n := range nodes {
			if n.equal(target) {
				return true
			}
		}
	}
	return false
}
