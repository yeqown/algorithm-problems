package array

import (
	"testing"
)

func Test_Jump2(t *testing.T) {
	arr := []int{2, 3, 1, 1, 4}
	excepted_step := 2
	if step := Jump2(arr); step != excepted_step {
		t.Fatal("Wrong jump alg with step:", step)
	}
}

func Test_Jump2_case2(t *testing.T) {
	arr := []int{2, 1}
	excepted_step := 1
	if step := Jump2(arr); step != excepted_step {
		t.Fatal("Wrong jump alg with step:", step)
	}
}

func Test_Jump2_case3(t *testing.T) {
	arr := []int{1, 1, 1, 1}
	excepted_step := 3
	if step := Jump2(arr); step != excepted_step {
		t.Fatal("Wrong jump alg with step:", step)
	}
}

func Test_Jump2_case4(t *testing.T) {
	arr := []int{1, 2, 1, 1, 1}
	excepted_step := 3
	if step := Jump2(arr); step != excepted_step {
		t.Fatal("Wrong jump alg with step:", step)
	}
}

func Test_Jump2_case5(t *testing.T) {
	arr := []int{1, 2, 3}
	excepted_step := 2
	if step := Jump2(arr); step != excepted_step {
		t.Fatal("Wrong jump alg with step:", step)
	}
}

func Test_Jump2_case6(t *testing.T) {
	arr := []int{4, 1, 1, 3, 1, 1, 1}
	excepted_step := 2
	if step := Jump2(arr); step != excepted_step {
		t.Fatal("Wrong jump alg with step:", step)
	}
}
