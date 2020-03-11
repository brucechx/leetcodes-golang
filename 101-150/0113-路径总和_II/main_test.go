package _113_路径总和_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyPathSum(t *testing.T) {
	cases := []struct{
		root *TreeNode
		target int
		res [][]int
	}{
		{
			root: mockTree(),
			target: 22,
			res:  [][]int{
				{5, 4, 11, 2},
				{5, 8, 4, 5},
			},
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, MyPathSum(cas.root, cas.target))
	}
}

func mockTree() *TreeNode{
	return &TreeNode{
		Val:   5,
		Left:  &TreeNode{
			Val:   4,
			Left:  &TreeNode{
				Val:   11,
				Left:  &TreeNode{
					Val:   7,
				},
				Right: &TreeNode{
					Val:   2,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   8,
			Left:  &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val:5},
				Right: &TreeNode{Val:1},
			},
		},
	}
}