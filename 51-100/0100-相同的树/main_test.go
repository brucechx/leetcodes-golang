package _100_相同的树

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameTree(t *testing.T){
	cases := []struct{
		p *TreeNode
		q *TreeNode
		res bool
	}{
		{
			p: mockTree1(),
			q: mockTree1(),
			res:true,
		},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, isSameTree(cas.p, cas.q))
	}
}


func mockTree1( ) *TreeNode{
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	return root
}
