package chapter3

type Stack struct {
	top    *StackNode
	length int
}

func (s *Stack) Push(val int) {
	node := &StackNode{Val: val}
	s.length++
	if s.top == nil {
		s.top = node
		return
	}
	node.Prev = s.top
	s.top = node
}

func (s *Stack) Pop() int {
	if s.top == nil {
		return 0
	}
	s.length--
	val := s.top.Val
	s.top = s.top.Prev
	return val
}

func (s *Stack) Length() int {
	if s == nil {
		return 0
	}
	return s.length
}

type MultiStack struct {
	childStackLength  int
	currentStackIndex int
	Stacks            []*Stack
}

func NewMultiStack(size int) *MultiStack {
	return &MultiStack{
		childStackLength:  size,
		currentStackIndex: 0,
		Stacks:            []*Stack{new(Stack)},
	}
}

func (ms *MultiStack) Push(val int) {
	if ms.Stacks[ms.currentStackIndex].Length() == ms.childStackLength {
		newStack := &Stack{}
		newStack.Push(val)
		ms.Stacks = append(ms.Stacks, newStack)
		ms.currentStackIndex++
		return
	}
	ms.Stacks[ms.currentStackIndex].Push(val)
}

func (ms *MultiStack) Pop() int {
	if ms.Stacks[ms.currentStackIndex].Length() == 0 {
		ms.currentStackIndex--
		if ms.currentStackIndex > -1 {
			return ms.Stacks[ms.currentStackIndex].Pop()
		} else {
			return 0
		}
	}
	return ms.Stacks[ms.currentStackIndex].Pop()
}

func (ms *MultiStack) PopAt(index int) int {
	if index < len(ms.Stacks) {
		return ms.Stacks[index].Pop()
	}
	return 0
}
