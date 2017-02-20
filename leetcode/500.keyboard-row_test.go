package leetcode

import (
	"testing"
)

func Test_findWords(t *testing.T) {
	dstWords := findWords([]string{"Hello", "Alaska", "Dad", "Peace"})
	if len(dstWords) != 2 {
		t.Error(dstWords)
		return
	}
	for i, word := range []string{"Alaska", "Dad"} {
		if dstWords[i] != word {
			t.Error(i, word, dstWords[i])
			break
		}
	}
}
