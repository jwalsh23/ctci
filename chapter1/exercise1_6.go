package chapter1

import "fmt"

func StringCompression(s string) string {
	prevRune := s[0]
	count := 0
	outStr := ""
	for i := range s {
		if s[i] == prevRune {
			count++
		} else {
			outStr += fmt.Sprintf("%s%d", string(prevRune), count)
			prevRune = s[i]
			count = 1
		}
	}
	outStr += fmt.Sprintf("%s%d", string(prevRune), count)
	if len(outStr) < len(s) {
		return outStr
	}
	return s
}