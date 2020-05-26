package _207_课程表

// bfs + 拓扑排序
func canFinish(numCourses int, prerequisites [][]int) bool {
	indegrees := make([]int, numCourses)
	adjacency := make([][]int, numCourses)
	// 排序队列
	queue := make([]int, 0)
	// 初始化入度和链接矩阵
	for _, prerequisite := range prerequisites{
		curr := prerequisite[0]
		pre := prerequisite[1]
		indegrees[curr] += 1
		adjacency[pre] = append(adjacency[pre], curr)
	}
	// 入度为0的点
	for i, indegree := range indegrees{
		if indegree == 0 {
			queue = append(queue, i)
		}
	}
	//BFS TopSort.
	for len(queue) > 0 {
		pre := queue[0]
		queue = queue[1:]
		numCourses -= 1
		for _, curr := range adjacency[pre]{
			indegrees[curr] -= 1
			if indegrees[curr] == 0{
				queue = append(queue, curr)
			}
		}
	}
	return numCourses == 0
}