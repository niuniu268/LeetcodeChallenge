package main

import (
	"Leetcode/List/RemoveElements/Methods"
	"fmt"
)

func main() {

	list := &Methods.ListNode{Val: 1, Next: &Methods.ListNode{Val: 2, Next: &Methods.ListNode{Val: 3, Next: &Methods.ListNode{Val: 4, Next: &Methods.ListNode{Val: 5, Next: &Methods.ListNode{Val: 6, Next: &Methods.ListNode{Val: 4, Next: &Methods.ListNode{Val: 7, Next: &Methods.ListNode{Val: 8}}}}}}}}}
	displayList(list)
	valToRemove := 4
	result := Methods.RemoveElements(list, valToRemove)

	// Displaying the modified linked list after removal
	fmt.Printf("\nList after removing elements with value %d: ", valToRemove)
	displayList(result)
}

func displayList(head *Methods.ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
