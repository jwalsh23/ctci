package chapter3

const (
	Cat PetType = iota
	Dog
)

type PetType int

type AnimalShelter struct {
	animals []*Pet
}

type Pet struct {
	Name string
	Type PetType
}

func (as *AnimalShelter) Enqueue(p Pet) {
	as.animals = append(as.animals, &p)
}

func (as *AnimalShelter) DequeueAny() *Pet {
	switch len(as.animals) {
	case 0:
		return nil
	case 1:
		retVal := as.animals[0]
		as.animals = nil
		return retVal
	default:
		retVal := as.animals[0]
		as.animals = as.animals[1:]
		return retVal
	}
}

func (as *AnimalShelter) DequeueDog() *Pet {
	return as.dequeue(func(pet *Pet) bool {
		return pet.Type == Dog
	})
}

func (as *AnimalShelter) DequeueCat() *Pet {
	return as.dequeue(func(pet *Pet) bool {
		return pet.Type == Cat
	})
}

func (as *AnimalShelter) dequeue(filter func(pet *Pet) bool) *Pet {
	for i := 0; i < len(as.animals); i++ {
		if filter(as.animals[i]) {
			returnVal := as.animals[i]
			if i < len(as.animals)-2 {
				as.animals = append(as.animals[:i], as.animals[i+1:]...)
			} else {
				as.animals = as.animals[:i]
			}
			return returnVal
		}
	}
	return nil
}
