// trie 前缀树
// http://www.cnblogs.com/huangxincheng/archive/2012/11/25/2788268.html

// 前缀树有如下特点：
// 	1. root节点不包含字符
// 	2. 从根节点到某一节点，路径上经过的字符连接起来，就是该节点对应的字符串。
// 	3. 每个单词的公共前缀作为一个字符节点保存。
// 使用场景：
// 	1. 词频统计（节省内存）
// 	2. 前缀匹配（高效）

package tree

// trieNode ...
type trieNode struct {
	childrenNodes [27]*trieNode
	cnt           int
}

// NewTrie ...
func NewTrie(words ...string) *trieNode {
	root := &trieNode{
		childrenNodes: [27]*trieNode{},
		cnt:           0,
	}

	for _, word := range words {
		root.Add([]byte(word))
	}

	return root
}

func (d *trieNode) offset(c byte) int {
	return int(c - 'a')
}

// Add ...
func (d *trieNode) Add(word []byte) {
	d.cnt++
	if len(word) == 0 {
		return
	}

	// 继续操作孩子
	offset := word[0] - 'a'
	if d.childrenNodes[offset] == nil {
		// 未被设置
		d.childrenNodes[offset] = &trieNode{
			childrenNodes: [27]*trieNode{},
			cnt:           0,
		}
	}

	d.childrenNodes[offset].Add(word[1:])
}

// Delete ...
func (d *trieNode) Delete(word []byte) {
	d.cnt--
	if d.cnt == 0 {
		// 没有其他引用则删除该节点
		d = nil
		return
	}

	if len(word) == 0 {
		// 已经删除到最后一个字符
		return
	}

	// 继续操作孩子
	offset := word[0] - 'a'
	if d.childrenNodes[offset] == nil {
		// 不存在的字符
		return
	}

	d.childrenNodes[offset].Delete(word[1:])
}

// Search ...
func (d *trieNode) SearchPrefix(prefix []byte) (int, bool) {
	n := len(prefix)
	if n == 0 {
		if d.cnt == 0 {
			return d.cnt, false
		}

		return d.cnt, true
	}

	offset := d.offset(prefix[0])
	if d.childrenNodes[offset] == nil {
		return 0, false
	}

	return d.childrenNodes[offset].SearchPrefix(prefix[1:])
}
