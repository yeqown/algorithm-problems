package cachereplacing_test

import (
	"testing"

	cp "github.com/yeqown/algorithm-problems/cache_replacing"
)

// ["LRUCache","put","get","put","get","get"]
// [[1],[2,1],[2],[3,2],[2],[3]]
func TestLRUKCache(t *testing.T) {
	cache := cp.NewLRUKCache(2, 2)
	cache.Put(2, 1)
	if _, hit := cache.Get(2); hit {
		t.Error("could not get 2 hit")
	}
	cache.Put(3, 1)
	if _, hit := cache.Get(3); hit {
		t.Error("could not get 3 hit")
	}
	cache.Put(3, 1)
	if _, hit := cache.Get(3); !hit {
		t.Error("should get 3 hit")
	}
	cache.Put(4, 1)
	cache.Put(4, 1)
	if _, hit := cache.Get(2); hit {
		t.Error("could not get 2 hit")
	}
	if _, hit := cache.Get(4); !hit {
		t.Error("should get 2 hit")
	}
	if _, hit := cache.Get(3); !hit {
		t.Error("should get 3 hit")
	}
}
