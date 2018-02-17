package tree

import (
	"testing"
)

func Test_TrieTree(t *testing.T) {
	words := []string{"abc", "and", "as", "both", "del", "test"}
	root := NewTrie(words)

	if root == nil {
		t.Fatal("new trie got error")
	}

}
