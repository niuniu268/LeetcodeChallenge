package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
	"container/list"
)

func countNodes1(root *Methods.TreeNode) int {
	if root == nil {
		return 0
	}
	leftH, rightH := 0, 0
	leftNode := root.Left
	rightNode := root.Right
	for leftNode != nil {
		leftNode = leftNode.Left
		leftH++
	}
	for rightNode != nil {
		rightNode = rightNode.Right
		rightH++
	}
	if leftH == rightH {
		return (2 << leftH) - 1
	}
	return countNodes1(root.Left) + countNodes1(root.Right) + 1
}

func countNodes2(root *Methods.TreeNode) int {
	if root == nil {
		return 0
	}
	q := list.New()
	q.PushBack(root)
	res := 0
	for q.Len() > 0 {
		n := q.Len()
		for i := 0; i < n; i++ {
			node := q.Remove(q.Front()).(*Methods.TreeNode)
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
			res++
		}
	}
	return res
}
