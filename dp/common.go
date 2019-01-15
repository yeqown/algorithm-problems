package dp

import "fmt"

// Matrix ...
type Matrix [][]int

// PrintMatrix ...
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
