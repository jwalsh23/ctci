package chapter1_test

import (
	"github.com/jwalsh23/ctci/chapter1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringCompression(t *testing.T) {
	t.Run("is compressed", func(t *testing.T) {
		input := "aabcccccaaa"
		expected := "a2b1c5a3"

		assert.Equal(t, expected, chapter1.StringCompression(input))
	})
	t.Run("is not compressed", func(t *testing.T) {
		input := "aab"

		assert.Equal(t, input, chapter1.StringCompression(input))
	})
}