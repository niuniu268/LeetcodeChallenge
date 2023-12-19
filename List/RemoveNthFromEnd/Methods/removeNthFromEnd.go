package Methods

import (
	"Leetcode/List/RemoveElements/Methods"
)

func removeNthFromEnd(head *Methods.ListNode, n int) *Methods.ListNode {

	dummyHead := &Methods.ListNode{}
	dummyHead.Next = head
	cur := head
	prev := dummyHead

	i := 1

	for cur != nil {
		cur = cur.Next
		if i > n {
			prev = prev.Next
		}

		i++
	}

	prev.Next = prev.Next.Next

	return dummyHead.Next
}
