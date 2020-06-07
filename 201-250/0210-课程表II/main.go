package _210_课程表II

// bfs 拓扑排序
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		// 连接矩阵
		edges = make([][]int, numCourses)
		// 入度列表
		indeg = make([]int, numCourses)
		// 拓扑排序
		result = make([]int, 0)
	)
	// init
	for _, prerequisite := range prerequisites{
		edges[prerequisite[1]] = append(edges[prerequisite[1]], prerequisite[0])
		indeg[prerequisite[0]]++
	}
	// 入度为0的入队列
	queue := make([]int, 0)
	for i:=0; i< numCourses; i++{
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		u := queue[0] // 出队列
		result = append(result, u)
		queue = queue[1:]
		// u 相邻节点，入度减1
		for _, v := range edges[u]{
			indeg[v]--
			if indeg[v] == 0{ // 如果入度为0，入队列
				queue = append(queue, v)
			}
		}
	}
	if len(result) == numCourses{
		return result
	}
	return []int{}
}
