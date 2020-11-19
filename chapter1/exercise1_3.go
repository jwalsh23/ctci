package chapter1

import (
	"unicode"
)

//func Urlify(s string) string {
//	return strings.Join(strings.Split(strings.TrimRight(s, " "), " "), "%20")
//}

func Urlify(s string) string {
	countSpaces := 0
	for i := len(s) -1; i > -1; i-- {
		if unicode.IsSpace(rune(s[i])) {
			countSpaces++
		} else {
			break
		}
	}
	trueLength := len(s) - countSpaces
	index := len(s) - 1
	byteSlc := make([]byte, len(s))
	for i := trueLength-1; i >= 0; i-- {
		if s[i] == ' ' {
			byteSlc[index] = '0'
			byteSlc[index -1] = '2'
			byteSlc[index -2] = '%'
			index-=3
		} else {
			byteSlc[index] = s[i]
			index--
		}
	}
	return string(byteSlc)
}
