package Methods

import (
	`Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods`
)

var pre int

func convertBST(root *Methods.TreeNode) *Methods.TreeNode {
	pre = 0
	traversal(root)
	return root
}

func traversal(cur *Methods.TreeNode) {
	if cur == nil {
		return
	}
	traversal(cur.Right)
	cur.Val += pre
	pre = cur.Val
	traversal(cur.Left)
}
