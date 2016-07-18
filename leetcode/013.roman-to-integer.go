// https://leetcode.com/problems/roman-to-integer/

package leetcode

func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1, 'X': 10, 'C': 100, 'M': 1000, 'V': 5, 'L': 50, 'D': 500,
	}
	var n int
	for i := 0; i < len(s); i++ {
		if j := i + 1; j < len(s) && roman[s[i]] < roman[s[j]] {
			n += roman[s[j]] - roman[s[i]]
			i++
		} else {
			n += roman[s[i]]
		}
	}
	return n
}
