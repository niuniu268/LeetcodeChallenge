package Methods

import (
	"Leetcode/List/RemoveElements/Methods"
)

func reverseList(head *Methods.ListNode) *Methods.ListNode {

	var pre *Methods.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next

	}

	return pre

}

func reverList2(head *Methods.ListNode) *Methods.ListNode {

	return help(nil, head)
}

func help(pre, head *Methods.ListNode) *Methods.ListNode {

	if head == nil {
		return pre
	}

	next := head.Next
	head.Next = pre
	return help(head, next)
}
