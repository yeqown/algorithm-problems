package array

import (
	"testing"
)

func Test_BinarySearch(t *testing.T) {
	var (
		arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	)

	// left target
	pos := BinarySearch(arr, 0, len(arr)-1, 4)
	if pos != 3 {
		t.Fatalf("find target %d-[%d], actual is: %d", 99, pos, MISS)
	}

	// miss target
	pos = BinarySearch(arr, 0, len(arr)-1, 99)
	if pos != MISS {
		t.Fatalf("find target %d-[%d], actual is: %d", 99, pos, MISS)
	}

	// right target
	pos = BinarySearch(arr, 0, len(arr)-1, 10)
	if pos != 9 {
		t.Fatalf("find target %d-[%d], actual is: %d", 10, pos, 9)
	}
}

func TestJumpSearch(t *testing.T) {
	var (
		arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	)

	// left target
	pos := JumpSearch(arr, 4)
	if pos != 3 {
		t.Fatalf("find target %d-[%d], actual is: %d", 99, pos, MISS)
	}

	// miss target
	pos = JumpSearch(arr, 99)
	if pos != MISS {
		t.Fatalf("find target %d-[%d], actual is: %d", 99, pos, MISS)
	}

	// right target
	pos = JumpSearch(arr, 10)
	if pos != 9 {
		t.Fatalf("find target %d-[%d], actual is: %d", 10, pos, 9)
	}
}
