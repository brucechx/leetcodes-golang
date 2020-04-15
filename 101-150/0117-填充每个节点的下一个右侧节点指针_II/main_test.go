package _117_填充每个节点的下一个右侧节点指针_II

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T){
	root := &Node{
		Val:   1,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}
	res := connect(root)
	fmt.Println(res)
}
