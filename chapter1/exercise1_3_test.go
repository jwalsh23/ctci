package chapter1_test

import (
	"github.com/jwalsh23/ctci/chapter1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlify(t *testing.T) {
	input := "Mr John Smith    "
	expected := "Mr%20John%20Smith"

	assert.Equal(t, expected, chapter1.Urlify(input))
}
