// https://leetcode.com/problems/add-strings/

package leetcode

func addStrings(num1 string, num2 string) string {
	add := func(b1, b2 byte, c int) (b byte, n int) {
		n1, n2 := b1-'0', b2-'0'
		s := int(n1) + int(n2) + c
		b = byte(s%10 + '0')
		n = s / 10
		return
	}

	var (
		num string
		b   byte
		c   int
	)
	i, j := len(num1)-1, len(num2)-1
	for i >= 0 && j >= 0 {
		b, c = add(num1[i], num2[j], c)
		num = string(b) + num
		i--
		j--
	}
	for i >= 0 {
		b, c = add(num1[i], '0', c)
		num = string(b) + num
		i--
	}
	for j >= 0 {
		b, c = add(num2[j], '0', c)
		num = string(b) + num
		j--
	}
	if c > 0 {
		num = string(c+'0') + num
	}
	return num
}
