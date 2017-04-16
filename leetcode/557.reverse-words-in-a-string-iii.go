// https://leetcode.com/problems/reverse-words-in-a-string-iii

package leetcode

func reverseWords3(s string) string {
	reverse := func(str string) string {
		n := len(str)
		ss := make([]rune, n)
		for i, s := range str {
			ss[n-i-1] = s
		}
		return string(ss)
	}
	var (
		i  int
		s2 string
	)
	for j, r := range s {
		if r != ' ' {
			continue
		}
		s2 += reverse(s[i:j]) + " "
		i = j + 1
	}
	s2 += reverse(s[i:])
	return s2
}
