package data_structures

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	N     *Node
	Count int
}

func (l *LinkedList) Add(value int) {
	if l == nil {
		return
	}
	node := &Node{Value: value}
	l.Count++
	if l.N == nil {
		l.N = node
		return
	}
	n := l.N
	for n.Next != nil {
		n = n.Next
	}
	n.Next = node
}

func (l *LinkedList) DeleteValue(value int) {
	if l.N == nil {
		return
	}
	n := l.N
	if l.N.Value == value {
		l.N = n.Next
		l.Count--
	}
	for n.Next != nil {
		if n.Next.Value == value {
			l.Count--
			n.Next = n.Next.Next
			continue
		}
		n = n.Next
	}
	if l.Count == 0 {
		*l = LinkedList{}
	}
}

func (l *LinkedList) DeleteIndex(index int) {
	if index > l.Count-1 {
		return
	}
	n := l.N
	if index == 0 {
		l.Count--
		l.N = n.Next
		return
	}
	count := 0
	for n.Next != nil {
		count++
		if count == index {
			l.Count--
			n.Next = n.Next.Next
			return
		}
		n = n.Next
	}
}

func (l *LinkedList) Slice() []int {
	var nodeSlc []int
	if l == nil || l.N == nil {
		return nodeSlc
	}
	n := l.N
	for n != nil {
		nodeSlc = append(nodeSlc, n.Value)
		n = n.Next
	}
	return nodeSlc
}

func (l *LinkedList) Empty() bool {
	return l == nil || l.Count == 0
}
