package chapter1_test

import (
	"testing"

	"github.com/jwalsh23/ctci/chapter1"
	"github.com/stretchr/testify/assert"
)

func TestIsRotation(t *testing.T) {
	t.Run("is rotation", func(t *testing.T) {
		s1 := "whatisit"
		s2 := "atisitwh"

		assert.True(t, chapter1.IsRotation(s1, s2))
	})
	t.Run("is not rotation", func(t *testing.T) {
		s1 := "whatisit"
		s2 := "isitthat"

		assert.False(t, chapter1.IsRotation(s1, s2))
	})
}

func TestIsRotationEasy(t *testing.T) {
	t.Run("is rotation", func(t *testing.T) {
		s1 := "whatisit"
		s2 := "atisitwh"

		assert.True(t, chapter1.IsRotationEasy(s1, s2))
	})
	t.Run("is not rotation", func(t *testing.T) {
		s1 := "whatisit"
		s2 := "isitthat"

		assert.False(t, chapter1.IsRotationEasy(s1, s2))
	})
}
