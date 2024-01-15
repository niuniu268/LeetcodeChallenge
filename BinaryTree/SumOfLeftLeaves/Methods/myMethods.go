package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

func sumOfLeftLeave(root *Methods.TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 0
	}

	leftValue := sumOfLeftLeave(root.Left)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftValue = root.Left.Val
	}

	rightValue := sumOfLeftLeave(root.Right)
	return leftValue + rightValue
}
