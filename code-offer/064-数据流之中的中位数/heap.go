package _64_数据流之中的中位数

import "sort"

type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}

// 入堆
func Push(h Interface, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1)
}

// 由子节点至父节点依次重建堆
func up(h Interface, j int){
	for {
		i := (j - 1) /2 // 父
		if i == j || !h.Less(j, i){
			break
		}
		h.Swap(i, j)
		j = i
	}
}

//下滤：由父节点至子节点依次建堆
//parent      : i
//left child  : 2*i+1
//right child : 2*i+2
func down(h Interface, i0, n int) bool{
	//构建最小堆,父小于两子节点值,最大堆同理，视Less函数实现
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		//找出两个节点中最小的(less: a<b)
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1){
			j = j2
		}
		//然后与父节点i比较,如果父节点小于这个子节点最小值,则break,否则swap
		if !h.Less(j, i){
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}


// Pop removes the minimum element (according to Less) from the heap
// and returns it. The complexity is O(log(n)) where n = h.Len().
// It is equivalent to Remove(h, 0).
//
func Pop(h Interface) interface{}{
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

