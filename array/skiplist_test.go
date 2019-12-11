package array_test

import (
	"testing"

	"github.com/yeqown/algorithm-problems/array"
)

func Test_Skiplist(t *testing.T) {
	list := array.NewSkiplist(32, 0.25) // new skiplist
	// list.ToDotGraphviz()             // output: skiplist.dot

	list.Insert(0.4, "alice", 0)
	list.Insert(0.5, "john", 0)
	list.Insert(0.2, "smith", 0)
	list.Insert(0.3, "mary", 0)
	list.Insert(0.23, "mary2", 0)

	if err := list.ToDotGraphviz(""); err != nil {
		t.Error(err)
	}
}

func Test_Skiplist_Graphviz(t *testing.T) {
	list := array.NewSkiplist(32, 0.25)

	if err := list.ToDotGraphviz(""); err != nil {
		t.Error(err)
	}
}
