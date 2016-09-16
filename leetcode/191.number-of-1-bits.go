// https://leetcode.com/problems/number-of-1-bits/

package leetcode

/*
#include <stdlib.h>

typedef unsigned int uint32_t;

int hammingWeight(uint32_t n) {
    int c = 0;
    while (n > 0) {
        c += (n & 1);
        n >>= 1;
    }
    return c;
}
*/
import "C"

func hammingWeight(n uint32) int {
	c := C.hammingWeight(C.uint32_t(n))
	return int(c)
}
