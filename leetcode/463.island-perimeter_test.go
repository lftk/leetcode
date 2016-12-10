package leetcode

import (
	"testing"
)

func Test_islandPerimeter(t *testing.T) {
	grid := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	num := islandPerimeter(grid)
	if num != 16 {
		t.Error(grid, 16, num)
	}
}
