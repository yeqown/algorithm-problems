// 最长公共子序列

/*
* 如果str1的所有字符按其在字符串中的顺序出现在另外一个str2中，
* 则str1称之为str2的子串。注意，并不要求子串（str1）的字符必须连续出现在str2中。
* 如： abcd 是 a1b2c3d4的子串
* 请编写一个函数，输入两个字符串，求它们的最长公共子串，并打印出最长公共子串。
 */

/* 解法
* 令 X=<x1,x2,...,xm> 和 Y=<y1,y2,...,yn> 为两个序列，Z=<z1,z2,z3,...,zk>为X和Y的任意LCS。则
* 如果xm = yn，则 zk=xm=yn 且Zk−1是Xm−1和Yn−1的一个LCS。
* 如果xm ≠ yn，那么zk ≠ xm，意味着Z是Xm−1和Y的一个LCS。
* 如果xm ≠ yn，那么zk ≠ yn，意味着Z是X和Yn−1的一个LCS。
*
*              |- 0                             (i = 0 | j = 0)
*              |
*   LCS[i,j] = |- LCS[i-1, j-1]                 (i, j > 0 & Str1[i] = Str2[j])
*              |
*              |- MAX(LCS[i, j-1], LCS[i-1, j]) (j, j> 0 & Str1[i] ≠ Str2[j])
 */

package dp

import (
	"fmt"
)

type Matrix [][]int

func constructMatrix(rows, cols int) Matrix {
	m := make([][]int, rows+1)
	for i := 0; i <= rows; i++ {
		m[i] = make([]int, cols+1)
	}
	return m
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

func MaxOfTwo(a, b int) int {
	max := a
	if b > max {
		max = b
	}
	return max
}

func PrintLCS(m Matrix, str string, i, j int) {
	if i == 0 || j == 0 {
		return
	}
	if m[i][j] == m[i-1][j] {
		PrintLCS(m, str, i-1, j)
	} else if m[i][j] == m[i][j-1] {
		PrintLCS(m, str, i, j-1)
	} else {
		PrintLCS(m, str, i-1, j-1)
		fmt.Print(string(str[i-1]), " ")
		lcs = append(lcs, str[i-1])
	}
}

var lcs []byte = []byte{}

func LCS(strA, strB string) string {
	var (
		l_stra int = len(strA)
		l_strb int = len(strB)
	)

	// matrix := constructMatrix(rows, cols)
	matrix := constructMatrix(l_stra, l_strb)

	for i := 1; i <= l_stra; i++ {
		for j := 1; j <= l_strb; j++ {
			if strA[i-1] == strB[j-1] {
				matrix[i][j] = matrix[i-1][j-1] + 1
			} else {
				matrix[i][j] = MaxOfTwo(matrix[i-1][j], matrix[i][j-1])
			}
		}
	}

	PrintLCS(matrix, strA, l_stra, l_strb)
	fmt.Println()

	return string(lcs)
}
