// https://leetcode.com/problems/count-and-say/

package leetcode

func countAndSay(n int) string {
	val := "1"
	for i := 1; i < n; i++ {
		var (
			c   int
			v   byte
			str string
		)
		for j := 0; j < len(val); j++ {
			if val[j] == v {
				c++
				continue
			}
			if c != 0 {
				str += string(c+'0') + string(v)
			}
			c, v = 1, val[j]
		}
		val = str + string(c+'0') + string(v)
	}
	return val
}
