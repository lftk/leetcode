// https://leetcode.com/problems/move-zeroes/

package leetcode

/*
void moveZeroes(int* nums, int numsSize) {
	int i = 0, k = 0;
	while (i < numsSize) {
		if (nums[i] == 0) {
			++k;
		} else if (k != 0) {
			nums[i - k] = nums[i];
		}
		++i;
	}
	while (k > 0) {
		nums[numsSize - k] = 0;
		--k;
	}
}
*/
import "C"
import "unsafe"

func moveZeroes(nums []int32) []int32 {
    if len(nums) > 0 {
        C.moveZeroes((*C.int)(unsafe.Pointer(&nums[0])), (C.int)(len(nums)))
    }
    return nums
}
