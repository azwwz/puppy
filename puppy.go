package puppy

import (
	"fmt"

	"github.com/azwwz/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks() + "v1.0.0")
}

func From11() {
	fmt.Println("I'm from version 1.1.0")
}
