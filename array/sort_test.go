package array

import (
	. "alg/_utils"
	"testing"
)

var (
	arr      = []int{93, 454, 3463, 8, 1, 287, 9, 11, 2, 0, 9, 31}
	excepted = []int{0, 1, 2, 8, 9, 9, 11, 31, 93, 287, 454, 3463}
)

func Test_BubbleSort(t *testing.T) {
	if sorted := BubbleSort(arr); !EqualArray(sorted, excepted) {
		t.Fatal("want: ", excepted, "have: ", sorted)
	}
}

func Test_DirectSelectSort(t *testing.T) {
	if sorted := DirectSelectSort(arr); !EqualArray(sorted, excepted) {
		t.Fatal("want: ", excepted, "have: ", sorted)
	}
}

func Test_InsertSort(t *testing.T) {
	if sorted := InsertSort(arr); !EqualArray(sorted, excepted) {
		t.Fatal("want: ", excepted, "have: ", sorted)
	}
}

func Test_MergeSort(t *testing.T) {
	if sorted := MergeSort(arr, 0, len(arr)-1); !EqualArray(sorted, excepted) {
		t.Fatal("want: ", excepted, "have: ", sorted)
	}
}

func Test_QuickSort(t *testing.T) {
	if sorted := QuickSort(arr); !EqualArray(sorted, excepted) {
		t.Fatal("want: ", excepted, "have: ", sorted)
	}
}

// func Test_QuickSort2(t *testing.T) {
// 	if sorted := QuickSort2(arr); !EqualArray(sorted, excepted) {
// 		t.Fatal("want: ", excepted, "have: ", sorted)
// 	}
// }
