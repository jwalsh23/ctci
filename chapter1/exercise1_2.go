package chapter1

func IsPermutation(a,b string) bool {
	if len(a) != len(b) {
		return false
	}
	var arrACharCount [128]int
	var arrBCharCount [128]int
	for i := 0; i < len(a); i ++ {
		arrACharCount[a[i]]++
		arrBCharCount[b[i]]++
	}
	return arrACharCount == arrBCharCount
}
