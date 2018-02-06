// Levenshtein Distance

/**** explain ***
* LD is measure of the similarity between two strings,
* which we will refer to as the source string (s) and the target string (t).
* The distance is the number of deletions, insertions,
* or substitutions required to transform s into t.
 */

/* example,
* If s is "test" and t is "test", then LD(s,t) = 0, because no transformations are needed. The strings are already identical.
* If s is "test" and t is "tent", then LD(s,t) = 1, because one substitution (change "s" to "n") is sufficient to transform s into t.
 */

/* used
* Spell checking
* Speech recognition
* DNA analysis
* Plagiarism detection
 */

// Step	Description
// 1	Set n to be the length of s;Set m to be the length of t;If n = 0, return m and exit;If m = 0, return n and exit;Construct a matrix containing 0..m rows and 0..n columns.
// 2	Initialize the first row to 0-n; Initialize the first column to 0-m.
// 3	Examine each character of s (i from 1 to n).
// 4	Examine each character of t (j from 1 to m).
// 5	If s[i] equals t[j], the cost is 0; If s[i] doesn't equal t[j], the cost is 1.
// 6	Set cell d[i,j] of the matrix equal to the minimum of:
// 		a. The cell immediately above plus 1: d[i-1,j] + 1.
// 		b. The cell immediately to the left plus 1: d[i,j-1] + 1.
// 		c. The cell diagonally above and to the left plus the cost: d[i-1,j-1] + cost.
// 7	After the iteration steps (3, 4, 5, 6) are complete, the distance is found in cell d[n,m].

package dp

import (
	"fmt"
	// "math"
)

type Matrix [][]int

type LD struct {
	M          Matrix
	Rows, Cols int
}

func (ld *LD) constructMatrix() {

	ld.M = make([][]int, ld.Rows+1)

	for i := 0; i <= ld.Rows; i++ {
		ld.M[i] = make([]int, ld.Cols+1)
		if i == 0 {
			for j := 0; j < ld.Cols+1; j++ {
				ld.M[0][j] = j
			}
		}
		ld.M[i][0] = i
	}
}

func (ld *LD) setMatrix(cost, row, col int) {
	ld.M[row][col] = cost
}

func (ld *LD) getMatrix(cost, row, col int) int {
	return ld.M[row][col]
}

func PrintMatrix(m Matrix) {
	fmt.Print("col: \t ")
	for i := 0; i < len(m[0]); i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()

	for i := 0; i < len(m); i++ {
		fmt.Println("row: ", i, m[i])
	}
}

func minOfThree(a, b, c int) (min int) {
	min = a
	if min > b {
		min = b
	}
	if min > c {
		min = c
	}
	return
}

func LevenshteinDistance(source, dest string) int {
	var cols, rows int = len(source), len(dest)
	if cols == 0 {
		return rows
	}
	if rows == 0 {
		return cols
	}
	var ld *LD = &LD{Rows: rows, Cols: cols}
	ld.constructMatrix()
	// PrintMatrix(ld.M)

	// step 5
	for c := 1; c <= cols; c++ {
		for r := 1; r <= rows; r++ {
			var cur_cost int = 1

			if source[c-1] == dest[r-1] {
				cur_cost = 0
			}
			// step 6
			cost := minOfThree(ld.M[r-1][c-1]+cur_cost, ld.M[r-1][c]+1, ld.M[r][c-1]+1)
			// step 7
			ld.setMatrix(cost, r, c)
		}
	}

	PrintMatrix(ld.M)

	return ld.M[rows][cols]
}
