package chapter1

import (
	"strings"
)

func IsRotation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var startPoints []int
	for i := range s2 {
		if s2[i] == s1[0] {
			startPoints = append(startPoints, i)
		}
	}
	if len(startPoints) == 0 {
		return false
	}
	for _, sp := range startPoints {
		s1Index := 1
		s2Index := sp + 1
		for {
			if s2Index == len(s2) {
				s2Index = 0
			}
			if s1[s1Index] != s2[s2Index] {
				break
			}
			s1Index++
			s2Index++
			if s1Index == len(s1) {
				return true
			}
		}
	}
	return false
}

func IsRotationEasy(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1 += s1
	return strings.Count(s1, s2) == 1
}
