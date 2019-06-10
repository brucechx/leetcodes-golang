package _66_矩阵中的路径

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHasPath(t *testing.T) {
	str := "ABCESFCSADEE"
	var matrix []rune
	for _, v := range str{
		matrix = append(matrix, v)
	}
	assert.Equal(t, true, HasPath(matrix, 3, 4, "ABCCED"))
	assert.Equal(t, true, HasPath(matrix, 3, 4, "SEE"))
	assert.Equal(t, false, HasPath(matrix, 3, 4, "ABCB"))
}
