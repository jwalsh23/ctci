package chapter3

type MyQueue struct {
	stack *Stack
}

func NewMyQueue() *MyQueue {
	return &MyQueue{stack: new(Stack)}
}

func (mq *MyQueue) Enqueue(val int) {
	mq.stack.Push(val)
}

func (mq *MyQueue) Dequeue() int {
	// 1 2 3 4 5
	tempStack := new(Stack)
	val := mq.stack.Pop()
	for val != 0 {
		tempStack.Push(val)
		val = mq.stack.Pop()
	}
	returnVal := tempStack.Pop()
	val = tempStack.Pop()
	for val != 0 {
		mq.stack.Push(val)
		val = tempStack.Pop()
	}
	return returnVal
}
