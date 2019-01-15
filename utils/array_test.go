package utils

import (
	"testing"
)

func Test_MidOfThree(t *testing.T) {
	if mid := MidOfThree(8, 11, 31); mid != 11 {
		t.Fatal("wrong mid")
	}

	if mid := MidOfThree(-8, -2, -23); mid != -8 {
		t.Fatal("wrong mid")
	}

	if mid := MidOfThree(-8, 2, -23); mid != -8 {
		t.Fatal("wrong mid")
	}
}
