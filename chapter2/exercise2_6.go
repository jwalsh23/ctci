package chapter2

import "github.com/jwalsh23/ctci/data_structures"

func IsPalindrome(list data_structures.LinkedList) bool {
	valStack := data_structures.Stack{}
	n := list.N
	for n != nil {
		valStack.Push(n.Value)
		n = n.Next
	}
	n = list.N
	for n != nil {
		if n.Value != valStack.Pop() {
			return false
		}
		n = n.Next
	}
	return true
}
