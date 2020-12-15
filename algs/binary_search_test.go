package algs_test

import (
	"testing"

	"github.com/jwalsh23/ctci/algs"
	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	t.Run("value not found", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 8, 10, 20, 30}

		found := algs.BinarySearch(input, 9)
		assert.False(t, found)
	})

	t.Run("value found", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 8, 10, 20, 30}

		found := algs.BinarySearch(input, 3)
		assert.True(t, found)
	})
}
