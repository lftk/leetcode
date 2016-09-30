// https://leetcode.com/problems/remove-linked-list-elements/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	ptr := &ListNode{Next: head}
	cur, next := ptr, ptr.Next
	for next != nil {
		if next.Val == val {
			cur.Next = next.Next
		} else {
			cur = next
		}
		next = next.Next
	}
	return ptr.Next
}
