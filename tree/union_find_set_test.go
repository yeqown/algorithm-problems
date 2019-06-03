package tree_test

import (
	"testing"

	"github.com/yeqown/algorithm-problems/tree"
)

func Test_UnionFindSet(t *testing.T) {
	tree.InitUnionFindSet()

	tree.Union(2, 9)
	tree.Union(4, 9)
	tree.Union(3, 4)
	tree.Union(5, 6)

	if v1, v2 := tree.Find(3), tree.Find(5); v1 == v2 {
		t.Logf("3 is connected to 5, but should not")
		t.Fail()
	}

	if v1, v2 := tree.Find(2), tree.Find(9); v1 != v2 {
		t.Logf("2[%d] is not connected to 9[%d]", v1, v2)
		t.Fail()
	}
}

func TestUnionFindSetOptmize(t *testing.T) {
	tree.InitUnionFindSet()

	tree.UnionByRank(2, 9)
	tree.UnionByRank(4, 9)
	tree.UnionByRank(3, 4)
	tree.UnionByRank(5, 6)

	if v1, v2 := tree.PathCompressionFind(3), tree.PathCompressionFind(5); v1 == v2 {
		t.Logf("3 is connected to 5, but should not")
		t.Fail()
	}

	if v1, v2 := tree.PathCompressionFind(2), tree.PathCompressionFind(9); v1 != v2 {
		t.Logf("2[%d] is not connected to 9[%d]", v1, v2)
		t.Fail()
	}
}
