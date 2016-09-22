// https://leetcode.com/problems/reverse-bits/

package leetcode

/*
#include <stdlib.h>

typedef unsigned int uint32_t;

uint32_t reverseBits(uint32_t n) {
    uint32_t m = 0;
    for (int i = 0; i < 32; i++) {
        m = (m << 1) | (n & 1);
        n = n >> 1;
    }
    return m;
}
*/
import "C"

func reverseBits(n uint32) uint32 {
	m := C.reverseBits(C.uint32_t(n))
	return uint32(m)
}
