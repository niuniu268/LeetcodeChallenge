package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

func insertIntoBST(node *Methods.TreeNode, val int) *Methods.TreeNode {

	if node == nil {
		return &Methods.TreeNode{Val: val}
	}

	var traversal func(node *Methods.TreeNode, val int)

	traversal = func(node *Methods.TreeNode, val int) {
		if node == nil {

			NewNode := &Methods.TreeNode{
				Val: val,
			}

			if val > node.Val {
				node.Right = NewNode
			} else {
				node.Left = NewNode
			}
			return
		}

		if node.Val > val {
			traversal(node.Left, val)
		}
		if node.Val < val {
			traversal(node.Right, val)
		}

		return

	}

	traversal(node, val)

	return node
}
