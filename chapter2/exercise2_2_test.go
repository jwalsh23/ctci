package chapter2_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/jwalsh23/ctci/chapter2"
	"github.com/stretchr/testify/assert"
)

func TestReturnKthNode(t *testing.T) {
	t.Run("returns error when kth to last doesnt exist", func(t *testing.T) {
		testObject := chapter2.LinkedList{}
		testObject.Add(1)
		testObject.Add(2)
		testObject.Add(3)
		testObject.Add(4)

		_, actualErr := chapter2.ReturnKthNode(testObject, 5)

		expectedErrStr := "list contains 4 elements. 5th element to last does not exist"
		assert.EqualError(t, actualErr, expectedErrStr)
	})
	t.Run("returns expected value", func(t *testing.T) {
		testObject := chapter2.LinkedList{}
		testObject.Add(1)
		testObject.Add(2)
		testObject.Add(3)
		testObject.Add(4)

		actual, actualErr := chapter2.ReturnKthNode(testObject, 2)
		require.NoError(t, actualErr)
		assert.Equal(t, 3, actual)
	})
}
