package algs

func BinarySearch(input []int, val int) bool {
	if len(input) == 0 {
		return false
	}
	if len(input) == 1 {
		return input[0] == val
	}
	middle := len(input) / 2
	if input[middle] == val {
		return true
	} else if input[middle] < val {
		return BinarySearch(input[middle:], val)
	} else {
		return BinarySearch(input[:middle], val)
	}
}
