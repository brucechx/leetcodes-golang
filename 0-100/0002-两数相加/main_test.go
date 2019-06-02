package _002_两数相加

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	cases := []struct {
		x *ListNode
		y *ListNode
		expect string
	}{
		{createList([]int{2, 4, 3}), createList([]int{5, 6, 4}), "7 --> 0 --> 8"},
		{createList([]int{2, 4, 3, 1}), createList([]int{5, 6, 4}), "7 --> 0 --> 8 --> 1"},
	}

	for _, cas := range cases{
		res := AddTwoNumbers(cas.x, cas.y)
		assert.Equal(t, cas.expect, res.String() )
	}
}