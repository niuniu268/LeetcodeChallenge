package Methods

type MyQueue struct {
	stackIn  []int
	stackOut []int
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}

func (q MyQueue) Push(x int) {
	q.stackIn = append(q.stackIn, x)
}

func (q MyQueue) Pop() int {
	inLen, outLen := len(q.stackIn), len(q.stackOut)
	if outLen == 0 {
		if inLen == 0 {
			return -1
		}
		for i := inLen - 1; i >= 0; i-- {

			q.stackOut = append(q.stackOut, q.stackIn[i])

		}
		q.stackIn = []int{}
		outLen = len(q.stackOut)
	}
	val := q.stackOut[outLen-1]
	q.stackOut = q.stackOut[:outLen-1]
	return val
}

func (q MyQueue) Peek() int {
	val := q.Pop()
	if val == -1 {
		return -1
	}

	q.stackOut = append(q.stackOut, val)
	return val
}

func (q MyQueue) Empty() bool {
	return len(q.stackIn) == 0 && len(q.stackOut) == 0
}
