package chapter2_test

import (
	"testing"

	"github.com/jwalsh23/ctci/chapter2"
	"github.com/jwalsh23/ctci/data_structures"
	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	t.Run("is not", func(t *testing.T) {
		testObject := data_structures.LinkedList{}
		testObject.Add(5)
		testObject.Add(4)
		testObject.Add(4)
		testObject.Add(6)

		assert.False(t, chapter2.IsPalindrome(testObject))
	})
	t.Run("is palindrome", func(t *testing.T) {
		testObject := data_structures.LinkedList{}
		testObject.Add(1)
		testObject.Add(5)
		testObject.Add(4)
		testObject.Add(4)
		testObject.Add(5)
		testObject.Add(1)

		assert.True(t, chapter2.IsPalindrome(testObject))
	})
}

func TestIsPalindromeTwoPointer(t *testing.T) {
	t.Run("is not", func(t *testing.T) {
		testObject := data_structures.LinkedList{}
		testObject.Add(5)
		testObject.Add(4)
		testObject.Add(4)
		testObject.Add(6)

		assert.False(t, chapter2.IsPalindromeTwoPointer(testObject))
	})
	t.Run("is palindrome even", func(t *testing.T) {
		testObject := data_structures.LinkedList{}
		testObject.Add(1)
		testObject.Add(5)
		testObject.Add(4)
		testObject.Add(4)
		testObject.Add(5)
		testObject.Add(1)

		assert.True(t, chapter2.IsPalindromeTwoPointer(testObject))
	})
	t.Run("is palindrome odd", func(t *testing.T) {
		testObject := data_structures.LinkedList{}
		testObject.Add(1)
		testObject.Add(5)
		testObject.Add(6)
		testObject.Add(5)
		testObject.Add(1)

		assert.True(t, chapter2.IsPalindromeTwoPointer(testObject))
	})
}
