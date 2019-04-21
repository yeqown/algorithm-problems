// Copyright 2019 The yeqown Author. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// bst.go

package tree

import (
	"errors"
)

var (
	errNullRoot = errors.New("root is NULL")
)

// BSTNode ... node contains left & right child pointer and value or key
type BSTNode struct {
	// val of current node contains
	Val int

	// left child node
	Left *BSTNode

	// right child node
	Right *BSTNode
}

// IsLeaf ... judge a BSTNode is a leaf or not
func (node *BSTNode) IsLeaf() bool {
	return node.Left == nil && node.Right == nil
}

// BSTree ... binary search tree
type BSTree *BSTNode

// Search one target element in the tree
func Search(root *BSTNode, target int) *BSTNode {
	if root == nil {
		return nil
	}

	// fmt.Println(root.Val, target)
	if root.Val == target {
		return root
	}

	// search left sub-BST
	if target < root.Val {
		if left := Search(root.Left, target); left != nil {
			return left
		}
	} else {
		if right := Search(root.Right, target); right != nil {
			return right
		}
	}

	return nil
}

// Insert a element into the BST
func Insert(root *BSTNode, val int) *BSTNode {
	if root == nil {
		root = new(BSTNode)
		root.Val = val
		return root
	}

	if val < root.Val {
		root.Left = Insert(root.Left, val)
	} else if val > root.Val {
		root.Right = Insert(root.Right, val)
	}

	return root
}

// Inorder ... of BSTree
func Inorder(root *BSTNode) []int {
	output := make([]int, 0)
	if root != nil {
		output = append(output, Inorder(root.Left)...)
		output = append(output, root.Val)
		output = append(output, Inorder(root.Right)...)
	}

	return output
}
