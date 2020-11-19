package chapter1_test

import (
	"testing"

	"github.com/jwalsh23/ctci/chapter1"
	"github.com/stretchr/testify/assert"
)

func TestOneAway(t *testing.T) {
	tests := []struct {
		inputA   string
		inputB   string
		expected bool
	}{
		{
			inputA:   "pale",
			inputB:   "ple",
			expected: true,
		},
		{
			inputA:   "pales",
			inputB:   "pale",
			expected: true,
		},
		{
			inputA:   "pale",
			inputB:   "bale",
			expected: true,
		},
		{
			inputA:   "helllo",
			inputB:   "hello",
			expected: true,
		},
		{
			inputA:   "pale",
			inputB:   "bake",
			expected: false,
		},
	}
	for _, test := range tests {
		actual := chapter1.OneAway(test.inputA, test.inputB)
		assert.Equal(t, test.expected, actual, "%s and %s = %t", test.inputA, test.inputB, actual)
	}
}
