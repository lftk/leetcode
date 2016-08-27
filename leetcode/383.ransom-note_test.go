package leetcode

import (
	"testing"
)

func Test_canConstruct(t *testing.T) {
	testCanConstruct(t, "a", "b", false)
	testCanConstruct(t, "aa", "ab", false)
	testCanConstruct(t, "aa", "aab", true)
}

func testCanConstruct(t *testing.T, ransomNote, magazine string, b bool) {
	b1 := canConstruct(ransomNote, magazine)
	if b1 != b {
		str := "can be  constructed"
		if b == false {
			str = "not " + str
		}
		t.Errorf("\"%s\" %s \"%s\"", ransomNote, str, magazine)
	}
}
