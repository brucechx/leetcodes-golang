package _18

import "testing"

func TestIdenticalTrees(t *testing.T) {
	left := &Node{
		nil,
		nil,
		1,
	}
	right := &Node{
		nil,
		nil,
		2,
	}
	root := &Node{
		left,
		right,
		0,
	}

	t.Run("IdenticalTrees", func(t *testing.T) {
		got := IdenticalTrees(root, root)
		if !got {
			t.Errorf("%v sholuld be true", got)
		}
	})

	tree2 := &Node{
		nil,
		nil,
		5,
	}
	t.Run("Not identicalTrees", func(t *testing.T) {
		got := IdenticalTrees(root, tree2)
		if got {
			t.Errorf("%v sholuld be false", got)
		}
	})
}

