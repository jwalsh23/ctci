package chapter1

import "unicode"

func PalindromePermutation(s string) bool {
	seenCharacters := make(map[rune]int)
	for _, c := range s {
		if unicode.IsSpace(c) {
			continue
		}
		if c -'a' < 0 {
			c = c + 32
		}
		seenCharacters[c]++
	}
	var oneLetter bool
	for _, v := range seenCharacters{
		if v == 1 {
			if oneLetter {
				return false
			}
			oneLetter = true
			continue
		}
		if v % 2 != 0 {
			return false
		}
	}
	return true
}