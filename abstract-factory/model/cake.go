package model

import "fmt"

type Cake struct{}

func (Cake) Eat() {
	fmt.Println("Eat Case")
}
