package Methods

type MyStack struct {
	queue []int //创建一个队列
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{ //初始化
		queue: make([]int, 0),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	//添加元素
	this.queue = append(this.queue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	n := len(this.queue) - 1 //判断长度
	for n != 0 {             //除了最后一个，其余的都重新添加到队列里
		val := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, val)
		n--
	}
	//弹出元素
	val := this.queue[0]
	this.queue = this.queue[1:]
	return val

}

/** Get the top element. */
func (this *MyStack) Top() int {
	//利用Pop函数，弹出来的元素重新添加
	val := this.Pop()
	this.queue = append(this.queue, val)
	return val
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
