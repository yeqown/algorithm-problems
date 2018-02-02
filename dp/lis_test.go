package dp

import (
	"testing"
)

type QA struct {
	Input    []int
	Excepted []int
}

func isEqualArray(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Test_LIS(t *testing.T) {
	// qa := &QA{
	// 	Input:    []int{2, 2, 2, 2, 2, 2},
	// 	Excepted: []int{2},
	// }
	// ans := LIS(qa.Input)
	// if !isEqualArray(ans, qa.Excepted) {
	// 	t.Fatal("not excepted answer")
	// }

	qa := &QA{
		Input:    []int{1, 2, 4, 5, 7, 6, 8, 4, 3, 2, 1, 10, 910},
		Excepted: []int{1, 2, 4, 5, 7, 8, 10, 910},
	}

	ans := LIS(qa.Input)
	if !isEqualArray(ans, qa.Excepted) {
		t.Fatal("not excepted answer")
	}
}
