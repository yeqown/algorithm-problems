package div_conquer

/*
Given a non-negative integer c, your task is to decide whether there're two integers a and b such that a2 + b2 = c.

Example 1:
	Input: 5
	Output: True
	Explanation: 1 * 1 + 2 * 2 = 5

Example 2:
	Input: 3
	Output: False
*/
func judgeSquareSum(c int) bool {
	for a := 0; a*a <= c; a++ {
		b := c - a*a
		if binarySearch(0, b, b) {
			return true
		}
	}
	return false
}

func binarySearch(s, e, target int) bool {
	if s > e {
		return false
	}
	mid := s + (e-s)/2
	if mid*mid == target {
		return true
	}
	if mid*mid > target {
		return binarySearch(s, mid-1, target)
	}
	return binarySearch(mid+1, e, target)
}
