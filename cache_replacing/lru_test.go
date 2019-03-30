package cachereplacing_test

import (
	cp "github.com/yeqown/algorithm-problems/cache_replacing"
	"testing"
)

// ["LRUCache","put","get","put","get","get"]
// [[1],[2,1],[2],[3,2],[2],[3]]
func TestLRUCache(t *testing.T) {
	lru := cp.Constructor(1)
	lru.Put(2, 1)
	lru.Get(2)
	lru.Put(3, 2)
	lru.Get(2)
	lru.Get(3)
}
