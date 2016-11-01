package leetcode

import (
	"testing"
)

func Test_compareVersion(t *testing.T) {
	n1 := compareVersion("1", "1.2")
	if n1 != -1 {
		t.Error(n1, -1)
	}
}
