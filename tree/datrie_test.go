package tree

import (
	"testing"
)

func TestTrieTree(t *testing.T) {
	words := []string{"abc", "and", "as", "both", "del", "test"}
	root := NewTrie(words...)

	if root == nil {
		t.Fatal("new trie got error")
	}

	type args struct {
		prefix []byte
	}
	type testcase struct {
		name    string
		args    args
		wantCnt int
		wantOK  bool
	}

	cases := []testcase{
		{
			name: "case 1",
			args: args{
				prefix: []byte("a"),
			},
			wantCnt: 3,
			wantOK:  true,
		},
		{
			name: "case 2",
			args: args{
				prefix: []byte("g"),
			},
			wantCnt: 0,
			wantOK:  false,
		},
		{
			name: "case 3",
			args: args{
				prefix: []byte("test"),
			},
			wantCnt: 1,
			wantOK:  true,
		},
	}

	for _, tt := range cases {
		cnt, ok := root.SearchPrefix(tt.args.prefix)
		if cnt != tt.wantCnt || ok != tt.wantOK {
			t.Errorf("%s root.SearchPrefix(%s) want=(%d, %v), got(%d, %v)",
				tt.name, tt.args.prefix, tt.wantCnt, tt.wantOK, cnt, ok)
			t.FailNow()
		}
	}

	root.Delete([]byte("as"))
	cnt, ok := root.SearchPrefix([]byte("a"))
	if cnt != 2 || ok != true {
		t.Errorf("root.SearchPrefix(%s) want=(%d, %v), got(%d, %v)",
			"a", 2, true, cnt, ok)
		t.FailNow()
	}
	cnt, ok = root.SearchPrefix([]byte("as"))
	if cnt != 0 || ok != false {
		t.Errorf("root.SearchPrefix(%s) want=(%d, %v), got(%d, %v)",
			"as", 0, false, cnt, ok)
		t.FailNow()
	}
}
