package _119_杨辉三角_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRow(t *testing.T){
	res := []int{1, 3, 3, 1}
	assert.Equal(t, res, getRow(3))
}
