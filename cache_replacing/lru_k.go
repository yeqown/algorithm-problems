package cachereplacing

import (
	"container/list"
	"fmt"
)

// LRUKCache .
type LRUKCache struct {
	K           uint // the K setting
	rest        uint // max - used = rest
	historyRest uint // historyMax - used = historyRest

	// history doubly linked list
	history *list.List

	// history get op O(1)
	historyHash map[int]*list.Element

	// cache doubly linked list, save
	cache *list.List

	// cache get op O(1)
	cacheHash map[int]*list.Element
}

// NewLRUKCache .
func NewLRUKCache(k uint, capacity uint) *LRUKCache {
	if k <= 1 {
		k = 2 // suggest k(min) = 2
	}

	return &LRUKCache{
		K:           k,
		history:     list.New(),
		cache:       list.New(),
		rest:        capacity,
		historyRest: capacity * ((k % 3) + 1), // TODO: how to limit history max, manually or auto calculatting
		historyHash: make(map[int]*list.Element),
		cacheHash:   make(map[int]*list.Element),
	}
}

// reorderCache cache linked-list to sort (Least Recently Used)
func (c *LRUKCache) reorderCache(ele *list.Element) {
	// rmeove cache data from linked-list
	c.cache.Remove(ele)
	// push it into head
	c.cache.PushFront(ele.Value)
}

func (c *LRUKCache) cacheReplacingCheck() {
	if c.rest == 0 {
		tail := c.cache.Back()
		c.cache.Remove(tail)
		delete(c.cacheHash, tail.Value.(data).key)
		c.rest++
	}
}

func (c *LRUKCache) historyReplacingCheck() {
	if c.historyRest == 0 {
		// true: trigger LRU history
		tail := c.history.Back()
		c.history.Remove(tail)
		delete(c.historyHash, tail.Value.(historyCounter).key)
		c.historyRest++
	}
}

func (c *LRUKCache) delfromHistory(key int) {
	ele, ok := c.historyHash[key]
	if !ok {
		return
	}
	delete(c.historyHash, key)
	c.history.Remove(ele)
	c.historyRest++
}

func (c *LRUKCache) addtoCache(d data) {
	c.cacheReplacingCheck()
	// push into cache
	ele := c.cache.PushFront(d)
	c.cacheHash[d.key] = ele
	c.rest--
}

func (c *LRUKCache) addtoHistory(key, value int) historyCounter {
	var (
		hc historyCounter
	)

	// missed in cache, then check history to incr visited time
	hEle, ok := c.historyHash[key]
	if ok {
		// auto-incr visited count
		hc = hEle.Value.(historyCounter)
		hc.visited++

		if hc.visited >= c.K {
			// true: removed from history, and add into cache
			c.delfromHistory(key)
			c.addtoCache(data{key: key, value: value})
			return hc
		}
		// write back to history
		hEle.Value = hc
	} else {
		// new an history visited log
		c.historyReplacingCheck()
		hc = historyCounter{key: key, visited: 1}
		hEle = c.history.PushFront(hc)
		c.historyRest--
	}

	c.historyHash[key] = hEle

	return hc
}

// Get key of cache in LRUKCache .
func (c *LRUKCache) Get(key int) (value int, ok bool) {
	ele, ok := c.cacheHash[key]
	if ok {
		c.reorderCache(ele)
		return ele.Value.(data).value, true
	}

	// false: missed
	return 0, false
}

// Put of LRUKCache
func (c *LRUKCache) Put(key, value int) {
	ele, ok := c.cacheHash[key]
	if ok {
		// true: hit the key
		c.reorderCache(ele)
		return
	}

	_ = c.addtoHistory(key, value)
	fmt.Printf("cacheRest: %d / cacheMap: %v \nhistoryRest: %d / historyMap: %v\n\n",
		c.rest, c.cacheHash, c.historyRest, c.historyHash)
	return
}
