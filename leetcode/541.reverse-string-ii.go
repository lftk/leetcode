// https://leetcode.com/problems/reverse-string-ii

package leetcode

func reverseStr(s string, k int) string {
	f := func(s string) string {
		var s2 string
		for i := len(s) - 1; i >= 0; i-- {
			s2 += string(s[i])
		}
		return s2
	}

	var (
		i  int
		s2 string
	)
	for i < len(s) {
		if i+k > len(s) {
			s2 += f(s[i:])
			break
		}
		s2 += f(s[i : i+k])
		i += k

		if i+k > len(s) {
			s2 += s[i:]
			break
		}
		s2 += s[i : i+k]
		i += k
	}
	return s2
}
