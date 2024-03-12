package model

import "fmt"

type Bread struct{}

func (Bread) Eat() {
	fmt.Println("Eat Bread!")
}
