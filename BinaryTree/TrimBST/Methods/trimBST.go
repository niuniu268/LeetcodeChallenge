package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

// 递归
func trimBST(root *Methods.TreeNode, low int, high int) *Methods.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low { //如果该节点值小于最小值，则该节点更换为该节点的右节点值，继续遍历
		right := trimBST(root.Right, low, high)
		return right
	}
	if root.Val > high { //如果该节点的值大于最大值，则该节点更换为该节点的左节点值，继续遍历
		left := trimBST(root.Left, low, high)
		return left
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
