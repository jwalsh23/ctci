package chapter2

import (
	"errors"
	"fmt"
)

func ReturnKthNode(list LinkedList, k int) (int, error) {
	var nodeSlc []int
	n := list.n
	for n != nil {
		nodeSlc = append(nodeSlc, n.Value)
		n = n.Next
	}
	if nodeSlc == nil {
		return 0, errors.New("input list is empty")
	}
	if k > len(nodeSlc) {
		return 0, fmt.Errorf("list contains %d elements. %d%s element to last does not exist", len(nodeSlc), k, getPostFix(k))
	}
	return nodeSlc[len(nodeSlc)-k], nil
}

func getPostFix(i int) string {
	switch i {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}
