package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

func searchBST(root *Methods.TreeNode) []int {
	mapNode := make(map[int]int)
	var result []int
	maxFreq := 0

	var traversal func(node *Methods.TreeNode)
	traversal = func(node *Methods.TreeNode) {
		if node == nil {
			return
		}

		traversal(node.Left)

		if freq, ok := mapNode[node.Val]; ok {
			mapNode[node.Val] = freq + 1
		} else {
			mapNode[node.Val] = 1
		}

		traversal(node.Right)
	}

	traversal(root)

	for k, v := range mapNode {
		if v > maxFreq {
			maxFreq = v
			result = []int{k} // Start a new list of modes with this value.
		} else if v == maxFreq {
			result = append(result, k) // Append to the list of modes.
		}
	}

	return result
}
