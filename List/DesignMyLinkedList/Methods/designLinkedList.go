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
