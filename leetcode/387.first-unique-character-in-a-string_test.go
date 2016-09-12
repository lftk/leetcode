package leetcode

import (
	"testing"
)

func Test_firstUniqChar(t *testing.T) {
	i1 := firstUniqChar("loveleetcode")
	if i1 != 2 {
		t.Error(i1, 2)
	}
}
