// https://leetcode.com/problems/palindrome-linked-list/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindromeLinkedList(head *ListNode) bool {
	var size int
	for n := head; n != nil; n = n.Next {
		size++
	}
	n1, n2 := head, head
	for i := 0; i < (size+1)/2; i++ {
		n2 = n2.Next
	}

	if n2 != nil {
		next := n2.Next
		n2.Next = nil
		for next != nil {
			n2, next, next.Next = next, next.Next, n2
		}
	}

	for n2 != nil && n1.Val == n2.Val {
		n1 = n1.Next
		n2 = n2.Next
	}
	return n2 == nil
}
