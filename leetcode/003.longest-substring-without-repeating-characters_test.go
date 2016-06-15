package leetcode

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
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
