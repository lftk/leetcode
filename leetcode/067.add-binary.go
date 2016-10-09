// https://leetcode.com/problems/add-binary/

package leetcode

func addBinary(a string, b string) string {
	var (
		s string
		c byte
	)
	an, bn := len(a)-1, len(b)-1
	for an >= 0 && bn >= 0 {
		ac, bc := a[an], b[bn]
		cc := ac - '0' + bc - '0' + c
		s, c = string(cc%2+'0')+s, cc/2
		an, bn = an-1, bn-1
	}
	for an >= 0 {
		cc := a[an] - '0' + c
		s, c = string(cc%2+'0')+s, cc/2
		an--
	}
	for bn >= 0 {
		cc := b[bn] - '0' + c
		s, c = string(cc%2+'0')+s, cc/2
		bn--
	}
	if c == 1 {
		s = "1" + s
	}
	return s
}
