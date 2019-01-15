package tree

import (
	"testing"
)

func xTestTrieTree(t *testing.T) {
	words := []string{"abc", "and", "as", "both", "del", "test"}
	root := NewDATrie(words...)

	if root == nil {
		t.Fatal("new trie got error")
	}

}
