package dp

import (
	"testing"
)

func Test_ConstructMatrix(t *testing.T) {
	matr := constructMatrix(6, 4)
	PrintMatrix(matr)
}

func Test_LCS(t *testing.T) {
	var (
		strA string   = "BDCABA"
		strB string   = "ABCBDAB"
		lcs  []string = []string{"BCBA", "BDAB", "BCAB"}
	)

	result_lcs := LCS(strA, strB)

	if result_lcs == "" {
		t.Fatal("didn't got any lcs")
	}

	for _, l := range lcs {
		if l == result_lcs {
			t.Log("got one lcs")
			goto label_success
		}
	}
	t.Fatal("didn't got any lcs")
label_success:
	t.Log("dp lcs pass")
}
