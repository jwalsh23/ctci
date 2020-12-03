package chapter2

func RemoveMiddle(list LinkedList) {
	var nodeSlc []*Node
	n := list.n
	for n != nil {
		nodeSlc = append(nodeSlc, n)
		n = n.Next
	}
	if len(nodeSlc) < 2 {
		return
	}
	middle := len(nodeSlc) / 2
	nodeSlc[middle-1].Next = nodeSlc[middle+1]
}
