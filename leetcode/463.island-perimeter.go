// https://leetcode.com/problems/island-perimeter/

package leetcode

func islandPerimeter(grid [][]int) int {
	var perimeter int
	for i, line := range grid {
		for j, val := range line {
			if val == 0 {
				continue
			}
			perimeter += 4
			if i > 0 && grid[i-1][j] == 1 {
				perimeter--
			}
			if i < len(grid)-1 && grid[i+1][j] == 1 {
				perimeter--
			}
			if j > 0 && grid[i][j-1] == 1 {
				perimeter--
			}
			if j < len(line)-1 && grid[i][j+1] == 1 {
				perimeter--
			}
		}
	}
	return perimeter
}
