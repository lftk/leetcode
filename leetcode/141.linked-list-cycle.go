// https://leetcode.com/problems/linked-list-cycle/

package leetcode

/*
#include <stdbool.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

bool hasCycle(struct ListNode *head) {
    struct ListNode *p1 = head, *p2 = head;

    if (head == 0) {
        return false;
    }

    while (p2->next != 0 && p2->next->next != 0) {
        p1 = p1->next;
        p2 = p2->next->next;
        if (p1 == p2) {
            return true;
        }
    }
    return false;
}
*/
import "C"

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	p1, p2 := head, head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			return true
		}
	}
	return false
}
