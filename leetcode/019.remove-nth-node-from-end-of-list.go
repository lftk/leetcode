package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	n1, n2 := head, head
	for i := 0; i < n; i++ {
		n2 = n2.Next
	}
	if n2 == nil {
		return head.Next
	}
	for n2.Next != nil {
		n1 = n1.Next
		n2 = n2.Next
	}
	n1.Next = n1.Next.Next
	return head
}
