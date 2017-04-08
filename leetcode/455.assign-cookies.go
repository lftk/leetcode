// https://leetcode.com/problems/assign-cookies

package leetcode

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	var (
		n int
		i = len(g) - 1
		j = len(s) - 1
	)
	sort.Ints(g)
	sort.Ints(s)
	for i >= 0 && j >= 0 {
		if s[j] >= g[i] {
			n++
			j--
		}
		i--
	}
	return n
}
