package __打印二叉树的边界节点

/*
// 标准一
1. 得到二叉树的最左，最右节点
2. 从上到下，打印所有层的左节点
3. 先序遍历，打印那些不属于某一层最左或者最右的节点，但同时又是叶节点的节点
4. 从下到上，打印所有层的右节点
*/

type Node struct {
	Left *Node
	Right *Node
	Val int
}

type Tree struct {
	root *Node
}

func (t *Tree) printEdge1() []int{
	res := make([]int, 0)
	h := t.height()
	edgeMap := t.getEdgeMap(h)
	for l:=0; l<h; l++{
		// 1. 左边界节点
		res = append(res, edgeMap[l][0].Val)
		//fmt.Println(edgeMap[l][0].Val)
	}
	// 2. 既不是左边界 也不是右边界节点
	leaf := t.leafNotEdgMap(edgeMap)
	//fmt.Println("val=", leaf)
	res = append(res, leaf...)
	// 3. 右边界节点
	for l:=h-1; l>0; l--{
		res = append(res, edgeMap[l][1].Val)
		//fmt.Println(edgeMap[l][1].Val)
	}
	return res
}

func (t *Tree) getEdgeMap(h int) [][]*Node{
	// level, 0:l;1:r
	edgeMap := make([][]*Node, h)
	for i:=0; i<h; i++{
		edgeMap[i] = make([]*Node, 2)
	}
	getEdgeMap(t.root, 0, edgeMap)
	return edgeMap
}

func getEdgeMap(node *Node, l int, edgeMap [][]*Node) {
	if node == nil{
		return
	}
	// 左边界节点
	if edgeMap[l][0] == nil{
		edgeMap[l][0] = node
	}
	// 右边界节点
	edgeMap[l][1] = node
	getEdgeMap(node.Left, l+1, edgeMap)
	getEdgeMap(node.Right, l+1, edgeMap)
}

func (t *Tree)leafNotEdgMap(edgeMap [][]*Node) []int{
	res := make([]int, 0)
	leafNodeEdgeMap(t.root, 0, edgeMap, &res)
	return res
}

func leafNodeEdgeMap(node *Node, l int, edgeMap [][]*Node, res *[]int){
	if node == nil{
		return
	}
	// 根节点
	if node.Left == nil && node.Right == nil && node != edgeMap[l][0] && node != edgeMap[l][1]{
		*res = append(*res, node.Val)
	}
	leafNodeEdgeMap(node.Left, l+1, edgeMap, res)
	leafNodeEdgeMap(node.Right, l+1, edgeMap, res)
}

func (t *Tree) height() int{
	return height(t.root)
}

func height(node *Node) int{
	if node == nil{
		return 0
	}
	return max(height(node.Left), height(node.Right)) + 1
}

func max(a, b int) int{
	if a > b{
		return a
	}
	return b
}