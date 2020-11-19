package chapter1_test

import (
	"github.com/jwalsh23/ctci/chapter1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindromePermutation(t *testing.T) {
	t.Run("is permutation", func(t *testing.T) {
		input := "Tact Coa"

		assert.True(t, chapter1.PalindromePermutation(input))
	})
	t.Run("is not permutation", func(t *testing.T) {
		input := "actc"

		assert.False(t, chapter1.PalindromePermutation(input))
	})
}