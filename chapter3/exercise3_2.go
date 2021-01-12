package chapter3

type StackNode struct {
	Val  int
	Prev *StackNode
}

type StackWithMin struct {
	min  *StackNode
	last *StackNode
}

func (s *StackWithMin) Push(i int) {
	node := &StackNode{Val: i, Prev: s.last}
	if s.min == nil {
		s.min = node
	} else {
		if s.min.Val > i {
			s.min = &StackNode{Val: i, Prev: s.min}
		}
	}
	s.last = node
}

func (s *StackWithMin) Pop() int {
	if s.last == nil {
		return -1
	}
	val := s.last.Val
	s.last = s.last.Prev
	if s.min != nil && s.min.Val == val {
		s.min = s.min.Prev
	}
	return val
}

func (s *StackWithMin) Min() int {
	if s.min == nil {
		return -1
	}
	return s.min.Val
}
