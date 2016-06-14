package leetcode

import "testing"

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

func TestLengthOfLongestSubstring(t *testing.T) {
	l1 := lengthOfLongestSubstring("abcabcbb")
	if l1 != 3 {
		t.Fatal("lengthOfLongestSubstring(\"abcabcbb\") != 3")
	}
	l2 := lengthOfLongestSubstring("bbbbb")
	if l2 != 1 {
		t.Fatal("lengthOfLongestSubstring(\"bbbbb\") != 1")
	}
	l3 := lengthOfLongestSubstring("pwwkew")
	if l3 != 3 {
		t.Fatal("lengthOfLongestSubstring(\"pwwkew\") != 3")
	}
	l4 := lengthOfLongestSubstring("abba")
	if l4 != 2 {
		t.Fatal("lengthOfLongestSubstring(\"abba\") != 2")
	}
}
