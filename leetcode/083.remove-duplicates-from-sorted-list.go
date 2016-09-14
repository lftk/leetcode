// https://leetcode.com/problems/remove-duplicates-from-sorted-list/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p1, p2 := head, head.Next
	for p2 != nil {
		if p1.Val == p2.Val {
			p1.Next, p2 = p2.Next, p2.Next
		} else {
			p1, p2 = p2, p2.Next
		}
	}
	return head
}
