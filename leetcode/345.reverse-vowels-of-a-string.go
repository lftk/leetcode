// https://leetcode.com/problems/reverse-vowels-of-a-string/

package leetcode

func reverseVowels(s string) string {
	b := []byte(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		for l < r && !isVowels(b[l]) {
			l++
		}
		for l < r && !isVowels(b[r]) {
			r--
		}
		b[l], b[r] = b[r], b[l]
	}
	return string(b)
}

func isVowels(c byte) bool {
	for _, v := range "AaEeIiOoUu" {
		if c == byte(v) {
			return true
		}
	}
	return false
}
