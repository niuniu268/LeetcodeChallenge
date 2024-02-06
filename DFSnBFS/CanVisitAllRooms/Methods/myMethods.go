package Methods

func canVisitAllRooms(rooms [][]int) bool {
	visit := make([]bool, len(rooms))

	var dfs func(rooms [][]int, key int, visit []bool)

	dfs = func(rooms [][]int, key int, visit []bool) {
		if visit[key] == true {
			return
		}
		visit[key] = true
		keys := rooms[key]

		for _, elem := range keys {
			dfs(rooms, elem, visit)
		}

	}

	dfs(rooms, 0, visit)

	for _, elem := range visit {
		if elem != true {
			return false
		}
	}

	return true
}
