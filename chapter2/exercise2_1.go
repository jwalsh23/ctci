package chapter2

func RemoveDuplicates(list LinkedList) {
	seen := make(map[int]struct{})
	n := list.n
	seen[n.Value] = struct{}{}
	for n.Next != nil {
		if _, ok := seen[n.Next.Value]; ok {
			n.Next = n.Next.Next
			continue
		}
		if n.Next != nil {
			seen[n.Next.Value] = struct{}{}
		}
		n = n.Next
	}
}
