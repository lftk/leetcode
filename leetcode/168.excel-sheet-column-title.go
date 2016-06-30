// https://leetcode.com/problems/excel-sheet-column-title/

package leetcode

func convertToTitle(n int) string {
	const l string = "ZABCDEFGHIJKLMNOPQRSTUVWXY"
	var s string
	for n > 0 {
		i := n % 26
		n = (n - 1) / 26
		s = string(l[i]) + s
	}
	return s
}
