package _210_课程表II

// dfs 拓扑排序
func findOrder2(numCourses int, prerequisites [][]int) []int {
	var (
		edges = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result []int
		invalid bool
		dfs func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1 // 已搜索
		for _, v := range edges[u] {
			if visited[v] == 0 { // 未搜索
				dfs(v)
				if invalid {
					return
				}
			} else if visited[v] == 1 {
				invalid = true
				return
			}
		}
		visited[u] = 2 // 所有相邻节点均搜索完毕
		result = append(result, u)
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && !invalid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if invalid {
		return []int{}
	}
	for i := 0; i < len(result)/2; i ++ {
		result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i]
	}
	return result
}
