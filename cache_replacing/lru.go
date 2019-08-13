package cachereplacing

import (
	"container/list"
	"fmt"
)

/*
Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

Follow up:
Could you do both operations in O(1) time complexity?

Example:

LRUCache cache = new LRUCache( 2 capacity );

	cache.put(1, 1);
	cache.put(2, 2);
	cache.get(1);       // returns 1
	cache.put(3, 3);    // evicts key 2
	cache.get(2);       // returns -1 (not found)
	cache.put(4, 4);    // evicts key 1
	cache.get(1);       // returns -1 (not found)
	cache.get(3);       // returns 3
	cache.get(4);       // returns 4
*/

// https://blog.csdn.net/elricboa/article/details/78847305 LRU的几种实现
// https://blog.csdn.net/wangsifu2009/article/details/6757352 几种页面置换算法

// LRUCache ...
type LRUCache struct {
	recoder map[int]*list.Element // key hit for get O(1)
	linked  *list.List            // linked-list
	rest    int
}

// Constructor ....
func Constructor(capacity int) LRUCache {
	c := LRUCache{
		rest:    capacity,
		linked:  list.New(),
		recoder: make(map[int]*list.Element),
	}

	return c
}

// Get ... O(1) wanted
func (c *LRUCache) Get(key int) int {
	v, ok := c.recoder[key]
	if !ok {
		return -1
	}
	// hit the key
	_ = c.linked.Remove(v)
	c.recoder[key] = c.linked.PushFront(v.Value)
	fmt.Println("get", key, c.linked.Len(), c.recoder)

	return v.Value.(*data).value
}

// Put ... O(1) wanted
func (c *LRUCache) Put(key int, value int) {
	if v, ok := c.recoder[key]; ok {
		c.linked.Remove(v)
		goto h
	}

	if c.rest == 0 {
		tail := c.linked.Back()
		c.linked.Remove(tail)
		delete(c.recoder, tail.Value.(*data).key)
	} else {
		c.rest--
	}
h:
	head := c.linked.PushFront(&data{key, value})
	c.recoder[key] = head
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
