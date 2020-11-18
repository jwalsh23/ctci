package chapter1_test

import (
	"github.com/jwalsh23/ctci/chapter1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPermutation (t *testing.T) {
	t.Run("is permutation", func(t *testing.T) {
		str1 := "abcd"
		str2 := "cabd"

		assert.True(t, chapter1.IsPermutation(str1, str2))
	})

	t.Run("is not permutation", func(t *testing.T) {
		t.Run("same length", func(t *testing.T) {
			str1 := "abcd"
			str2 := "dajb"

			assert.False(t, chapter1.IsPermutation(str1, str2))
		})
		t.Run("differnt lengths", func(t *testing.T) {
			str1 := "abcderer"
			str2 := "dajb"

			assert.False(t, chapter1.IsPermutation(str1, str2))
		})
	})
}
