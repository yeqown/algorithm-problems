package array

import (
	// "github.com/yeqown/algorithm-problems/array"

	"testing"
)

func Test_StackOrder(t *testing.T) {
	if n := StackOrder("ABCD"); n != 14 {
		t.Errorf("wrong stackOrder result, want: %d, got: %d", 14, n)
	}
}
