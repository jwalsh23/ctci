package chapter1_test

import (
	"github.com/jwalsh23/ctci/chapter1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate(t *testing.T) {
	input := [][]int{
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11,12},
		{13,14,15,16},
	}

	expected := [][]int {
		{13,9,5,1},
		{14,10,6,2},
		{15,11,7,3},
		{16,12,8,4},
	}

	actual := chapter1.Rotate(input)
	assert.Equal(t, expected, actual)
}