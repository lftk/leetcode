package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	b1 := isPalindrome(123454321)
	if !b1 {
		t.Fatal("123454321 is palindrome")
	}
	b2 := isPalindrome(123456789)
	if b2 {
		t.Fatal("123456789 not is palindrome")
	}
}
