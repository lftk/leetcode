// https://leetcode.com/problems/ransom-note/

package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, r := range magazine {
		m[r] = m[r] + 1
	}
	for _, r := range ransomNote {
		if m[r]--; m[r] < 0 {
			return false
		}
	}
	return true
}
