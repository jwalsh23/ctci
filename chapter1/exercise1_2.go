package chapter1

func IsPermutation(a,b string) bool {
	if len(a) != len(b) {
		return false
	}
	countSlc := make([]int, 128)
	for i := 0; i < len(a); i ++ {
		countSlc[a[i]]++
		countSlc[b[i]]--
	}
	for _, count := range countSlc {
		if count != 0 {
			return false
		}
	}
	return true
}
