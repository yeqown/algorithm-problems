package tree

// union find set
// refer to:  https://subetter.com/algorithm/union-find-set.html

import "fmt"

const (
	// N of count of array
	N = 10
)

var (
	pre  [N]int
	rank [N]int
)

// InitUnionFindSet .
// use an array [10]
func InitUnionFindSet() {
	for i := 0; i < N; i++ {
		pre[i] = i
		rank[i] = 0
	}
}

// Find . find the root of the node(i)
func Find(i int) int {
	if pre[i] == i {
		return i
	}

	iParent := pre[i]
	iRoot := Find(iParent)

	return iRoot
}

// PathCompressionFind .
func PathCompressionFind(i int) int {
	if pre[i] == i {
		return i
	}

	iParent := pre[i]
	iRoot := Find(iParent)

	pre[i] = iRoot // 路径压缩

	return iRoot
}

// Union . union the node(j) and node(j) into one tree
func Union(i, j int) {
	iRoot := Find(i)
	jRoot := Find(j)

	if iRoot == jRoot {
		return
	}

	pre[iRoot] = jRoot
	fmt.Printf("union %d and %d, got: %v\n", i, j, pre)
}

// UnionByRank ,
func UnionByRank(i, j int) {
	// iRoot := Find(i)
	// jRoot := Find(j)

	// replaced by Path Compression
	iRoot := PathCompressionFind(i)
	jRoot := PathCompressionFind(j)

	if iRoot == jRoot {
		return
	}

	if rank[iRoot] > rank[jRoot] {
		iRoot, jRoot = jRoot, iRoot
	}

	pre[iRoot] = jRoot

	if rank[iRoot] == rank[jRoot] {
		// 两个秩相同的树合并，则整体的秩就会增加 1
		rank[jRoot]++
	}
	fmt.Printf("union %d and %d, got: %v\n", i, j, pre)
}
