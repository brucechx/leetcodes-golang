package __打印二叉树的边界节点

import (
	"fmt"
	"testing"
)

func TestPrintEdge1(t *testing.T){
	tree := mockTree()
	fmt.Println(tree.printEdge1())
}

func mockTree() *Tree{
	return &Tree{root:&Node{
		Left:  &Node{
			Left:  nil,
			Right: &Node{
				Left:  &Node{
					Left:  nil,
					Right: nil,
					Val:   7,
				},
				Right: &Node{
					Left:  nil,
					Right: &Node{
						Left:  &Node{
							Left:  nil,
							Right: nil,
							Val:   13,
						},
						Right: &Node{
							Left:  nil,
							Right: nil,
							Val:   14,
						},
						Val:   11,
					},
					Val:   8,
				},
				Val:   4,
			},
			Val:   2,
		},
		Right: &Node{
			Left:  &Node{
				Left:  &Node{
					Left:  &Node{
						Left:  &Node{
							Left:  nil,
							Right: nil,
							Val:   15,
						},
						Right: &Node{
							Left:  nil,
							Right: nil,
							Val:   16,
						},
						Val:   12,
					},
					Right: nil,
					Val:   9,
				},
				Right: &Node{
					Left:  nil,
					Right: nil,
					Val:   10,
				},
				Val:   5,
			},
			Right: &Node{
				Left:  nil,
				Right: nil,
				Val:   6,
			},
			Val:   3,
		},
		Val:   1,
	}}

}
