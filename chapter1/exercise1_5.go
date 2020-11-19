package chapter1

func OneAway(a, b string) bool {
	var useA bool
	if len(a) > len(b) {
		useA = true
	}
	bigMap, littleMap := make(map[rune]int), make(map[rune]int)
	if useA {
		for _, char := range a {
			bigMap[char]++
		}
		for _, char := range b {
			littleMap[char]++
		}
	} else {
		for _, char := range a {
			bigMap[char]++
		}
		for _, char := range b {
			littleMap[char]++
		}
	}
	oneOffs := 0
	for k, v := range bigMap {
		if littleMap[k] != v {
			oneOffs++
		}
	}

	return oneOffs < 2
}
