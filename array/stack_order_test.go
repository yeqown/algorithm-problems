package array_test

import (
	"github.com/yeqown/algorithm-problems/array"

	"testing"
)

func TestStackOrder(t *testing.T) {
	if n := stackOrder("ABCDF"); n != 12 {
		t.Errorf("wrong stackOrder result")
	}
}
