// https://leetcode.com/problems/detect-capital

package leetcode

func detectCapitalUse(word string) bool {
	var n int
	for _, w := range word {
		if w >= 'A' && w <= 'Z' {
			n++
		}
	}
	return (n == 0 || n == len(word) || (n == 1 && word[0] >= 'A' && word[0] <= 'Z'))
}
