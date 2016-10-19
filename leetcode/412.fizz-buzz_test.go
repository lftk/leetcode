package leetcode

import (
	"strings"
	"testing"
)

func Test_fizzBuzz(t *testing.T) {
	s15 := strings.Join(fizzBuzz(15), "-")
	if s15 != "1-2-Fizz-4-Buzz-Fizz-7-8-Fizz-Buzz-11-Fizz-13-14-FizzBuzz" {
		t.Error(15, s15)
	}
}
