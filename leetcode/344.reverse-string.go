// https://leetcode.com/problems/reverse-string/

package leetcode

func reverseString(s string) string {
	var ss []byte
	if n := len(s); n > 0 {
		ss = make([]byte, n)
		for i := 0; i < n; i++ {
			ss[i] = s[n-i-1]
		}
	}
	return string(ss)
}
