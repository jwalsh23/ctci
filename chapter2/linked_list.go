package chapter2

import (
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	n *Node
}

func (l *LinkedList) Add(val int) {
	node := &Node{Value: val}
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

func (l *LinkedList) Delete(val int) {
	n := l.n
	if n == nil {
		return
	}
	if n.Value == val {
		l.n = n.Next
	}
	for n.Next != nil {
		if n.Next.Value == val {
			n.Next = n.Next.Next
		}
		n = n.Next
	}
}

func (l *LinkedList) String() string {
	n := l.n
	var valSlc []string
	for n != nil {
		valSlc = append(valSlc, strconv.Itoa(n.Value))
		n = n.Next
	}
	return strings.Join(valSlc, ", ")
}
