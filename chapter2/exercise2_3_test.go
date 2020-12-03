package chapter2_test

import (
	"testing"

	"github.com/jwalsh23/ctci/chapter2"
	"github.com/stretchr/testify/assert"
)

func TestRemoveMiddle(t *testing.T) {
	t.Run("removes expected when odd", func(t *testing.T) {
		testObject := chapter2.LinkedList{}
		testObject.Add(1)
		testObject.Add(2)
		testObject.Add(3)
		testObject.Add(4)
		testObject.Add(5)

		expected := "1, 2, 4, 5"

		chapter2.RemoveMiddle(testObject)

		assert.Equal(t, expected, testObject.String())
	})
	t.Run("removes expected when even", func(t *testing.T) {
		testObject := chapter2.LinkedList{}
		testObject.Add(1)
		testObject.Add(2)
		testObject.Add(3)
		testObject.Add(4)

		expected := "1, 2, 4"

		chapter2.RemoveMiddle(testObject)

		assert.Equal(t, expected, testObject.String())
	})
}
