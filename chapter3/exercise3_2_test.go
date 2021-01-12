package chapter3_test

import (
	"testing"

	"github.com/jwalsh23/ctci/chapter3"
	"github.com/stretchr/testify/assert"
)

func TestStackWithMin_Min(t *testing.T) {
	t.Run("gets min val", func(t *testing.T) {
		testObject := new(chapter3.StackWithMin)
		testObject.Push(20)
		testObject.Push(11)
		testObject.Push(12)
		testObject.Push(14)
		testObject.Push(1)
		testObject.Push(5)

		assert.Equal(t, 1, testObject.Min())
		testObject.Pop()
		assert.Equal(t, 1, testObject.Min())
		testObject.Pop()
		assert.Equal(t, 11, testObject.Min())
		testObject.Pop()
		assert.Equal(t, 11, testObject.Min())
	})
}
