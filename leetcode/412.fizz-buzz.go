// https://leetcode.com/problems/fizz-buzz/

package leetcode

import (
	"fmt"
)

func fizzBuzz(n int) []string {
	ss := make([]string, n+1)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			ss[i] = "FizzBuzz"
		} else if i%5 == 0 {
			ss[i] = "Buzz"
		} else if i%3 == 0 {
			ss[i] = "Fizz"
		} else {
			ss[i] = fmt.Sprintf("%d", i)
		}
	}
	return ss[1:]
}
