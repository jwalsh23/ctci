package chapter2

import (
	"github.com/jwalsh23/ctci/data_structures"
)

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

func IsPalindromeTwoPointer(list data_structures.LinkedList) bool {
	fast, slow := list.N, list.N
	stack := data_structures.Stack{}
	for fast != nil && fast.Next != nil {
		stack.Push(slow.Value)
		slow = slow.Next
		fast = fast.Next.Next
	}
	// this is for odd lists
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		if slow.Value != stack.Pop() {
			return false
		}
		slow = slow.Next
	}
	return true
}
