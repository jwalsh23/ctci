package chapter1_test

import (
	"github.com/jwalsh23/ctci/chapter1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZero(t *testing.T) {
	t.Run("single zero", func(t *testing.T) {
		input := [][]int{
			{1,2,3,4},
			{5,0,7,8},
			{9,10,11,12},
			{13,14,15,16},
		}

		expected := [][]int {
			{1,0,3,4},
			{0,0,0,0},
			{9,0,11,12},
			{13,0,15,16},
		}

		chapter1.Zero(input)
		assert.Equal(t, expected, input)
	})
	t.Run("zero in same column", func(t *testing.T) {
		input := [][]int{
			{1,2,3,4},
			{5,0,7,8},
			{9,0,11,12},
			{13,14,15,16},
		}

		expected := [][]int {
			{1,0,3,4},
			{0,0,0,0},
			{0,0,0,0},
			{13,0,15,16},
		}

		chapter1.Zero(input)
		assert.Equal(t, expected, input)
	})
	t.Run("zero in same row", func(t *testing.T) {
		input := [][]int{
			{1,2,3,4},
			{5,0,7,0},
			{9,10,11,12},
			{13,14,15,16},
		}

		expected := [][]int {
			{1,0,3,0},
			{0,0,0,0},
			{9,0,11,0},
			{13,0,15,0},
		}

		chapter1.Zero(input)
		assert.Equal(t, expected, input)
	})
}