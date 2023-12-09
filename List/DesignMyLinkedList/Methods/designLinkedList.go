package Methods

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
