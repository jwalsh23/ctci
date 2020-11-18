package chapter1

func IsUniqueHashMap(s string) bool {
	uniqueChars := make(map[rune]struct{})
	for _, c := range s {
		if _, ok := uniqueChars[c]; ok {
			return false
		}
		uniqueChars[c] = struct{}{}
	}
	return true
}

func IsUniqueArray(s string) bool {
	charSeen := make([]bool, 128)
	for _, c := range s {
		if charSeen[c] {
			return false
		}
		charSeen[c] = true
	}
	return true
}

func IsUniqueNoExtraDataStructure(s string) bool {
	var existsInt int
	for _, c := range s {
		val := c - 'a'
		if (existsInt & 1 << val) > 0 {
			return false
		}
		existsInt |= 1 << val
	}
	return true
}