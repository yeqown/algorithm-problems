package tree

// CommonTreeNode .
type CommonTreeNode struct {
	Value      int
	LeftChild  *CommonTreeNode
	RightChild *CommonTreeNode
}

// IsLeaf .
func IsLeaf(node *CommonTreeNode) bool {
	return node.LeftChild == nil && node.RightChild == nil
}

// IsEmptyNode .
func IsEmptyNode(node *CommonTreeNode) bool {
	return node == nil
}

// CommonTree .
type CommonTree struct {
	Root *CommonTreeNode
}
