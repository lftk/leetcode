package leetcode

import (
	"testing"
)

func Test_detectCapitalUse(t *testing.T) {
	b1 := detectCapitalUse("USA")
	if !b1 {
		t.Error("USA", b1)
		return
	}

	b2 := detectCapitalUse("leetcode")
	if !b2 {
		t.Error("leetcode", b2)
		return
	}

	b3 := detectCapitalUse("iPhone")
	if b3 {
		t.Error("iPhone", b3)
		return
	}
}
