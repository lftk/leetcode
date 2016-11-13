// https://leetcode.com/problems/longest-palindrome/

package leetcode

func longestPalindrome(s string) int {
	ss := make(map[rune]int)
	for _, c := range s {
		ss[c]++
	}

	var nn, odd int
	for _, n := range ss {
		if n%2 == 0 {
			nn += n
		} else {
			nn += n - 1
			odd = 1
		}
	}
	return nn + odd
}
