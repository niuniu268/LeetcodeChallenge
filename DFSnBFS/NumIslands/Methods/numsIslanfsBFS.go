package Methods

import (
	"container/list"
	"fmt"
)

type Solution struct {
	visited [][]bool
	move    [][]int
}

func (s *Solution) numIslands(grid [][]byte) int {
	res := 0
	s.visited = make([][]bool, len(grid))
	for i := range s.visited {
		s.visited[i] = make([]bool, len(grid[0]))
	}

	s.move = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !s.visited[i][j] && grid[i][j] == '1' {
				s.bfs(grid, i, j)
				res++
			}
		}
	}
	return res
}

func (s *Solution) bfs(grid [][]byte, y, x int) {
	queue := list.New()
	queue.PushBack([]int{y, x})
	s.visited[y][x] = true

	for queue.Len() > 0 {
		cur := queue.Front().Value.([]int)
		queue.Remove(queue.Front())

		m, n := cur[0], cur[1]
		for i := 0; i < 4; i++ {
			nexty := m + s.move[i][0]
			nextx := n + s.move[i][1]

			if nextx < 0 || nexty == len(grid) || nexty < 0 || nextx == len(grid[0]) {
				continue
			}

			if !s.visited[nexty][nextx] && grid[nexty][nextx] == '1' {
				queue.PushBack([]int{nexty, nextx})
				s.visited[nexty][nextx] = true
			}
		}
	}
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	solution := Solution{}
	result := solution.numIslands(grid)
	fmt.Println(result)
}
