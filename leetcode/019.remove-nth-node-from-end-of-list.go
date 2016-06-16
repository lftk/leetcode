// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	/* 4 ms
	var count int
	for node := head; node != nil; node = node.Next {
		count++
	}
	if n = count - n + 1; n == 1 {
		return head.Next
	}
	for node := head; node != nil; node = node.Next {
		if n--; n == 1 {
			if next := node.Next; next != nil {
				node.Next = next.Next
			}
			break
		}
	}
	return head
	*/
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
