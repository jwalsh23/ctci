package chapter1_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsUniqueStandard(t *testing.T) {
	t.Run("is not unique", func(t *testing.T) {
		input := "abcdefghiayur"

		isUnique := chapter1.IsUniqueHashMap(input)

		assert.False(t, isUnique)
	})
	t.Run("is unique", func(t *testing.T) {
		input := "abcdef63g/jk"

		isUnique := chapter1.IsUniqueHashMap(input)

		assert.True(t, isUnique)
	})
}

func TestIsUniqueArray(t *testing.T) {
	t.Run("is not unique", func(t *testing.T) {
		input := "abcdefghiayur"

		isUnique := chapter1.IsUniqueArray(input)

		assert.False(t, isUnique)
	})
	t.Run("is unique", func(t *testing.T) {
		input := "abcdef63g/jk"

		isUnique := chapter1.IsUniqueArray(input)

		assert.True(t, isUnique)
	})
}
func TestIsUniqueNoExtraDataStructure(t *testing.T) {
	t.Run("is not unique", func(t *testing.T) {
		input := "abcdefghiayur"

		isUnique := chapter1.IsUniqueNoExtraDataStructure(input)

		assert.False(t, isUnique)
	})
	t.Run("is unique", func(t *testing.T) {
		input := "abcdef63g/jk"

		isUnique := chapter1.IsUniqueNoExtraDataStructure(input)

		assert.True(t, isUnique)
	})
}