package utils

import (
	"testing"
)

var (
	chars = []string{
		"+-*/",
		"()",
		"1234567890",
	}
)

// no more need
func Test_LogNeedChar(t *testing.T) {
	cs := "09+-*/()"

	for _, c := range cs {
		t.Logf("Char=%s,Code=%d", string(c), c)
	}
}

func Test_ConvRuneNum2Int(t *testing.T) {
	t.Log("want 1234567890")
	for _, c := range chars[2] {
		t.Log(ConvRuneNum2Int(c))
	}
}
