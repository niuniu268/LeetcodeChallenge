package Methods

import (
	"Leetcode/List/RemoveElements/Methods"
)

func swapPairs(head *Methods.ListNode) *Methods.ListNode {
	dummy := &Methods.ListNode{Next: head}

	pre := dummy

	for head != nil && head.Next != nil {
		pre.Next = head.Next
		next := head.Next.Next
		head.Next.Next = head

		head.Next = next
		pre = head
		head = next
	}

	return dummy.Next

}
