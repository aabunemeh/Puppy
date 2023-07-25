package puppy

import (
	"github.com/aabunemeh/dog"
)

func Bark() string {

	return "Woof!"
}

func Barks() string {

	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrowsUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrowsUp(Barks())
}

func From11() string {
	return "I am from 1.1.0"
}
