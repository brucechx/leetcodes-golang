package _052_N皇后_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalNQueens(t *testing.T)  {
	cases := []struct{
		input int
		res int
	}{
		{
			input: 4,
			res: 2,
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, totalNQueens(cas.input))
	}
}
