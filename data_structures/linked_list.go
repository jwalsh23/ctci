package data_structures

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	n     *Node
	count int
}

func (l *LinkedList) Add(value int) {
	if l == nil {
		return
	}
	node := &Node{Value: value}
	l.count++
	if l.n == nil {
		l.n = node
		return
	}
	n := l.n
	for n.Next != nil {
		n = n.Next
	}
	n.Next = node
}

func (l *LinkedList) DeleteValue(value int) {
	if l.n == nil {
		return
	}
	n := l.n
	if l.n.Value == value {
		l.n = n.Next
		l.count--
	}
	for n.Next != nil {
		if n.Next.Value == value {
			l.count--
			n.Next = n.Next.Next
			continue
		}
		n = n.Next
	}
	if l.count == 0 {
		*l = LinkedList{}
	}
}

func (l *LinkedList) DeleteIndex(index int) {
	if index > l.count-1 {
		return
	}
	n := l.n
	if index == 0 {
		l.count--
		l.n = n.Next
		return
	}
	count := 0
	for n.Next != nil {
		count++
		if count == index {
			l.count--
			n.Next = n.Next.Next
			return
		}
		n = n.Next
	}
}

func (l *LinkedList) Slice() []int {
	var nodeSlc []int
	if l == nil || l.n == nil {
		return nodeSlc
	}
	n := l.n
	for n != nil {
		nodeSlc = append(nodeSlc, n.Value)
		n = n.Next
	}
	return nodeSlc
}

func (l *LinkedList) Empty() bool {
	return l == nil || l.count == 0
}
