package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

func lowestCommonAncestor(root, p, q *Methods.TreeNode) *Methods.TreeNode {
	if root == nil {
		return nil
	}
	for {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		}
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		}
		if (root.Val-p.Val)*(root.Val-q.Val) <= 0 {
			return root
		}
	}
	return root
}
