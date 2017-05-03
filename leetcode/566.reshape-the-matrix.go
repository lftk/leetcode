// https://leetcode.com/problems/reshape-the-matrix

package leetcode

func matrixReshape(nums [][]int, r int, c int) [][]int {
	l := len(nums[0])
	if l*len(nums) < r*c {
		return nums
	}
	newNums := make([][]int, r)
	for i := 0; i < r; i++ {
		newNums[i] = make([]int, c)
		for j := 0; j < c; j++ {
			n := i*c + j
			i1, j1 := n/l, n%l
			newNums[i][j] = nums[i1][j1]
		}
	}
	return newNums
}
