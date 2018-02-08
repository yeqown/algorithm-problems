package div_conquer

import (
	"testing"
)

func Test_MaxSubArray(t *testing.T) {
	arr := []int{1, 2, 3, 45, 9}
	excepted := 60

	if max_sub := MaxSubArray(arr, 0, len(arr)-1); max_sub != excepted {
		t.Fatal("want: ", excepted, "have: ", max_sub)
	}
}

func Test_MaxSubArray_case1(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	excepted := 6

	if max_sub := MaxSubArray(arr, 0, len(arr)-1); max_sub != excepted {
		t.Fatal("want: ", excepted, "have: ", max_sub)
	}
}

func Test_MaxOfThree(t *testing.T) {
	var a, b, c int = 1, 2, 3
	if max := MaxOfThree(a, b, c); max != c {
		t.Fatal("Wrong Judge with MaxOfThree")
	}
}
