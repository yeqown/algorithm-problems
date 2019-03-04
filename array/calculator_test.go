package array_test

import (
	"github.com/yeqown/algorithm-problems"
	"testing"
)

func Test_ConvMid2Polish(t *testing.T) {
	expr := "1 + ( ( 222 + 43 ) * 4 ) - 45"
	want := "-45,-43,1,-42,-43,222,43,4,45,"
	// want - + 1 × + 2 3 4 5
	if have := ConvMid2Polish(expr); have != want {
		t.Fatalf("want: %s, have: %s", want, have)
	}
}

func Test_ConvMid2Polish_case2(t *testing.T) {
	expr := "1+((222+43)*4)-45"
	want := "-45,-43,1,-42,-43,222,43,4,45,"
	// want - + 1 × + 2 3 4 5
	if have := ConvMid2Polish(expr); have != want {
		t.Fatalf("want: %s, have: %s", want, have)
	}
}

func Test_PolishTonationCalcu(t *testing.T) {
	expr := "1+((222+43)*4)-45"
	want := 1016
	// want - + 1 × + 2 3 4 5
	polish := ConvMid2Polish(expr)
	if have := PolishTonationCalcu(polish); have != want {
		t.Fatalf("Wrong answer, have: %d, want: %d", have, want)
	}
}
