package _063_不同路径_II

import (
	"fmt"
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T){
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}
