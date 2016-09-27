// https://leetcode.com/problems/rotate-function/

package leetcode

func maxRotateFunction(A []int) int {
	fn := func(nums []int, n int) (sum int) {
		for _, num := range nums {
			if n++; n >= len(nums) {
				n = 0
			}
			sum += n * num
		}
		return
	}

	var max int
	for i := 0; i < len(A); i++ {
		sum := fn(A, i)
		if i == 0 || max < sum {
			max = sum
		}
	}
	return max
}
