package array

import (
	"testing"
)

func Test_ConvMid2Polish(t *testing.T) {
	expr := "(1+2x3)-89"
	// want - x + 1 2 3 89
	s := ConvMid2Polish(expr)

}
