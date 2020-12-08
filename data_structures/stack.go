package data_structures

type stackNode struct {
	Value int
	prev *stackNode
}

type Stack struct {
	top *stackNode
}

func (s *Stack) Push(val int) {
	node := &stackNode{Value: val}
	if s.top == nil {
		s.top = node
		return
	}
	node.prev = s.top
	s.top = node
}

func(s *Stack) Pop() int {
	if s.top == nil {
		return 0
	}
	val := s.top.Value
	s.top = s.top.prev
	return val
}
