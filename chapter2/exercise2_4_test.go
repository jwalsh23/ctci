package chapter2_test

import (
	"testing"

	"github.com/jwalsh23/ctci/chapter2"
	"github.com/jwalsh23/ctci/data_structures"
)

func TestPartition(t *testing.T) {
	testObject := data_structures.LinkedList{}
	testObject.Add(3)
	testObject.Add(5)
	testObject.Add(8)
	testObject.Add(5)
	testObject.Add(10)
	testObject.Add(2)
	testObject.Add(1)

	chapter2.Partition(&testObject, 5)

	t.Log(testObject.Slice())
}
