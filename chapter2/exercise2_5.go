package chapter2

import "github.com/jwalsh23/ctci/data_structures"

func SumListBackward(l1, l2 data_structures.LinkedList) data_structures.LinkedList {
	n1 := l1.N
	n2 := l2.N
	returnList := data_structures.LinkedList{}
	carryOver := 0
	for n1 != nil || n2 != nil {
		var n1Val, n2Val int
		if n1 != nil {
			n1Val = n1.Value
		}
		if n2 != nil {
			n2Val = n2.Value
		}
		sum := n1Val + n2Val + carryOver
		if sum > 10 {
			carryOver = 1
			sum = sum % 10
		}
		returnList.Add(sum)
		if n1 != nil {
			n1 = n1.Next
		}
		if n2 != nil {
			n2 = n2.Next
		}
	}
	return returnList
}
