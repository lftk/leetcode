// https://leetcode.com/problems/excel-sheet-column-number/

package leetcode

func titleToNumber(s string) int {
	var n int
	for _, c := range s {
		i := int(c-'A') + 1
		n = 26*n + i
	}
	return n
}
