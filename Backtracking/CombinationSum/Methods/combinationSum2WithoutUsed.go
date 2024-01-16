package Methods

import (
	"sort"
)

var (
// res [][]int
// path  []int
)

func combinationSum3(candidates []int, target int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, len(candidates))
	sort.Ints(candidates) // 排序，为剪枝做准备
	dfs2(candidates, 0, target)
	return res
}

func dfs2(candidates []int, start int, target int) {
	if target == 0 { // target 不断减小，如果为0说明达到了目标值
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target { // 剪枝，提前返回
			break
		}
		// i != start 限制了这不对深度遍历到达的此值去重
		if i != start && candidates[i] == candidates[i-1] { // 去重
			continue
		}
		path = append(path, candidates[i])
		dfs(candidates, i+1, target-candidates[i])
		path = path[:len(path)-1]
	}
}
