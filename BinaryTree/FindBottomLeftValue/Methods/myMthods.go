package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

func findBottomLeftValue(root *Methods.TreeNode) int {
	depth := 0
	leftdepth, leftval := LeftTraversal(root, depth, root.Val)
	rightdepth, righttval := RightTraversal(root, depth, root.Val)

	if leftdepth >= rightdepth {
		return leftval
	}

	return righttval

}
func LeftTraversal(node *Methods.TreeNode, depth int, value int) (int, int) {
	if node.Left == nil && node.Right == nil {
		return depth, value
	}
	depth, value = LeftTraversal(node.Left, depth+1, node.Left.Val)

	return depth, value
}
func RightTraversal(node *Methods.TreeNode, depth int, value int) (int, int) {
	if node.Left == nil && node.Right == nil {
		return depth, value
	}
	depth, value = RightTraversal(node.Right, depth+1, node.Left.Val)

	return depth, value
}
