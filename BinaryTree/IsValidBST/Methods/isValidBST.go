package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

// 中序遍历解法
func isValidBST(root *Methods.TreeNode) bool {
	// 保存上一个指针
	var prev *Methods.TreeNode
	var travel func(node *Methods.TreeNode) bool
	travel = func(node *Methods.TreeNode) bool {
		if node == nil {
			return true
		}
		leftRes := travel(node.Left)
		// 当前值小于等于前一个节点的值，返回false
		if prev != nil && node.Val <= prev.Val {
			return false
		}
		prev = node
		rightRes := travel(node.Right)
		return leftRes && rightRes
	}
	return travel(root)
}
