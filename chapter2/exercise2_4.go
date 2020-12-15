package chapter2

import "github.com/jwalsh23/ctci/data_structures"

func Partition(list *data_structures.LinkedList, val int) {
	var lessThan, greaterThan, lessThanFirst, greaterThanFirst *data_structures.Node

	n := list.N
	count := 0
	for n != nil {
		if n.Value < val {
			if lessThan == nil {
				lessThan = &data_structures.Node{Value: n.Value}
				lessThanFirst = lessThan
			} else {
				lessThan.Next = &data_structures.Node{Value: n.Value}
				lessThan = lessThan.Next
			}
		} else {
			if greaterThan == nil {
				greaterThan = &data_structures.Node{Value: n.Value}
				greaterThanFirst = greaterThan
			} else {
				greaterThan.Next = &data_structures.Node{Value: n.Value}
				greaterThan = greaterThan.Next
			}
		}
		n = n.Next
		count++
	}
	var firstNode *data_structures.Node
	if lessThanFirst != nil {
		firstNode = lessThanFirst
		lessThan.Next = greaterThanFirst
	}
	*list = data_structures.LinkedList{N: firstNode, Count: count}
}
