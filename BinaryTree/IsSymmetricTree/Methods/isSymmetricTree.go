package Methods

import (
	Methods2 "Leetcode/BinaryTree/BuildTree/Methods"
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

func Compare(left, right *Methods.TreeNode) bool {

	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return Compare(left.Left, right.Right) && Compare(right.Left, left.Right)

}

func IsSymmetric1(root *Methods.TreeNode) bool {

	return Compare(root.Left, root.Right)
}

func isSymmetric2(root *Methods2.TreeNode) bool {
	var queue []*Methods2.TreeNode
	if root != nil {
		queue = append(queue, root.Left, root.Right)
	}
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, right.Left, left.Right)
	}
	return true
}
