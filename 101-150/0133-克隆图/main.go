package _133_克隆图

// dfs
/*
// 值为原图中节点，键为新图中节点
假设所有节点个数为n，
时间复杂度O(n)，每个节点处理一次，栈调用时间复杂度O(H),H为图的最大深度，综合复杂度O(n)
空间复杂度O(n)，哈希表需要O(n)，栈需要O(H)
 */

type Node struct {
 	Val int
 	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	var dfs func(n *Node) *Node
	dfs = func(n *Node) *Node{
		if n == nil{
			return nil
		}
		if v, ok := visited[n]; ok{
			return v
		}
		v := &Node{
			Val:       n.Val,
			Neighbors: make([]*Node, len(n.Neighbors)),
		}
		visited[n] = v
		for i, val := range n.Neighbors{
			v.Neighbors[i] = dfs(val)
		}
		return v
	}
	return dfs(node)
}
