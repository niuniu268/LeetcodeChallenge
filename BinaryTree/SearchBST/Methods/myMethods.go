package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

func searchBST(root *Methods.TreeNode, value int) *Methods.TreeNode {

	if root == nil || root.Val == value {
		return root
	}

	if root.Val > value {
		return searchBST(root.Left, value)
	} else if root.Val < value {
		return searchBST(root.Right, value)
	}

	return nil
}
