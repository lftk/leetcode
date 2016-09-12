// https://leetcode.com/problems/find-the-difference/

package leetcode

/*
#include <string.h>
#include <stdlib.h>

char findTheDifference(char* s, char* t) {
	int i, ch;
	ch = t[strlen(t) - 1];
    for (i = 0; i < strlen(s); ++i) {
        ch += t[i] - s[i];
    }
    return (char)ch;
}
*/
import "C"
import "unsafe"

func findTheDifference(s, t string) byte {
	cs, ct := C.CString(s), C.CString(t)
	defer func() {
		C.free(unsafe.Pointer(cs))
		C.free(unsafe.Pointer(ct))
	}()
	c := C.findTheDifference(cs, ct)
	return byte(c)
}
