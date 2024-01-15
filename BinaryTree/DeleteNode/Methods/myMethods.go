package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

func deleteNode(node *Methods.TreeNode, val int) *Methods.TreeNode {
	if node == nil {
		return node
	}
	//第二种情况：左右孩子都为空（叶子节点），直接删除节点， 返回NULL为根节点
	if node.Right == nil && node.Left == nil && node.Val == val {
		return nil
	}
	//第三种情况：删除节点的左孩子为空，右孩子不为空，删除节点，右孩子补位，返回右孩子为根节点
	if node.Left == nil && node.Val == val {
		return node.Right
	}
	//第四种情况：删除节点的右孩子为空，左孩子不为空，删除节点，左孩子补位，返回左孩子为根节点
	if node.Right == nil && node.Val == val {
		return node.Left
	}
	//第五种情况：左右孩子节点都不为空，则将删除节点的左子树头结点（左孩子）放到删除节点的右子树的最左面节点的左孩子上，返回删除节点右孩子为新的根节点

	var leftnode func(node *Methods.TreeNode) *Methods.TreeNode

	leftnode = func(node *Methods.TreeNode) *Methods.TreeNode {

		for node.Left != nil {
			node = node.Left
		}
		return node
	}

	if node.Val == val {

		inOrderSuccessor := leftnode(node.Right)
		node.Val = inOrderSuccessor.Val

		node.Right = deleteNode(node.Right, inOrderSuccessor.Val)

		return node
	}
	if node.Val > val {
		node.Left = deleteNode(node.Left, val)
	}
	if node.Val < val {
		node.Right = deleteNode(node.Right, val)
	}

	return node
}
