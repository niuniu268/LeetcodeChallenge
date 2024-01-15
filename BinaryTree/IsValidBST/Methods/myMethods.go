package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

var list []int

func Traversal(node *Methods.TreeNode) {
	if node == nil {
		return
	}
	Traversal(node.Left)
	list = append(list, node.Val)
	Traversal(node.Right)
}

func IsValidBST(root *Methods.TreeNode) bool {
	Traversal(root)
	for i := 1; i < len(list); i++ {
		if list[i-1] > list[i] {
			return false
		}
	}
	return true
}
