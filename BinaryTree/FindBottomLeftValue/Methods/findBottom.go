package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

var depth int // 全局变量 最大深度
var res int   // 记录最终结果
func findBottomLeftValue2(root *Methods.TreeNode) int {
	depth, res = 0, 0 // 初始化
	dfs(root, 1)
	return res
}

func dfs(root *Methods.TreeNode, d int) {
	if root == nil {
		return
	}
	// 因为先遍历左边，所以左边如果有值，右边的同层不会更新结果
	if root.Left == nil && root.Right == nil && depth < d {
		depth = d
		res = root.Val
	}
	dfs(root.Left, d+1) // 隐藏回溯
	dfs(root.Right, d+1)
}
