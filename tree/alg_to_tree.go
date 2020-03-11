package tree

import (
	"github.com/yeqown/algorithm-problems/utils"
)

/**
 * 计算二叉树的直径
 */
func diameterOfBinaryTree(root *CommonTreeNode) (diameter int) {
	diameter = 1
	// d := max(diameter, root.Left+root.Right+1)
	_ = depth(root, &diameter)
	return diameter - 1
}

func depth(node *CommonTreeNode, diameter *int) int {
	if IsEmptyNode(node) {
		return 0
	}
	if IsLeaf(node) {
		return 1
	}

	dl := depth(node.LeftChild, diameter)
	dr := depth(node.RightChild, diameter)
	*diameter = utils.Max(*diameter, dl+dr+1)

	return dl + dr + 1
}
