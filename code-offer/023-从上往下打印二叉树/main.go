package _23_从上往下打印二叉树

import "fmt"

type Node struct {
	Left *Node
	Right *Node
	Data Item
}

func PrintFromTopToBottom(root *Node){
	var q Queue
	if root == nil{
		return
	}
	q.Enqueue(root)
	for !q.isEmpty(){
		t := q.Dequeue()
		node := t.(*Node)
		fmt.Println(node.Data)
		if node.Left != nil{
			q.Enqueue(node.Left)
		}
		if node.Right != nil{
			q.Enqueue(node.Right)
		}
	}
}

