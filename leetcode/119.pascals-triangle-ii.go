// https://leetcode.com/problems/pascals-triangle-ii/

package leetcode

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	row := append(getRow(rowIndex-1), 1)
	for i := rowIndex - 1; i > 0; i-- {
		row[i] += row[i-1]
	}
	return row
}
