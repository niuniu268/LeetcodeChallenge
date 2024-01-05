package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

//102. 二叉树的递归遍历

func levelOrder(root *Methods.TreeNode) [][]int {
	arr := [][]int{}

	depth := 0

	var order func(root *Methods.TreeNode, depth int)

	order = func(root *Methods.TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(arr) == depth {
			arr = append(arr, []int{})
		}
		arr[depth] = append(arr[depth], root.Val)

		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	order(root, depth)

	return arr
}
