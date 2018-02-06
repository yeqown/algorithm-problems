package dp

import (
	"testing"
)

func Test_PrintMatrix(t *testing.T) {
	m := Matrix{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}
	PrintMatrix(m)
}

func Test_LD_ConstructMatrix(t *testing.T) {
	ld := &LD{Rows: 3, Cols: 6}
	ld.constructMatrix()

	if len(ld.M[0]) != 7 {
		t.Fatal("ConstructMatrix make a wrong matrix with cols")
	}
	if len(ld.M) != 4 {
		t.Fatal("ConstructMatrix make a wrong matrix with rows")
	}
	// display
	PrintMatrix(ld.M)
}

func Test_LevenshteinDistance(t *testing.T) {
	source := "GUMBO"
	dest := "GAMBOL"

	dis := LevenshteinDistance(source, dest)

	if dis != 2 {
		t.Fatalf("wrong dis is got, %d, actual: %d", dis, 2)
	}
}
