package tree_test

import (
	"reflect"
	"testing"

	"github.com/yeqown/algorithm-problems/tree"
)

func TestBSTree(t *testing.T) {
	input := []int{1, 23, 6, 12, 44, 67, 32, 88}

	var root *tree.BSTNode
	// build the tree
	for _, v := range input {
		root = tree.Insert(root, v)
	}

	output := tree.Inorder(root)
	want := []int{1, 6, 12, 23, 32, 44, 67, 88}
	if !reflect.DeepEqual(want, output) {
		t.Errorf("want: %v, got: %v", want, output)
	}

	for _, v := range input {
		if node := tree.Search(root, v); node == nil {
			t.Errorf("could not found: %d", v)
			continue
		}
	}
}
