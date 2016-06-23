// https://leetcode.com/problems/reverse-words-in-a-string/

package leetcode

/*
#include <stdlib.h>

void reverse(char* begin, char* end) {
    while (begin < end) {
        char c = *begin;
        *begin = *end;
        *end = c;
        ++begin;
        --end;
    }
}

void reverseWords(char *s) {
    if (s == 0) return;

    int begin = 0, end = 0;
    do {
        begin = end;
        while (s[begin] == ' ') ++begin;
        end = begin;
        while (s[end] != '\0' && s[end] != ' ') ++end;
        reverse(s + begin, s + end - 1);
    } while (begin < end);

    reverse(s, s + end - 1);

    int i = 0, j = 0;
    while (i < end) {
        while (i < end && s[i] == ' ') i++;
        while (i < end && s[i] != ' ') s[j++] = s[i++];
        s[j++] = ' ';
    }
    while (--j >= 0 && (s[j] == '\0' || s[j] == ' '));
    s[j + 1] = '\0';
}
*/
import "C"
import "unsafe"

func reverseWords(s string) string {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.reverseWords(cs)
	return C.GoString(cs)
}
