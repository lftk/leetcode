// https://leetcode.com/problems/pascals-triangle/

package leetcode

func generate(numRows int) [][]int {
	var row []int
	var ret [][]int
	for i := 1; i <= numRows; i++ {
		for j := i - 1; j >= 2; j-- {
			row[j-1] += row[j-2]
		}
		row = append(row, 1)
		val := make([]int, len(row))
		copy(val, row)
		ret = append(ret, val)
	}
	return ret
}
