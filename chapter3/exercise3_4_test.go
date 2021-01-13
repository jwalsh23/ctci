package chapter3_test

import (
	"testing"

	"github.com/jwalsh23/ctci/chapter3"
	"github.com/stretchr/testify/assert"
)

func TestMyQueue_Dequeue(t *testing.T) {
	t.Run("works as expected", func(t *testing.T) {
		testObject := chapter3.NewMyQueue()
		testObject.Enqueue(1)
		testObject.Enqueue(2)
		testObject.Enqueue(3)
		testObject.Enqueue(4)
		testObject.Enqueue(5)

		assert.Equal(t, 1, testObject.Dequeue())
		assert.Equal(t, 2, testObject.Dequeue())
		assert.Equal(t, 3, testObject.Dequeue())
		assert.Equal(t, 4, testObject.Dequeue())
		assert.Equal(t, 5, testObject.Dequeue())
	})
}
