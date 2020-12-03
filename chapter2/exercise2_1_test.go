package chapter2_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jwalsh23/ctci/chapter2"
)

func TestRemoveDuplicates(t *testing.T) {
	t.Run("removes dups", func(t *testing.T) {
		testObject := chapter2.LinkedList{}
		testObject.Add(1)
		testObject.Add(2)
		testObject.Add(3)
		testObject.Add(2)

		expectedStr := "1, 2, 3"

		chapter2.RemoveDuplicates(testObject)
		assert.Equal(t, expectedStr, testObject.String())
	})
	t.Run("no duplicates, so nothing removed", func(t *testing.T) {
		testObject := chapter2.LinkedList{}
		testObject.Add(1)
		testObject.Add(2)
		testObject.Add(3)
		testObject.Add(4)

		expectedStr := "1, 2, 3, 4"

		chapter2.RemoveDuplicates(testObject)
		assert.Equal(t, expectedStr, testObject.String())
	})
}
