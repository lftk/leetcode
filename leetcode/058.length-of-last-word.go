// https://leetcode.com/problems/length-of-last-word/

package leetcode

func lengthOfLastWord(s string) int {
	var n1, n2 int
	for _, c := range s {
		if c != ' ' {
			n2++
		} else if n2 != 0 {
			n1, n2 = n2, 0
		}
	}
	if n2 == 0 {
		n2 = n1
	}
	return n2
}
