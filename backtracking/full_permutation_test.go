package backtracking_test

import (
	"testing"

	"github.com/yeqown/algorithm-problems/backtracking"
)

func Test_FullPermutation(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	res := backtracking.FullPermutation(arr)
	for _, v := range res {
		t.Logf("%+v", v)
	}
}
