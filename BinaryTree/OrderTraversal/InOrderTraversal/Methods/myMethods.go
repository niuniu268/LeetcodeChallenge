package Methods

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Traversal(cur *TreeNode, val []int) {
	if cur == nil {
		return
	}
	Traversal(cur.Left, val)
	val = append(val, cur.Val)
	Traversal(cur.Right, val)
}

func InOrderTraversal(root *TreeNode) []int {

	result := make([]int, 0)
	Traversal(root, result)

	return result
}
