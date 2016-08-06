// https://leetcode.com/problems/nim-game/

package leetcode

/*
#include <stdbool.h>

bool canWinNim(int n) {
    return n % 4 != 0;
}
*/
import "C"

func canWinNim(n int) bool {
	return C.canWinNim(C.int(n)) == true
}
