package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

func deleteNodeTrue(root *Methods.TreeNode, key int) *Methods.TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	if root.Right == nil {
		return root.Left
	}
	if root.Left == nil {
		return root.Right
	}
	minnode := root.Right
	for minnode.Left != nil {
		minnode = minnode.Left
	}
	root.Val = minnode.Val
	root.Right = deleteNode1(root.Right)
	return root
}

func deleteNode1(root *Methods.TreeNode) *Methods.TreeNode {
	if root.Left == nil {
		pRight := root.Right
		root.Right = nil
		return pRight
	}
	root.Left = deleteNode1(root.Left)
	return root
}

// 迭代版本
func deleteOneNode(target *Methods.TreeNode) *Methods.TreeNode {
	if target == nil {
		return target
	}
	if target.Right == nil {
		return target.Left
	}
	cur := target.Right
	for cur.Left != nil {
		cur = cur.Left
	}
	cur.Left = target.Left
	return target.Right
}
func deleteNode2(root *Methods.TreeNode, key int) *Methods.TreeNode {
	// 特殊情况处理
	if root == nil {
		return root
	}
	cur := root
	var pre *Methods.TreeNode
	for cur != nil {
		if cur.Val == key {
			break
		}
		pre = cur
		if cur.Val > key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if pre == nil {
		return deleteOneNode(cur)
	}
	// pre 要知道是删除左孩子还有右孩子
	if pre.Left != nil && pre.Left.Val == key {
		pre.Left = deleteOneNode(cur)
	}
	if pre.Right != nil && pre.Right.Val == key {
		pre.Right = deleteOneNode(cur)
	}
	return root
}
