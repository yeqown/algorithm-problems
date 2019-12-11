package array

import (
	"math/rand"
	"strings"
	"time"
)

// skiplistLevel .
type skiplistLevel struct {
	forward *skiplistNode // forward pointer
	span    uint          // span
}

// skiplistNode .
type skiplistNode struct {
	levels   []*skiplistLevel // level datastructure to contains all levels of one node.
	backward *skiplistNode    // backward pointer to the last node.
	score    float32          // score to sort by, score could be duplicated
	key      string           // the key object also be the value.
}

func newskiplistNode(score float32, key string, level int) *skiplistNode {
	node := &skiplistNode{
		levels:   make([]*skiplistLevel, level),
		backward: nil,
		score:    score,
		key:      key,
	}

	return node
}

func (node *skiplistNode) level() int {
	return len(node.levels)
}

// comparekey return 1 if node.key is bigger than target, 0 means equal
func (node *skiplistNode) comparekey(targetKey string) int {
	return strings.Compare(node.key, targetKey)
}

// Skiplist .
// references are following:
//
// http://zhangtielei.com/posts/blog-redis-skiplist.html
// https://www.cnblogs.com/xudong-bupt/p/7468860.html
//
// Skiplist based Sorted Linked List, it contains random algorithm.
// 1. it contains levels which speeds up searching an item;
// 2. each level contains pointer to anothor node;
// 3. the levels[1] to traverse all items;
// 4. all levels in one node share the same key object;
type Skiplist struct {
	head, tail *skiplistNode // pointers to head and tail
	length     uint          // the length of skiplist
	level      int           // the maxnium level number in all nodes

	p        float32    // p, to control the average pointer used space with maxLevel
	maxLevel int        // the limitation of level
	r        *rand.Rand // rand.Rand
}

// NewSkiplist default params: maxLevel = 32, p = 0.25
func NewSkiplist(maxLevel int, p float32) *Skiplist {
	list := &Skiplist{
		length: 0,
		level:  1,

		maxLevel: maxLevel,
		p:        p,
		r:        rand.New(rand.NewSource(time.Now().UnixNano())), // rand seed
	}

	// init head node
	list.head = newskiplistNode(0, "head", maxLevel)
	for i := 0; i < list.maxLevel; i++ {
		list.head.levels[i] = &skiplistLevel{
			forward: nil,
			span:    0,
		}
	}
	list.head.backward = nil
	list.tail = nil

	return list
}

// Insert . insert and update levels
// assume key not exists
func (list *Skiplist) Insert(score float32, key string, level int) {
	update := make([]*skiplistNode, list.maxLevel)
	spans := make([]uint, list.maxLevel)

	visitNode := list.head
	for i := list.level - 1; i >= 0; i-- {
		if i != list.level-1 {
			spans[i] = spans[i+1]
		}

		// to find the position of current level
		for visitNode.levels[i].forward != nil &&
			(visitNode.levels[i].forward.score < score ||
				(visitNode.levels[i].forward.score == score && visitNode.comparekey(key) < 0)) {

			spans[i] += visitNode.levels[i].span
			visitNode = visitNode.levels[i].forward
		}

		update[i] = visitNode
	}

	if level == 0 {
		level = list.randomLevel()
	}

	if level > list.level {
		// true: level is bigger than current level max
		for i := list.level; i < level; i++ {
			spans[i] = 0
			update[i] = list.head
			update[i].levels[i].span = list.length
		}

		list.level = level
	}

	// create node
	insertNode := newskiplistNode(score, key, level)

	// update forward pointer
	for i := 0; i < level; i++ {
		insertNode.levels[i] = new(skiplistLevel)
		insertNode.levels[i].forward = update[i].levels[i].forward
		update[i].levels[i].forward = insertNode

		insertNode.levels[i].span = update[i].levels[i].span - (spans[0] - spans[i])
		update[i].levels[i].span = (spans[0] - spans[i]) + 1
	}

	// increment span for untouched levels
	for i := level; i < list.level; i++ {
		update[i].levels[i].span++
	}

	if update[0] == list.head {
		insertNode.backward = nil
	} else {
		insertNode.backward = update[0]
	}

	if insertNode.levels[0].forward != nil {
		insertNode.levels[0].forward.backward = insertNode
	} else {
		list.tail = insertNode
	}

	list.length++
}

// Remove .
// TODO:
func (list *Skiplist) Remove(key string) {

}

// Range .
// TODO:
func (list *Skiplist) Range(lowScore, hightScore float32) {

}

// RevRange . reverse range items
// TODO:
func (list *Skiplist) RevRange(lowScore, hightScore float32) {

}

// Rank of the key
// TODO:
func (list *Skiplist) Rank(key interface{}) int {
	return -1
}

// randomLevel depends on P and MaxLevel
func (list *Skiplist) randomLevel() int {
	level := 1

	for list.r.Float32() < list.p && level < list.maxLevel {
		level++
	}

	return level
}
