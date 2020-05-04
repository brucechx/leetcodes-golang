package _515_在每个树行中找最大值

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T){
	tree := mockTree()
	expectRes := []int{1, 3, 9}
	assert.Equal(t, expectRes, largestValues(tree))
}

func mockTree() *TreeNode{
	return &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   3,
			Left:  &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}
}
