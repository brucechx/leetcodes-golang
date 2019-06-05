package _40_数组中只出现一次的数字

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindNumsAppearOnce(t *testing.T){
	cases := []struct{
		input []int
		res1, res2 int
	}{
		{[]int{2, 4, 3, 6, 3, 2, 5, 5}, 4, 6},
	}

	for _, cas := range cases{
		num1, num2 := FindNumsAppearOnce(cas.input)
		assert.Equal(t, cas.res1, num1)
		assert.Equal(t, cas.res2, num2)
	}
}

