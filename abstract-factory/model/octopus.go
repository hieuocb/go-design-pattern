package model

import "fmt"

type Octopus struct{}

func (Octopus) Eat() {
	fmt.Println("Eat Octopus")
}
