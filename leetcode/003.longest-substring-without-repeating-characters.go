package leetcode

func lengthOfLongestSubstring(s string) int {
	var max, start int
	letter := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		prev, ok := letter[c]
		if ok && prev >= start {
			start = prev + 1
		}
		letter[c] = i
		len := i - start + 1
		if max < len {
			max = len
		}
	}
	return max
}
