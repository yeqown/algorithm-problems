package div_conquer

import (
	"testing"
)

func Test_Sqrt(t *testing.T) {
	if root := Sqrt(8); root != 2 {
		t.Fatal("not excepted answer")
	}

	if root := Sqrt(1); root != 1 {
		t.Fatal("not excepted answer")
	}
}

func Test_NewtonIterSqrt(t *testing.T) {
	if root := NewtonIterSqrt(4); root != 2.00 {
		t.Fatal("not excepted answer")
	}
}
