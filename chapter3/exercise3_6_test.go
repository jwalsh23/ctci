package chapter3_test

import (
	"testing"

	"github.com/jwalsh23/ctci/chapter3"
	"github.com/stretchr/testify/assert"
)

func TestAnimalShelter_DequeueAny(t *testing.T) {
	t.Run("dequeues pets in fifo order", func(t *testing.T) {
		testObject := new(chapter3.AnimalShelter)

		testObject.Enqueue(chapter3.Pet{Name: "Sparky", Type: chapter3.Cat})
		testObject.Enqueue(chapter3.Pet{Name: "Spots", Type: chapter3.Cat})
		testObject.Enqueue(chapter3.Pet{Name: "Bob", Type: chapter3.Dog})
		testObject.Enqueue(chapter3.Pet{Name: "Steve", Type: chapter3.Cat})

		assert.Equal(t, &chapter3.Pet{Name: "Sparky", Type: chapter3.Cat}, testObject.DequeueAny())
		assert.Equal(t, &chapter3.Pet{Name: "Spots", Type: chapter3.Cat}, testObject.DequeueAny())
		assert.Equal(t, &chapter3.Pet{Name: "Bob", Type: chapter3.Dog}, testObject.DequeueAny())
		assert.Equal(t, &chapter3.Pet{Name: "Steve", Type: chapter3.Cat}, testObject.DequeueAny())
	})
}

func TestAnimalShelter_DequeueCat(t *testing.T) {
	t.Run("dequeues cats in fifo order", func(t *testing.T) {
		testObject := new(chapter3.AnimalShelter)

		testObject.Enqueue(chapter3.Pet{Name: "Sparky", Type: chapter3.Cat})
		testObject.Enqueue(chapter3.Pet{Name: "Spots", Type: chapter3.Cat})
		testObject.Enqueue(chapter3.Pet{Name: "Bob", Type: chapter3.Dog})
		testObject.Enqueue(chapter3.Pet{Name: "Steve", Type: chapter3.Cat})

		assert.Equal(t, &chapter3.Pet{Name: "Sparky", Type: chapter3.Cat}, testObject.DequeueCat())
		assert.Equal(t, &chapter3.Pet{Name: "Spots", Type: chapter3.Cat}, testObject.DequeueCat())
		assert.Equal(t, &chapter3.Pet{Name: "Steve", Type: chapter3.Cat}, testObject.DequeueCat())
		var nilPet *chapter3.Pet
		assert.Equal(t, nilPet, testObject.DequeueCat())
	})
}

func TestAnimalShelter_DequeueDog(t *testing.T) {
	t.Run("dequeues dogs in fifo order", func(t *testing.T) {
		testObject := new(chapter3.AnimalShelter)

		testObject.Enqueue(chapter3.Pet{Name: "Sparky", Type: chapter3.Dog})
		testObject.Enqueue(chapter3.Pet{Name: "Spots", Type: chapter3.Dog})
		testObject.Enqueue(chapter3.Pet{Name: "Bob", Type: chapter3.Cat})
		testObject.Enqueue(chapter3.Pet{Name: "Steve", Type: chapter3.Dog})

		assert.Equal(t, &chapter3.Pet{Name: "Sparky", Type: chapter3.Dog}, testObject.DequeueDog())
		assert.Equal(t, &chapter3.Pet{Name: "Spots", Type: chapter3.Dog}, testObject.DequeueDog())
		assert.Equal(t, &chapter3.Pet{Name: "Steve", Type: chapter3.Dog}, testObject.DequeueDog())
		var nilPet *chapter3.Pet
		assert.Equal(t, nilPet, testObject.DequeueDog())
	})
}
