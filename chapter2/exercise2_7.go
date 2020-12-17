package chapter2

import (
	"github.com/jwalsh23/ctci/data_structures"
)

func Intersection(list1, list2 data_structures.LinkedList) bool {
	n1 := list1.N
	n2 := list2.N

	list1Seen := make(map[*data_structures.Node]struct{})
	list2Seen := make(map[*data_structures.Node]struct{})
	for n1 != nil || n2 != nil {
		if n1 != nil {
			if _, ok := list2Seen[n1]; ok {
				return true
			}
			list1Seen[n1] = struct{}{}
		}
		if n2 != nil {
			if _, ok := list1Seen[n2]; ok {
				return true
			}
			list2Seen[n2] = struct{}{}
		}
		if n1 != nil {
			n1 = n1.Next
		}
		if n2 != nil {
			n2 = n2.Next
		}
	}
	return false
}

func IntersectionNoExtraStructures(list1, list2 data_structures.LinkedList) *data_structures.Node {
	n1, n2 := list1.N, list2.N
	list1Count, list2Count := 1, 1
	for n1.Next != nil {
		list1Count++
		n1 = n1.Next
	}
	for n2.Next != nil {
		list2Count++
		n2 = n2.Next
	}
	if n1 != n2 {
		return nil
	}
	if list1Count > list2Count {
		n1 := list1.N
		n2 := list2.N
		for i := 0; i < list1Count-list2Count; i++ {
			n1 = n1.Next
		}
		for n1 != nil {
			if n1 == n2 {
				return n1
			}
			n1 = n1.Next
			n2 = n2.Next
		}
	} else {
		n1 := list1.N
		n2 := list2.N
		for i := 0; i < list2Count-list1Count; i++ {
			n2 = n2.Next
		}
		for n2 != nil {
			if n2 == n1 {
				return n2
			}
			n1 = n1.Next
			n2 = n2.Next
		}
	}
	return nil
}
