package chapter3

import "math"

type SortStack struct {
	stack  *Stack
	topVal int
}

func NewSortStack() *SortStack {
	return &SortStack{stack: new(Stack), topVal: math.MaxInt32}
}

func (s *SortStack) Push(val int) {
	if val < s.Peek() {
		s.stack.Push(val)
		s.topVal = val
		return
	}
	tempStack := new(Stack)
	stackValue := s.stack.Pop()
	for stackValue != 0 && stackValue < val {
		tempStack.Push(stackValue)
		stackValue = s.stack.Pop()
	}
	s.stack.Push(val)
	stackValue = tempStack.Pop()
	for stackValue != 0 {
		s.stack.Push(stackValue)
		stackValue = tempStack.Pop()
	}
}

func (s *SortStack) Pop() int {
	return s.stack.Pop()
}

func (s *SortStack) Peek() int {
	return s.topVal
}

func (s *SortStack) IsEmpty() bool {
	return s.stack.Length() == 0
}
