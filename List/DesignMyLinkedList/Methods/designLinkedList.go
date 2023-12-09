package Methods

import (
	"fmt"
)

type SingleNode struct {
	Val  int
	Next *SingleNode
}

type MyLinkedList struct {
	dummyHead *SingleNode
	Size      int
}

func Constructor() MyLinkedList {

	newNode := &SingleNode{
		Val:  -9999,
		Next: nil,
	}

	return MyLinkedList{
		dummyHead: newNode,
		Size:      0,
	}
}

func (this *MyLinkedList) Get(index int) int {

	if this == nil || index < 0 || index >= this.Size {
		return -1
	}

	cur := this.dummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	return cur.Val

}

func (this *MyLinkedList) AddAtHead(val int) {

	newNode := new(SingleNode)
	newNode.Val = val

	newNode.Next = this.dummyHead.Next
	this.dummyHead.Next = newNode

	this.Size++
}

func (this MyLinkedList) AddAtTail(val int) {

	newNode := new(SingleNode)
	newNode.Val = val

	cur := this.dummyHead

	for cur.Next != nil {

		cur = cur.Next
	}

	cur.Next = newNode
	this.Size++
}

func (this MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		index = 0
	} else if index > this.Size {
		return
	}

	s := new(SingleNode)
	s.Val = val

	cur := this.dummyHead

	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	s.Next = cur.Next
	cur.Next = s
	this.Size++
}

func (this MyLinkedList) DeleteAtIndex(index int) {

	if index < 0 || index >= this.Size {
		return
	}

	cur := this.dummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next

	}
	if cur.Next != nil {

		cur.Next = cur.Next.Next
	}

	this.Size--

}

func (this MyLinkedList) printLinkedList() {
	cur := this.dummyHead

	for cur.Next != nil {
		fmt.Println(cur.Next.Val)
		cur = cur.Next
	}
}
